package main

import (
	"fmt"
	"os"

	localsigner "github.com/matsuyoshi30/local-signer"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s </path/to/certFile> </path/to/keyFile> </path/to/contents>\n", os.Args[0])
		os.Exit(1)
	}

	pubKey, err := localsigner.ReadPublicKey(os.Args[1])
	if err != nil {
		fmt.Println("failed to read public key", err)
		os.Exit(1)
	}

	prvKey, err := localsigner.ReadPrivateKey(os.Args[2])
	if err != nil {
		fmt.Println("failed to read private key", err)
		os.Exit(1)
	}

	contents, err := os.ReadFile(os.Args[3])
	if err != nil {
		fmt.Println("failed to read contents", err)
		os.Exit(1)
	}

	signature, err := localsigner.SignPKCS7(pubKey, prvKey, contents)
	if err != nil {
		fmt.Println("failed to sign PKCS#7", err)
		os.Exit(1)
	}

	fmt.Printf("Signature: %x\n", signature)
}
