package loader

import (
	"archive/zip"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Loader struct {
	Manifest string
}

// Read one package
// Take pointer or not?
// Will we change the internal of the Loader at all?
func (loader Loader) ReadPackage(path string) error {
	_, err := os.Stat(path)

	// if errors.Is(err, os.ErrNotExist) {
	// 	return errors.New("File doesn't exist")
	// }

	if err != nil {
		return err
	}

	// check file extension
	if filepath.Ext(path) != ".avpi" {
		return errors.New("bad file extension")
	}

	archive, err := zip.OpenReader(path)
	if err != nil {
		return err
	}

	// TODO We may not close this until the loader restart?
	// Which means returing the archive from this function?
	defer archive.Close()

	for _, f := range archive.File {
		// check for the existence of the manifest

		fmt.Printf("name: %v\n", f.FileHeader.Name)
	}

	return nil
}
