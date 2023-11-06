package data

import (
	"crypto/sha1"
	"fmt"
	"os"
	"strings"

	"gogit/fs"
)

// Type of file to hash
type FileType string

const (
	Blob FileType = "blob"
	Tree FileType = "tree"
	Any  FileType = "any"
)

func HashObject(path string, filetype FileType) string {
	if filetype == Any {
		fmt.Println("A specific filetype must be specified")
		os.Exit(1)
	}

	// Hash the file
	file_bytes, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return HashBytes(file_bytes, filetype)
}

func HashBytes(file_bytes []byte, filetype FileType) string {

	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Create the `objects` directory if it doesn't exist
	fs.CreateDir(cwd + "/.gogit/objects")
	// Append the type of file to the beginning of the file_bytes
	file_bytes = append([]byte(filetype+"\x00"), file_bytes...)

	// Hash the file
	hash := sha1.New()
	hash.Write(file_bytes)

	hash_string := fmt.Sprintf("%x", hash.Sum(nil))

	// Write the file to the objects directory
	fs.WriteBlob(".gogit/objects/"+hash_string, file_bytes)

	fmt.Println("wrote obj " + hash_string + " " + string(filetype))
	return hash_string
}

func ReadObject(hash string, expected_type FileType) []byte {
	// Read the file_bytes into memory
	file_bytes, err := os.ReadFile(".gogit/objects/" + hash)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Check the type of the file (search for the first null byte)
	null_byte_index := 0
	for i, b := range file_bytes {
		if b == 0 {
			null_byte_index = i
			break
		}
	}

	// Check the type of the file
	filetype := FileType(file_bytes[:null_byte_index])

	if expected_type != Any && filetype != expected_type {
		fmt.Println("Expected file type `", expected_type, "` but got `", filetype, "`")
	}

	// Return the file_bytes without the type
	return file_bytes[null_byte_index+1:]
}

func WriteBlob(path string, oid string) {
	// Read the file_bytes into memory
	file_bytes := ReadObject(oid, Blob)

	// Create the directory if it doesn't exist
	fs.CreateDir(path[:len(path)-len(path[strings.LastIndex(path, "/"):])])

	fs.WriteBlob(path, file_bytes)
}
