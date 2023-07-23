package localsigner

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/fullsailor/pkcs7"
)

// ReadPublicKey takes a path to a public key and returns an x509.Certificate.
func ReadPublicKey(pubKey string) (*x509.Certificate, error) {
	certPEM, err := os.ReadFile(pubKey)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

// ReadPrivateKey takes a path to a private key and returns an rsa.PrivateKey.
func ReadPrivateKey(prvKey string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(prvKey)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("invalid private key data")
	}
	if block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("invalid key type")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA private key")
	}

	return rsaKey, nil
}

// SignPKCCS7 create PKCS#7 signature
func SignPKCS7(certificate *x509.Certificate, privateKey *rsa.PrivateKey, contents []byte) ([]byte, error) {
	toBeSigned, err := pkcs7.NewSignedData(contents)
	if err != nil {
		return nil, err
	}

	if err := toBeSigned.AddSigner(certificate, privateKey, pkcs7.SignerInfoConfig{}); err != nil {
		return nil, err
	}
	toBeSigned.Detach()

	signature, err := toBeSigned.Finish()
	if err != nil {
		return nil, err
	}

	return signature, nil
}
