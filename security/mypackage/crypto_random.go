package mypackage

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

func CryptoRandom() {
	limit := int64(math.MaxInt64)

	// rand.Int create random int
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(limit))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("random int64 number:", randomNumber)

	// Alternatively, you could generate the random bytes
	// and turn them into the specific data type needed.
	// binary.Read() will only read enough bytes to fill the data type
	var number uint32

	// binary.Read generate just enough random bytes to fill in number
	err = binary.Read(rand.Reader, binary.BigEndian, &number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Random 32 bit number: ", number)

	// Generate radom byte slice
	const numByte = 4
	randomBytes := make([]byte, numByte)
	_, err = rand.Read(randomBytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("random 4-byte value:", randomBytes)

}
