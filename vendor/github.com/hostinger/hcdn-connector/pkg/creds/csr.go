package creds

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// GenerateCSR creates a P-256 key and a CSR with an empty subject and no SANs:
// the Hostinger API sets the URI SAN server-side and refuses any subject with 400.
func GenerateCSR() (keyPEM, csrPEM []byte, err error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	der, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, nil, err
	}
	keyPEM = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: der})
	tmpl := &x509.CertificateRequest{SignatureAlgorithm: x509.ECDSAWithSHA256}
	csrDER, err := x509.CreateCertificateRequest(rand.Reader, tmpl, key)
	if err != nil {
		return nil, nil, err
	}
	csrPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrDER})
	return keyPEM, csrPEM, nil
}

// WriteCSR writes key (0600, or owner+SYSTEM on Windows) and CSR (0644),
// refusing to overwrite an existing key.
func WriteCSR(keyPath, csrPath string) error {
	keyPEM, csrPEM, err := GenerateCSR()
	if err != nil {
		return err
	}
	kf, err := createPrivate(keyPath)
	if err != nil {
		return fmt.Errorf("key: %w", err)
	}
	if _, err := kf.Write(keyPEM); err != nil {
		kf.Close()
		return err
	}
	if err := kf.Close(); err != nil {
		return err
	}
	if err := os.WriteFile(csrPath, csrPEM, 0o644); err != nil {
		return fmt.Errorf("csr: %w", err)
	}
	return nil
}
