package main

import (
	"encoding/json"
	"fmt"
	"log"

	"sample.dirtree/dt"
)

func main() {
	root, err := dt.NewZipTree("testdata/test.zip")
	if err != nil {
		log.Fatal(err)
	}

	byte, _ := json.Marshal(root)
	fmt.Print(string(byte))

}
