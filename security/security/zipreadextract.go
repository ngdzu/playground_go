package security

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Zipread() {
	filename := "valid_plugin3.zip"
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	//extract file (copy file from zip to local directory)
	targetDir := filepath.Join("./", filepath.Base(filename)[:len(filename)-len(filepath.Ext(filename))])
	for _, file := range zipReader.Reader.File {
		// open file inside zip file just like a normal file
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		if _, err := os.Stat(targetDir); err != nil {
			if os.IsNotExist(err) {
				os.MkdirAll(targetDir, os.ModePerm)
			}
		}

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
