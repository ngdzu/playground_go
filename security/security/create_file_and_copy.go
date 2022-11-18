package security

import (
	"io"
	"log"
	"os"
)

func createfile() {
	originalFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal("can't open file")
	}
	defer originalFile.Close()

	newFile, err := os.Create("newfile.txt")
	if err != nil {
		log.Fatal("can't create new file")
	}

	defer newFile.Close()

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal("can't copy file")
	}

	log.Printf("Copy %d bytes", bytesWritten)

	// Commit the file contents
	// Flushes memory to disk
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
