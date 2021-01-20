package internal

import (
	"fmt"
	"github.com/spf13/cobra"
)

const urlFlagName = "url"
const usernameFlagName = "username"
const passwordFlagName = "password"
const insecureFlagName = "insecure"
const outputFlagName = "output"

const jsonFormat outputFormat = "json"
const textFormat outputFormat = "text"

var outputFormats = []outputFormat{jsonFormat, textFormat}

type outputFormat string

func connectionFlags(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		cmd.PersistentFlags().StringP(urlFlagName, "x", "", "URL to pfsense xmlrpc endpoint")
		if err := cmd.MarkPersistentFlagRequired(urlFlagName); err != nil {
			panic(err)
		}
		cmd.PersistentFlags().StringP(usernameFlagName, "u", "", "Username for pfsense")
		if err := cmd.MarkPersistentFlagRequired(usernameFlagName); err != nil {
			panic(err)
		}
		cmd.PersistentFlags().StringP(passwordFlagName, "p", "", "Password for pfsense")
		if err := cmd.MarkPersistentFlagRequired(passwordFlagName); err != nil {
			panic(err)
		}
		cmd.PersistentFlags().BoolP(insecureFlagName, "i", false, "Disable SSL verification")
	}
}

func outputFormatFlags(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		cmd.PersistentFlags().StringP(outputFlagName, "o", "text", "Output format; one of: json|text")
	}
}

func getOutputFormat(cmd *cobra.Command) (outputFormat, error) {
	s, err := cmd.Flags().GetString(outputFlagName)
	if err != nil {
		return "", fmt.Errorf("failed to parse %s flag; %w", outputFlagName, err)
	}
	for _, f := range outputFormats {
		if string(f) == s {
			return f, nil
		}
	}
	return "", fmt.Errorf("passed value %s is not in the list of supported formats: %v", s, outputFormats)
}
