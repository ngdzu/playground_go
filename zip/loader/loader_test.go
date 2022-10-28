package loader

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestLoader(t *testing.T) {
	l := Loader{}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	// fullpath := filepath.Join(path, "sample.avpi")
	// fmt.Println(fullpath)

	fullpath := "sample.avpi"

	l.ReadPackage(fullpath)
}
