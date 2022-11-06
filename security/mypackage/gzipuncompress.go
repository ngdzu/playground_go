package mypackage

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func Gzipuncompress() {
	file, err := os.Open("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()

	outputFile, err := os.Create("uncompressed.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, gzipReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("uncompressed test.txt.gz to uncompressed.txt")
}
