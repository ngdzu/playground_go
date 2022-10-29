package mypackage

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Zipread() {
	zipReader, err := zip.OpenReader("valid_plugin3.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	for _, file := range zipReader.Reader.File {
		// open file inside zip file just like a normal file
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		//extract file (copy file from zip to local directory)
		targetDir := "./"
		exactractedFilePath := filepath.Join(targetDir, file.Name)

		if file.FileInfo().IsDir() {
			log.Println("creating directory:", exactractedFilePath)
			os.MkdirAll(exactractedFilePath, file.Mode())
		} else { // is a file
			log.Println("extracting file: ", file.Name)
			outputFile, err := os.OpenFile(
				exactractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			// "Extract" the file by copying zipped file
			// contents to the output file
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}

		}
	}

}
