package cmd

import (
	"fmt"
	"os"

	"gogit/data"

	"github.com/spf13/cobra"
)

type TreeEntry struct {
	name      string
	oid       string
	file_type data.FileType
}

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

func writeTree(directory string) string {

	// List the entries and directories in the current directory
	entries, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tree_entries := []TreeEntry{}
	// Iterate over the files and directories
	for _, entry := range entries {
		if isIgnored(entry) {
			continue
		}

		path := directory + "/" + entry.Name()
		if entry.IsDir() {
			oid := writeTree(path)
			tree_entries = append(tree_entries, TreeEntry{entry.Name(), oid, data.Tree})
		} else {
			oid := data.HashObject(path, data.Blob)
			tree_entries = append(tree_entries, TreeEntry{entry.Name(), oid, data.Blob})
		}
	}

	// Write the tree_str object to a file
	tree_str := ""
	for _, tree_entry := range tree_entries {
		tree_str += fmt.Sprintf("%s %s %s\n", tree_entry.file_type, tree_entry.oid, tree_entry.name)
	}

	oid := data.HashBytes([]byte(tree_str), data.Tree)

	return oid
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
