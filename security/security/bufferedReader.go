package security

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func BufferedreaderDemo() {
	file, err := os.Open("testdata/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 9)

	// pass a file to NewReader
	bufferedReader := bufio.NewReader(file)
	bytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes\n", bytesRead)
	fmt.Printf("read contents: %s\n", byteSlice)

	// Read a single byte
	b, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read 1 byte: %c\n", b)

	// read until delimiter
	byteSlice, err = bufferedReader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read until endline: %s\n", byteSlice)

	// Read until '.', returns a string
	stringRead, err := bufferedReader.ReadString('.')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("string read: %s\n", stringRead)

}
