package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/fullsailor/pkcs7"
)

func readPublicKey(pubKey string) (*x509.Certificate, error) {
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

func readPrivateKey(prvKey string) (*rsa.PrivateKey, error) {
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

func main() {
	data := []byte("this is some data to be signed")

	pubKey, err := readPublicKey("signer.crt")
	if err != nil {
		fmt.Println("failed to read public key", err)
		os.Exit(1)
	}

	prvKey, err := readPrivateKey("signer.key")
	if err != nil {
		fmt.Println("failed to read private key", err)
		os.Exit(1)
	}

	toBeSigned, err := pkcs7.NewSignedData(data)
	if err != nil {
		fmt.Println("failed to create signed data", err)
		os.Exit(1)
	}

	if err := toBeSigned.AddSigner(pubKey, prvKey, pkcs7.SignerInfoConfig{}); err != nil {
		fmt.Println("failed to add signer", err)
		os.Exit(1)
	}
	toBeSigned.Detach()

	signature, err := toBeSigned.Finish()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Signature: %x\n", signature)
}
