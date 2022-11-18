package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

func generateKey() []byte {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal(err)
	}

	return randomBytes
}

func encrypt(key, message []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create the byte slice that will hold encrypted message
	cipherText := make([]byte, aes.BlockSize+len(message))

	// Generate the Initialization Vector (IV) nonce
	// which is stored at the beginning of the byte slice
	// The IV is the same length as the AES blocksize
	iv := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	// Choose the block cipher mode of operation
	// Using the cipher feedback (CFB) mode here.
	// CBCEncrypter also available.
	cfb := cipher.NewCFBEncrypter(block, iv)

	// Generate the encrypted message and store it
	// in the remaining bytes after the IV nonce
	cfb.XORKeyStream(cipherText[aes.BlockSize:], message)

	return cipherText, nil
}

func decrypt(key, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// separate nonce from the encrypted message byte
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Decrypt the message using the CFB block mode
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(cipherText, cipherText)

	return cipherText, nil

}
