package mypackage

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readfull() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 1000) // buffer is too large -> unexpected EOF
	bytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("bytes read: %d\n", bytesRead)

}
