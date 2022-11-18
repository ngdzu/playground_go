package security

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func SymbolicLink() {
	filename := "./testdata/symlink"

	fileinfo, err := os.Lstat(filename)
	if err != nil {
		log.Fatal(err)
	}

	if fileinfo.Mode()&os.ModeSymlink != 0 {
		realpath, err := filepath.EvalSymlinks(filename)
		if err == nil {
			fmt.Printf("real path: %s\n", realpath)
		}
	}

}
