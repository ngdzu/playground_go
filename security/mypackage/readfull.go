package mypackage

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Readfull() {
	file, err := os.Open("testdata/large.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 1000000) // buffer is too large -> unexpected EOF
	bytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("bytes read: %d\n", bytesRead)

	output, err := os.Create("testdata/large2.txt")
	if err != nil {
		log.Fatal(err)
	}
	_, err = output.Write(byteSlice)
	if err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
