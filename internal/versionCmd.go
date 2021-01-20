package internal

import (
	"github.com/spf13/cobra"
)

var version = "main"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + rootCmdName,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("%s %s", rootCmdName, version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
