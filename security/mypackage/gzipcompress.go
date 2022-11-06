package mypackage

import (
	"compress/gzip"
	"log"
	"os"
)

func gzipcompress() {
	outputFile, err := os.Create("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// create a text file called test.txt,
	// then compress this file to test.txt.gz
	// The resulting gz file will contain the only
	// on file test.txt
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()
	_, err = gzipWriter.Write([]byte("file content"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Compressed file to test.txt.gz")
}
