package cmd

import (
	"gogit/ops"

	"github.com/spf13/cobra"
)

var catFileCommand = &cobra.Command{
	Use:   "cat-file",
	Short: "Cats a file if it exists in the objects directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bytes := ops.ReadObject(args[0], ops.Any)
		print(string(bytes))
	},
}

func init() {
	RootCmd.AddCommand(catFileCommand)
}
