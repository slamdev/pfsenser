package internal

import (
	"fmt"
	"github.com/slamdev/golang-cli/pkg"
	"github.com/spf13/cobra"
)

var length uint8

var randomTextCmd = &cobra.Command{
	Use:              "random-text",
	Short:            "Print a random text",
	Long:             "Generate a random text string and print it",
	TraverseChildren: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		str, err := pkg.RndString(uint(length))
		if err != nil {
			return err
		}
		fmt.Print(str)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(randomTextCmd)
	randomTextCmd.Flags().Uint8VarP(&length, "length", "l", 0, "text length")
	if err := randomTextCmd.MarkFlagRequired("length"); err != nil {
		panic(err)
	}
}
