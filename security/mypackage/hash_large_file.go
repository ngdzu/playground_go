package mypackage

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

// We don't read the file all at once since the file is large
// - We create a reader from the file. 
// - THen create a hasher as 
// - Finally, copy the file reader to the hash writer. 
func HashLargeFile() {
	f, err := os.Open("testdata/large.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	hasher := md5.New()

	_, err = io.Copy(hasher, f) // haser implements io.Writer
	if err != nil {
		log.Fatal(err)
	}

	checksum := hasher.Sum(nil) // pass nil because we already copy file to hasher
	fmt.Printf("checksum: %x\n", checksum)
}
