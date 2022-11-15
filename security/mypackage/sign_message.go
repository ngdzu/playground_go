package mypackage

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

const private_key_filename = "testdata/private.pem"
const public_key_filename = "testdata/public.pem"
const message_filename = "testdata/test.txt"
const signature_filename = "signature.txt.256"
const key_size = 2048

func SignMessage(privateKey *rsa.PrivateKey, message []byte) []byte {
	hashed := sha256.Sum256(message)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		log.Fatal("error sigining message. ", err)
	}

	return signature
}

func loadMessageFromFile(messageFilename string) []byte {
	fileData, err := os.ReadFile(messageFilename)
	if err != nil {
		log.Fatal("error reading file:", err)
	}
	return fileData
}

func LoadPrivateKeyFromPemFile(privateKeyFilePath string) *rsa.PrivateKey {
	fileData, err := os.ReadFile(privateKeyFilePath)
	if err != nil {
		log.Fatal("error reading file:", err)
	}

	block, _ := pem.Decode(fileData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatal("Unable to load a valid private key.")
	}
	der := block.Bytes

	// // If the private key is encrypted with a password
	// encrypted := false
	// password := []byte("ca private key password")
	// if encrypted {
	// 	der, err = x509.DecryptPEMBlock(block, password)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	privateKey, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		log.Fatal("error loading private key", err)
	}

	return privateKey
}

func LoadPublicKeyFromPemFile(publicKeyFilename string) *rsa.PublicKey {
	fileData, err := os.ReadFile(publicKeyFilename)
	if err != nil {
		log.Fatal("error reading file:", err)
	}

	block, _ := pem.Decode(fileData)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("Unable to load a valid private key.")
	}

	// NOTE we use PKIX
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal("error loading private key", err)
	}

	return publicKey.(*rsa.PublicKey)
}

func WriteToFile(filename string, data []byte) error {
	file, err := os.OpenFile(
		filename,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func VerifyMessage(publicKey *rsa.PublicKey, signature, message []byte) error {
	// hash the message ourselves
	hashed := sha256.Sum256(message)

	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)

	return err
}

func SignMessage_demo() {
	privateKey := LoadPrivateKeyFromPemFile(private_key_filename)
	message := loadMessageFromFile(message_filename)
	signature := SignMessage(privateKey, message)
	WriteToFile(signature_filename, signature)

	if VerifyMessage(&privateKey.PublicKey, signature, message) == nil {
		fmt.Print("Verify successfully.")
	} else {
		fmt.Print("Verify unsuccessfully")
	}
}
