package cmd

import (
	"fmt"
	"os"

	"gogit/data"

	"github.com/spf13/cobra"
)

type TreeEntry struct {
	file_type data.FileType
	oid       string
	name      string
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

		oid := writeTree(path)

		fmt.Println(oid)
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

		path := directory + "/" + entry.Name()
		if data.IsIgnored(path) {
			continue
		}

		if entry.IsDir() {
			oid := writeTree(path)
			tree_entries = append(tree_entries, TreeEntry{data.Tree, oid, entry.Name()})
		} else {
			oid := data.HashObject(path, data.Blob)
			tree_entries = append(tree_entries, TreeEntry{data.Blob, oid, entry.Name()})
		}
	}

	// Write the tree_str object to a file, ensuring no newlines after the last line
	tree_str := ""
	for i, tree_entry := range tree_entries {
		if i == len(tree_entries)-1 {
			tree_str += fmt.Sprintf("%s %s %s", tree_entry.file_type, tree_entry.oid, tree_entry.name)
		} else {
			tree_str += fmt.Sprintf("%s %s %s\n", tree_entry.file_type, tree_entry.oid, tree_entry.name)
		}
	}

	oid := data.HashBytes([]byte(tree_str), data.Tree)

	return oid
}

func init() {
	RootCmd.AddCommand(writeTreeCommand)
}
