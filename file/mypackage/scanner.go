package mypackage

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// first call Scan()
// then call Tex() or Byte() to get the string or bytes just scanned
func scanner() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// byteSlice := make([]byte, 9)

	// pass a file to NewReader
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// scanner.Split(bufio.ScanWords)
	// scanner.Split(bufio.ScanRunes)
	// scanner.Split(bufio.ScanBytes)

	success := scanner.Scan()
	if !success { // error or EOF
		err = scanner.Err()
		if err == nil {
			log.Println("scan completed and reached EOF")
		} else {
			log.Fatal(err)
		}
	}

	// call scanner.Text() or scanner.Byte() to read what was scanned
	fmt.Printf("1st scan:\n %s\n", scanner.Text())

	// loop Scan() to read
	for scanner.Scan() {
		fmt.Print(scanner.Text(), "\n")
	}
}
