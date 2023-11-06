package cmd

import (
	"fmt"
	"os"

	"gogit/data"

	"github.com/spf13/cobra"
)

var writeTreeCommand = &cobra.Command{
	Use:   "write-tree",
	Short: "Writes the current state of the index to the objects directory",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) == 1 {
			path = args[0]
		}

		writeTree(path)
	},
}

func writeTree(directory string) {
	// List the entries and directories in the current directory
	entries, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Iterate over the files and directories
	for _, entry := range entries {
		if isIgnored(entry) {
			continue
		}

		path := directory + "/" + entry.Name()
		if entry.IsDir() {
			writeTree(path)
		} else {
			data.HashObject(path, data.Blob)
			fmt.Println("hashed " + path)
		}
	}
}

func isIgnored(entry os.DirEntry) bool {
	ignored_entries := []string{".gogit"}

	for _, ignored_entry := range ignored_entries {
		if entry.Name() == ignored_entry {
			return true
		}
	}
	return false
}

func init() {
	RootCmd.AddCommand(writeTreeCommand)
}
