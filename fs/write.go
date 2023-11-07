package fs

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func CreateFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func CreateDir(path string) {
	// Create the `objects` directory if it doesn't exist
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		os.Mkdir(path, 0755)
	}
}

func WriteBlob(path string, bytes []byte) []byte {
	// write a blob to the given path, overwriting any existing file
	CreateFile(path)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(bytes)

	return bytes
}

func DeleteFile(path string) {
	fmt.Println("Deleting file " + path)
	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}
