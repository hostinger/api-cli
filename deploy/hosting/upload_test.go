package hosting

import (
	"bytes"
	"cmp"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"testing"
	"time"
)

// tusServer requires each PATCH to start at the stored offset.
type tusServer struct {
	mu         sync.Mutex
	data       []byte
	length     int64
	created    bool
	failNext   int // PATCH requests to fail before accepting
	failStatus int // what they fail with; 500 when zero
	failCreate int // creation POSTs to fail with 429 before accepting
	path       string
	headers    http.Header
}

func (s *tusServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.path = r.URL.RequestURI()
	s.headers = r.Header.Clone()
	if r.Header.Get("X-Auth") != "auth" || r.Header.Get("X-Auth-Rest") != "rest" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	switch r.Method {
	case http.MethodPost:
		if s.failCreate > 0 {
			s.failCreate--
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		s.length, _ = strconv.ParseInt(r.Header.Get("Upload-Length"), 10, 64)
		s.created = true
		s.data = nil
		w.WriteHeader(http.StatusCreated)
	case http.MethodHead:
		w.Header().Set("Upload-Offset", strconv.Itoa(len(s.data)))
		w.WriteHeader(http.StatusOK)
	case http.MethodPatch:
		if s.failNext > 0 {
			s.failNext--
			w.WriteHeader(cmp.Or(s.failStatus, http.StatusInternalServerError))
			return
		}
		offset, _ := strconv.ParseInt(r.Header.Get("Upload-Offset"), 10, 64)
		if !s.created || offset != int64(len(s.data)) || r.Header.Get("Content-Type") != "application/offset+octet-stream" {
			w.WriteHeader(http.StatusConflict)
			return
		}
		body, _ := io.ReadAll(r.Body)
		s.data = append(s.data, body...)
		w.Header().Set("Upload-Offset", strconv.Itoa(len(s.data)))
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func tempFile(t *testing.T, content []byte) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "site.zip")
	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestUploadInChunks(t *testing.T) {
	ChunkSize = 7
	t.Cleanup(func() { ChunkSize = 10 * 1024 * 1024 })
	srv := &tusServer{}
	ts := httptest.NewServer(srv)
	defer ts.Close()

	content := bytes.Repeat([]byte("abcdefghij"), 5) // 50 bytes: 7 full chunks + 1 byte
	var reports []int64
	err := Upload(context.Background(), Credentials{URL: ts.URL + "/", AuthKey: "auth", RestAuthKey: "rest"}, tempFile(t, content), "dist/my site_1.zip", func(sent, total int64) {
		reports = append(reports, sent)
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(srv.data, content) {
		t.Errorf("server got %d bytes, want %d", len(srv.data), len(content))
	}
	if srv.length != 50 {
		t.Errorf("Upload-Length = %d", srv.length)
	}
	if srv.path != "/dist/my%20site_1.zip?override=true" {
		t.Errorf("path = %s", srv.path)
	}
	if len(reports) != 8 || reports[len(reports)-1] != 50 {
		t.Errorf("progress = %v", reports)
	}
}

func TestUploadResumesAfterServerError(t *testing.T) {
	ChunkSize = 16
	RetryDelays = []time.Duration{time.Millisecond, time.Millisecond}
	t.Cleanup(func() { ChunkSize = 10 * 1024 * 1024 })
	srv := &tusServer{failNext: 1}
	ts := httptest.NewServer(srv)
	defer ts.Close()

	content := bytes.Repeat([]byte("x"), 40)
	err := Upload(context.Background(), Credentials{URL: ts.URL, AuthKey: "auth", RestAuthKey: "rest"}, tempFile(t, content), "site.zip", nil)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(srv.data, content) {
		t.Errorf("server got %d bytes, want %d", len(srv.data), len(content))
	}
}

func TestUploadRetriesRateLimitAndLock(t *testing.T) {
	ChunkSize = 16
	RetryDelays = []time.Duration{time.Millisecond, time.Millisecond}
	t.Cleanup(func() { ChunkSize = 10 * 1024 * 1024 })
	for _, status := range []int{http.StatusTooManyRequests, http.StatusLocked} {
		srv := &tusServer{failNext: 1, failStatus: status, failCreate: 1}
		ts := httptest.NewServer(srv)
		content := bytes.Repeat([]byte("x"), 40)
		err := Upload(context.Background(), Credentials{URL: ts.URL, AuthKey: "auth", RestAuthKey: "rest"}, tempFile(t, content), "site.zip", nil)
		ts.Close()
		if err != nil {
			t.Fatalf("HTTP %d: %v", status, err)
		}
		if !bytes.Equal(srv.data, content) {
			t.Errorf("HTTP %d: server got %d bytes, want %d", status, len(srv.data), len(content))
		}
	}
}

func TestUploadGivesUpOnRejection(t *testing.T) {
	ts := httptest.NewServer(&tusServer{})
	defer ts.Close()
	err := Upload(context.Background(), Credentials{URL: ts.URL, AuthKey: "wrong", RestAuthKey: "rest"}, tempFile(t, []byte("x")), "site.zip", nil)
	if err == nil {
		t.Fatal("expected an error for bad credentials")
	}
}

func TestUploadRejectsUnsafePaths(t *testing.T) {
	for _, p := range []string{"", "/abs.zip", "../up.zip", "a//b.zip", `a\b.zip`} {
		if err := Upload(context.Background(), Credentials{}, "nope", p, nil); err == nil {
			t.Errorf("path %q accepted", p)
		}
	}
}
