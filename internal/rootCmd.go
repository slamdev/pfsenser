package internal

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

const rootCmdName = "pfsenser"

const verboseFlag = "verbose"

var RootCmd = &cobra.Command{
	Use: rootCmdName,
	Long: `CLI wrapper on top of PfSense XML-RPC API.
 Find more information at: https://github.com/slamdev/` + rootCmdName,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		verbose, err := cmd.Flags().GetBool(verboseFlag)
		if err != nil {
			return err
		}
		if !verbose {
			log.SetOutput(ioutil.Discard)
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().BoolP(verboseFlag, "v", false, "verbose output")
}

func ExecuteCmd() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
