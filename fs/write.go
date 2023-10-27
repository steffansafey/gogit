package fs

import (
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
	err := os.Mkdir(path, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
