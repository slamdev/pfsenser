package internal

import (
	"github.com/spf13/cobra"
)

const urlFlagName = "url"
const usernameFlagName = "username"
const passwordFlagName = "password"
const insecureFlagName = "insecure"

func connectionFlags(cmd *cobra.Command) {
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
