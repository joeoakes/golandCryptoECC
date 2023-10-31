package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Generate a new private-public key pair
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println(err.Error())
	}

	publicKey := &privateKey.PublicKey

	// Message to be signed
	message := "Hello, ECC!"

	// Hash the message
	hashedMessage := sha256.Sum256([]byte(message))

	// Sign the hashed message
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashedMessage[:])
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Signature:", r, s)

	// Verify the signature
	isVerified := ecdsa.Verify(publicKey, hashedMessage[:], r, s)
	fmt.Println("Signature verified:", isVerified)
}
