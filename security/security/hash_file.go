package security

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"log"
	"os"
)

func HashFile() {
	b, err := os.ReadFile("testdata/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("MD5: %x\n\n", md5.Sum(b))
	fmt.Printf("SHA1: %x\n\n", sha1.Sum(b))
	fmt.Printf("SHA256: %x\n\n", sha256.Sum256(b))
	fmt.Printf("SHA512: %x\n\n", sha512.Sum512(b))
}
