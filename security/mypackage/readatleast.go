package mypackage

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readatleast() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 400)
	bytesRead, err := io.ReadAtLeast(file, byteSlice, 400)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("bytes read: %d\n", bytesRead)

}
