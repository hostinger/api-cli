// Package creds loads the connector credentials (a client certificate, its
// key and the bootstrap endpoints) from credentials.json or from memory, and
// swaps them without a restart.
package creds

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"
	"time"

	"github.com/hostinger/hcdn-connector/internal/fileguard"
)

const (
	SANScheme     = "hcdn-tunnel:"
	WarnDaysLeft  = 30
	pollInterval  = 10 * time.Second
	warnInterval  = 24 * time.Hour
	gaugeInterval = time.Hour
	clockSkew     = 5 * time.Minute
)

// ErrTunnelIDMismatch: the credentials belong to another tunnel than the config.
var ErrTunnelIDMismatch = errors.New("credentials tunnel_id does not match config tunnel_id")

type File struct {
	TunnelID           string    `json:"tunnel_id"`
	CertificatePEM     string    `json:"certificate_pem"`
	PrivateKeyPEM      string    `json:"private_key_pem"`
	BootstrapEndpoints []string  `json:"bootstrap_endpoints"`
	NotAfter           time.Time `json:"not_after"`
}

// Credentials is an immutable, parsed snapshot of credentials.json.
type Credentials struct {
	File
	Cert     tls.Certificate
	Leaf     *x509.Certificate
	Serial   string
	ModTime  time.Time
	LoadedAt time.Time
}

func (c *Credentials) DaysLeft(now time.Time) float64 {
	return c.Leaf.NotAfter.Sub(now).Hours() / 24
}

// Load reads and validates the file. Files that other users can read are
// refused, and so are files that another account can change, itself or
// through the directory, as in config.Load.
func Load(path, wantTunnelID string) (*Credentials, error) {
	// Stat and read one open file: a rename between two path lookups could swap in a 0644 file.
	f, err := openCredentials(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, fmt.Errorf("credentials %s: not found; download credentials.json from hPanel and install it with %s", path, permHint)
		}
		var pe *fs.PathError
		if errors.As(err, &pe) {
			err = pe.Err
		}
		return nil, fmt.Errorf("credentials %s: %w", path, err)
	}
	defer f.Close()
	st, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("credentials %s: %w", path, err)
	}
	if !st.Mode().IsRegular() {
		return nil, fmt.Errorf("credentials %s: not a regular file", path)
	}
	if err := checkPerm(path, f, st); err != nil {
		return nil, err
	}
	if err := fileguard.Check(path, f); err != nil {
		return nil, fmt.Errorf("credentials %s: %w", path, err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	c, err := Parse(b, wantTunnelID)
	if err != nil {
		return nil, fmt.Errorf("credentials %s: %w", path, err)
	}
	c.ModTime = st.ModTime()
	return c, nil
}

func Parse(b []byte, wantTunnelID string) (*Credentials, error) {
	var f File
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	f.TunnelID = strings.ToLower(f.TunnelID)
	if wantTunnelID != "" && f.TunnelID != strings.ToLower(wantTunnelID) {
		return nil, ErrTunnelIDMismatch
	}
	return fromFile(f)
}

// FromPEM validates a certificate and key the same way as Parse, without the
// JSON envelope (for example after a CSR enrollment that keeps the key in memory).
func FromPEM(tunnelID string, certPEM, keyPEM []byte, bootstrap []string) (*Credentials, error) {
	return fromFile(File{
		TunnelID:           strings.ToLower(tunnelID),
		CertificatePEM:     string(certPEM),
		PrivateKeyPEM:      string(keyPEM),
		BootstrapEndpoints: append([]string(nil), bootstrap...),
	})
}

func fromFile(f File) (*Credentials, error) {
	if len(f.BootstrapEndpoints) == 0 {
		return nil, errors.New("bootstrap_endpoints: required")
	}
	for _, e := range f.BootstrapEndpoints {
		if e == "" || strings.ContainsAny(e, " /:") {
			return nil, fmt.Errorf("bootstrap_endpoints: %q is not a host name", e)
		}
	}
	cert, err := tls.X509KeyPair([]byte(f.CertificatePEM), []byte(f.PrivateKeyPEM))
	if err != nil {
		return nil, fmt.Errorf("certificate/key: %w", err)
	}
	leaf, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return nil, fmt.Errorf("certificate: %w", err)
	}
	cert.Leaf = leaf
	// tunneld requires exactly one hcdn-tunnel:<uuid> URI SAN (docs/protocol.md, section 1).
	var sans []string
	for _, u := range leaf.URIs {
		if u.Scheme == strings.TrimSuffix(SANScheme, ":") {
			sans = append(sans, strings.ToLower(u.Opaque))
		}
	}
	if len(sans) != 1 {
		return nil, fmt.Errorf("certificate: want exactly one %s<uuid> URI SAN, got %d", SANScheme, len(sans))
	}
	if sans[0] != f.TunnelID {
		return nil, fmt.Errorf("certificate URI SAN does not match tunnel_id %s", f.TunnelID)
	}
	if f.NotAfter.IsZero() {
		f.NotAfter = leaf.NotAfter
	}
	return &Credentials{
		File:     f,
		Cert:     cert,
		Leaf:     leaf,
		Serial:   leaf.SerialNumber.Text(16),
		LoadedAt: time.Now(),
	}, nil
}
