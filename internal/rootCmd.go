package internal

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

const rootCmdName = "golang-cli"

const verboseFlag = "verbose"

var rootCmd = &cobra.Command{
	Use:   rootCmdName,
	Short: rootCmdName + " Describe CLI tool purpose here",
	Long: rootCmdName + ` Describe CLI tool purpose here
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
	rootCmd.PersistentFlags().BoolP(verboseFlag, "v", false, "verbose output")
}

func ExecuteCmd() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
