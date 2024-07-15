package fileparser

import (
	"io"
	"log"
	"os"
)

func (f *parser) Readfile(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
