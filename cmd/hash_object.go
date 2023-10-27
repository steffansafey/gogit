package cmd

import (
	"crypto/sha1"
	"fmt"
	"os"

	"gogit/fs"

	"github.com/spf13/cobra"
)

func HashObject(path string) {
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

var hashObjectCommand = &cobra.Command{
	Use:   "hash-object",
	Short: "Compute object ID and creates a blob from a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Create the `objects` directory if it doesn't exist
		fs.CreateDir(cwd + "/.gogit/objects")

		// Hash the file
		HashObject(args[0])
	},
}

func init() {
	RootCmd.AddCommand(hashObjectCommand)
}
