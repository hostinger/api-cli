package creds

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/binary"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"
)

const (
	// QuickLifetime is the longest validity window tunneld accepts for a quick leaf.
	QuickLifetime = 24 * time.Hour
	quickBackdate = 5 * time.Minute
)

// QuickID derives the quick tunnel id: the first 16 bytes of
// SHA-256(SPKI DER ‖ uint64_be(NotBefore unix seconds)) with the UUIDv8
// version and variant bits. The edge (tunneld) shares this
// derivation (docs/protocol.md, "Quick class"); testdata/quick_id.json pins it.
func QuickID(spkiDER []byte, notBefore time.Time) string {
	h := sha256.New()
	h.Write(spkiDER)
	var ts [8]byte
	binary.BigEndian.PutUint64(ts[:], uint64(notBefore.Unix()))
	h.Write(ts[:])
	id := h.Sum(nil)[:16]
	id[6] = id[6]&0x0f | 0x80
	id[8] = id[8]&0x3f | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", id[0:4], id[4:6], id[6:8], id[8:10], id[10:16])
}

// IsQuickID reports a canonical UUID with version 8 (the quick class).
// Named tunnels use version 4 ids.
func IsQuickID(id string) bool {
	id = strings.ToLower(id)
	if len(id) != 36 || id[8] != '-' || id[13] != '-' || id[18] != '-' || id[23] != '-' {
		return false
	}
	if _, err := hex.DecodeString(strings.ReplaceAll(id, "-", "")); err != nil {
		return false
	}
	return id[14] == '8' && strings.IndexByte("89ab", id[19]) >= 0
}

// QuickHostname is the public hostname of a quick tunnel in zone.
func QuickHostname(id, zone string) string {
	return strings.ReplaceAll(strings.ToLower(id), "-", "") + "." + zone
}

// NewQuick creates anonymous quick tunnel credentials: a P-256 key and a
// self-signed client leaf whose id binds the key and NotBefore. The key
// stays in memory. Quick credentials cannot be renewed: a new leaf is a new id.
func NewQuick(bootstrap []string, now time.Time) (*Credentials, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	spki, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return nil, err
	}
	nb := now.UTC().Truncate(time.Second).Add(-quickBackdate)
	id := QuickID(spki, nb)
	serial, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, err
	}
	tmpl := &x509.Certificate{
		SerialNumber:          serial,
		NotBefore:             nb,
		NotAfter:              nb.Add(QuickLifetime),
		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		URIs:                  []*url.URL{{Scheme: strings.TrimSuffix(SANScheme, ":"), Opaque: id}},
		BasicConstraintsValid: true,
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		return nil, err
	}
	leaf, err := x509.ParseCertificate(der)
	if err != nil {
		return nil, err
	}
	if QuickID(leaf.RawSubjectPublicKeyInfo, leaf.NotBefore) != id {
		return nil, errors.New("quick certificate: id does not match the encoded key and NotBefore")
	}
	keyDER, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, err
	}
	return FromPEM(id,
		pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}),
		pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER}),
		bootstrap)
}
