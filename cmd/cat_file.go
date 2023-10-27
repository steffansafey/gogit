package cmd

import (
	"gogit/data"

	"github.com/spf13/cobra"
)

var catFileCommand = &cobra.Command{
	Use:   "cat-file",
	Short: "Cats a file if it exists in the objects directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bytes := data.ReadObject(args[0])
		print(string(bytes))
	},
}

func init() {
	RootCmd.AddCommand(catFileCommand)
}
