package data

import (
	"crypto/sha1"
	"fmt"
	"os"

	"gogit/fs"
)

func HashObject(path string) {
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create the `objects` directory if it doesn't exist
	fs.CreateDir(cwd + "/.gogit/objects")

	// Hash the file
	// Read the file_bytes into memory
	file_bytes, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Hash the file
	hash := sha1.New()
	hash.Write(file_bytes)

	hash_string := fmt.Sprintf("%x", hash.Sum(nil))

	// Write the file to the objects directory
	fs.WriteBlob(".gogit/objects/"+hash_string, file_bytes)

	fmt.Println(hash_string)
}

func ReadObject(hash string) []byte {
	// Read the file_bytes into memory
	file_bytes, err := os.ReadFile(".gogit/objects/" + hash)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return file_bytes
}
