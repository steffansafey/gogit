package cmd

import (
	"fmt"
	"os"

	"gogit/fs"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new repository",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Initializing empty repository in", cwd)
		fs.CreateDir(cwd + "/.gogit")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
