package fs

import (
	"log"
	"os"

	"path/filepath"
)

func RecursivelyListFilesInDir(dir string) []string {
	// Recursively list all files in a directory
	// Returns a list of file paths
	files := []string{}

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return files
}
