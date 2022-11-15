package mypackage

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"path/filepath"
)

// encode the private key as PEM file
// PEM is base64 encoding of a key
func GetPrivatePemFromKey(privateKey *rsa.PrivateKey) *pem.Block {
	encodedPrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privatePem := &pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: encodedPrivateKey,
	}
	return privatePem
}

func GeneratePublicPemFromKey(publicKey rsa.PublicKey) *pem.Block {
	encodedPublicKey, err := x509.MarshalPKIXPublicKey(&publicKey) // NOTE must pass a pointer to public key
	if err != nil {
		log.Fatal(err)
	}
	publicPem := &pem.Block{
		Type: "PUBLIC KEY", Bytes: encodedPublicKey,
	}

	return publicPem
}

func SavePemToFile(pemBlock *pem.Block, filename string) {

	parentPath := filepath.Dir(filename)
	if len(parentPath) > 0 {
		err := os.MkdirAll(parentPath, os.ModePerm)
		if err != nil {
			log.Fatal("Error creating path")
		}
	}
	// save public PEM to file
	publicPemOutputFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer publicPemOutputFile.Close()

	err = pem.Encode(publicPemOutputFile, pemBlock)
	if err != nil {
		log.Fatal(err)
	}

}

// # Generate the private key
// openssl genrsa -out priv.pem 2048
// # Extract the public key from the private key
// openssl rsa -in priv.pem -pubout -out public.pem
func GenerateRSAKeyPair(privateKeyFilename, publicKeyFilename string) {
	// generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, key_size)
	if err != nil {
		log.Fatal(err)
	}

	// Encode keys to PEM format
	privatePem := GetPrivatePemFromKey(privateKey)
	publicPem := GeneratePublicPemFromKey(privateKey.PublicKey)

	// save PEM outputs to files

	SavePemToFile(privatePem, privateKeyFilename)
	SavePemToFile(publicPem, publicKeyFilename)
}
