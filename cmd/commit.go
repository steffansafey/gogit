package cmd

import (
	"fmt"
	"os"
	"time"

	"gogit/data"

	"github.com/spf13/cobra"
)

// gogit commit -m "Initial commit"
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a new commit object",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// Create a new tree object from the current state of the index
		tree_oid := writeTree(".")

		// Get the commit message
		message, err := cmd.Flags().GetString("m")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Create the commit object
		commit_oid := createCommit(message, tree_oid)

		fmt.Println(commit_oid)

	},
}

func createCommit(message string, tree_oid string) string {
	time_string := time.Now().Format(time.RFC3339)
	author := "John Doe"
	commit := fmt.Sprintf("tree %s\nauthor %s\n%s\n\n%s\n", tree_oid, author, time_string, message)

	// Write the commit to the objects directory
	oid := data.HashBytes([]byte(commit), data.Commit)

	return oid
}

func init() {
	commitCmd.Flags().StringP("m", "m", "", "Commit message")
	RootCmd.AddCommand(commitCmd)
}
