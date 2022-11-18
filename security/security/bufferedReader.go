package security

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func bufferedreader() {
	file, err := os.Open("test.txt")
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

	b, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read 1 byte: %c\n", b)

	// read unitl delimiter
	byteSlice, err = bufferedReader.ReadBytes(',')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read until endline: %s\n", byteSlice)

	stringRead, err := bufferedReader.ReadString('.')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("string read: %s\n", stringRead)

}
