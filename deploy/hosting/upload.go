package hosting

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Credentials struct {
	URL         string `json:"url"`
	AuthKey     string `json:"auth_key"`
	RestAuthKey string `json:"rest_auth_key"`
}

// Tests lower these.
var (
	ChunkSize   int64 = 10 * 1024 * 1024
	RetryDelays       = []time.Duration{time.Second, 2 * time.Second, 4 * time.Second, 8 * time.Second, 16 * time.Second, 20 * time.Second}
	// The response timeout leaves a 10 MiB chunk as long as it takes on a slow
	// uplink; ctx still cancels.
	uploadHTTP = &http.Client{Transport: &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           (&net.Dialer{Timeout: 30 * time.Second}).DialContext,
		TLSHandshakeTimeout:   30 * time.Second,
		ResponseHeaderTimeout: 2 * time.Minute,
	}}
)

// Upload sends a file to remotePath (relative to the document root) over tus,
// in ChunkSize pieces, resuming from the server's offset after a failure.
func Upload(ctx context.Context, creds Credentials, localPath, remotePath string, progress func(sent, total int64)) error {
	if err := checkRemotePath(remotePath); err != nil {
		return err
	}
	f, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return err
	}
	size := info.Size()

	segments := strings.Split(remotePath, "/")
	for i, s := range segments {
		segments[i] = url.PathEscape(s)
	}
	target := strings.TrimRight(creds.URL, "/") + "/" + strings.Join(segments, "/") + "?override=true"
	auth := func(req *http.Request) {
		req.Header.Set("X-Auth", creds.AuthKey)
		req.Header.Set("X-Auth-Rest", creds.RestAuthKey)
		req.Header.Set("Tus-Resumable", "1.0.0")
	}

	for attempt := 0; ; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, target, nil)
		if err != nil {
			return err
		}
		auth(req)
		req.Header.Set("Upload-Length", strconv.FormatInt(size, 10))
		req.Header.Set("Upload-Offset", "0")
		status, body, err := send(req)
		if err == nil && status == http.StatusCreated {
			break
		}
		if !retryable(status, err) || attempt >= len(RetryDelays) {
			if err != nil {
				return fmt.Errorf("starting upload: %w", err)
			}
			return fmt.Errorf("starting upload: HTTP %d: %s", status, strings.TrimSpace(string(body)))
		}
		if err := wait(ctx, RetryDelays[attempt]); err != nil {
			return err
		}
	}

	var offset int64
	buf := make([]byte, ChunkSize)
	attempt := 0
	for offset < size {
		n, err := io.ReadFull(io.NewSectionReader(f, offset, ChunkSize), buf)
		if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) && !errors.Is(err, io.EOF) {
			return err
		}
		if n == 0 {
			return fmt.Errorf("%s changed while uploading", localPath)
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodPatch, target, bytes.NewReader(buf[:n]))
		if err != nil {
			return err
		}
		auth(req)
		req.Header.Set("Content-Type", "application/offset+octet-stream")
		req.Header.Set("Upload-Offset", strconv.FormatInt(offset, 10))
		status, body, err := send(req)
		if err == nil && status >= 200 && status < 300 {
			attempt = 0
			next := offset + int64(n)
			if h := req.Response; h != nil {
				if v, perr := strconv.ParseInt(h.Header.Get("Upload-Offset"), 10, 64); perr == nil && v > offset {
					next = v
				}
			}
			offset = next
			if progress != nil {
				progress(offset, size)
			}
			continue
		}
		if !retryable(status, err) {
			return fmt.Errorf("upload rejected: HTTP %d: %s", status, strings.TrimSpace(string(body)))
		}
		if attempt >= len(RetryDelays) {
			if err != nil {
				return fmt.Errorf("upload failed after %d retries: %w", attempt, err)
			}
			return fmt.Errorf("upload failed after %d retries: HTTP %d: %s", attempt, status, strings.TrimSpace(string(body)))
		}
		if err := wait(ctx, RetryDelays[attempt]); err != nil {
			return err
		}
		attempt++
		// Resume from what the server has, when it tells us.
		if head, herr := http.NewRequestWithContext(ctx, http.MethodHead, target, nil); herr == nil {
			auth(head)
			if resp, rerr := uploadHTTP.Do(head); rerr == nil {
				resp.Body.Close()
				if v, perr := strconv.ParseInt(resp.Header.Get("Upload-Offset"), 10, 64); perr == nil && resp.StatusCode < 300 && v >= 0 && v <= size {
					offset = v
				}
			}
		}
	}
	return nil
}

// retryable: no response, a 5xx, or a 4xx meaning "not now" (offset conflict,
// timeout, lock from a dropped connection, rate limit).
func retryable(status int, err error) bool {
	if err != nil || status >= 500 {
		return true
	}
	switch status {
	case http.StatusConflict, http.StatusRequestTimeout, http.StatusLocked, http.StatusTooManyRequests:
		return true
	}
	return false
}

func wait(ctx context.Context, d time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(d):
		return nil
	}
}

// send leaves the response on req.Response so headers can be read.
func send(req *http.Request) (int, []byte, error) {
	resp, err := uploadHTTP.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 64*1024))
	req.Response = resp
	return resp.StatusCode, body, nil
}

func checkRemotePath(p string) error {
	if p == "" || strings.HasPrefix(p, "/") || strings.Contains(p, "\\") {
		return fmt.Errorf("unsafe upload path %q", p)
	}
	for _, seg := range strings.Split(p, "/") {
		if seg == "" || seg == "." || seg == ".." {
			return fmt.Errorf("unsafe upload path %q", p)
		}
	}
	return nil
}
