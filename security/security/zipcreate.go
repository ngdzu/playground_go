package security

import (
	"archive/zip"
	"log"
	"os"
)

func zipcreate() {
	outfile, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	// create zip writer on top of file writer
	zipWriter := zip.NewWriter(outfile)
	var filesToArchive = []struct {
		Name, Body string
	}{
		{"file1.txt", "String contents of file"},
		{"file2.txt", "\x61\x62\x63\n"},
	}

	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fileWriter.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// clean up
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
