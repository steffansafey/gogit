package cmd

import (
	"gogit/ops"

	"github.com/spf13/cobra"
)

var hashObjectCommand = &cobra.Command{
	Use:   "hash-object",
	Short: "Compute object ID and creates a blob from a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Hash the file and write it to the objects directory
		ops.HashObject(args[0], ops.Blob)
	},
}

func init() {
	RootCmd.AddCommand(hashObjectCommand)
}
