package cmd

import (
	"fmt"
	"gogit/data"
	"strings"

	"github.com/spf13/cobra"
)

var readTreeCommand = &cobra.Command{
	Use:   "read-tree",
	Short: "Reads the tree object with the given oid",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		readTree(args[0])
	},
}

// Parse a tree, returning a list of TreeEntry objects
func getTreeEntries(tree_oid string) []TreeEntry {
	if len(tree_oid) != 40 {
		fmt.Printf("fatal: not a valid object name %s\n", tree_oid)
		return nil
	}

	tree := data.ReadObject(tree_oid, "tree")
	if tree == nil {
		return nil
	}
	tree_str := string(tree)

	// each line in the tree object is a tree entry
	// each tree entry is of the form:
	// <file_type> <name> <oid>
	// where <file_type> is either "blob" or "tree"

	entries := []TreeEntry{}
	for _, line := range strings.Split(tree_str, "\n") {
		components := strings.Split(line, " ")
		entries = append(entries, TreeEntry{data.FileType(components[0]), components[1], components[2]})
	}

	return entries
}

// Recurse an entire tree object and return a map of the form: {<name>: <oid>}
// where <name> is the path of the file, and <oid> is the oid of the object.
func buildFileMap(tree_oid string, base_path string) map[string]string {
	entries := getTreeEntries(tree_oid)
	tree := map[string]string{}

	for _, entry := range entries {
		if entry.file_type == data.Blob {
			tree[base_path+"/"+entry.name] = entry.oid
		} else if entry.file_type == data.Tree {
			// We recursively call getTree() on the tree objects.
			subtree := buildFileMap(entry.oid, base_path+"/"+entry.name)
			for k, v := range subtree {
				tree[k] = v
			}
		}
	}
	return tree
}

// Read the tree object with the given oid and write the files to the working directory
func readTree(oid string) {
	entries := buildFileMap(oid, ".")

	for k, v := range entries {
		data.WriteBlob(k, v)
	}

	fmt.Println(oid)
}

func init() {
	RootCmd.AddCommand(readTreeCommand)
}
