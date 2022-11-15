package mypackage

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
)

// Hash the password with HMAC.

// secretKey should be unique, protected, private,
// and not hard-coded like this. Store in environment var
// or in a secure configuration file.
// This is an arbitrary key that should only be used
// for example purposes.
const secretKey = "neictr98y85klfgneghre"

func generateSalt() string {
	randomByte := make([]byte, 32)
	_, err := rand.Read(randomByte)
	if err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(randomByte)
}

// hash the password with a private key
// Use hmac to create a new hashing based on sha256
// Append salt to the password
// Use the newly created hash to hash the password
func HashPassword(plainText string, salt string) string {
	hash := hmac.New(sha256.New, []byte(secretKey))
	io.WriteString(hash, plainText+salt) // append salt to the end of password
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue)
}

func HashingPassword_demo() {
	salt := generateSalt()
	password := "i don't knowit&#"
	hashedPass := HashPassword(password, salt)

	fmt.Printf("password: %s\n", password)
	fmt.Printf("salt: %s\n", salt)
	fmt.Printf("hashed password: %s\n", hashedPass)

	fmt.Println()
	fmt.Println()

	// different salt -> different hash
	salt = generateSalt()
	hashedPass = HashPassword(password, salt)

	fmt.Printf("password: %s\n", password)
	fmt.Printf("different salt: %s\n", salt)
	fmt.Printf("hashed password: %s\n", hashedPass)

}
