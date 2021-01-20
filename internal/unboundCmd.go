package internal

import (
	"fmt"
	"github.com/slamdev/pfsenser/pkg/unbound"
	"github.com/slamdev/pfsenser/pkg/xmlrpc"
	"github.com/spf13/cobra"
)

var client xmlrpc.PfsenseClient

var unboundCmd = &cobra.Command{
	Use:              "unbound",
	Short:            "Operate with Unbound DNS resolver",
	TraverseChildren: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		url, err := cmd.Flags().GetString(urlFlagName)
		username, err := cmd.Flags().GetString(usernameFlagName)
		password, err := cmd.Flags().GetString(passwordFlagName)
		insecure, err := cmd.Flags().GetBool(insecureFlagName)
		if err != nil {
			return fmt.Errorf("failed to parse flag; %w", err)
		}
		client, err = xmlrpc.NewPfsenseClient(url, username, password, insecure)
		return err
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		if err := client.Close(); err != nil {
			return fmt.Errorf("failed to close pfsense client; %w", err)
		}
		return nil
	},
}

var createDescription string
var createAliases string

var unboundCreateCmd = &cobra.Command{
	Use:   "create [hostname] [IP address]",
	Short: "Create host override",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return unbound.Create(client, unbound.Host{
			Name:        args[0],
			IP:          args[1],
			Description: createDescription,
			Aliases:     createAliases,
		})
	},
}

var unboundDeleteCmd = &cobra.Command{
	Use:   "delete [hostname]",
	Short: "Delete host override",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return unbound.Delete(client, args[0])
	},
}

var unboundGetCmd = &cobra.Command{
	Use:   "get [hostname]",
	Short: "Get host override",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		host, exists, err := unbound.Get(client, args[0])
		if err != nil {
			return err
		}
		if !exists {
			cmd.Printf("%s host is not found", args[0])
			return nil
		}
		fmt.Printf("%+v\n", host)
		return nil
	},
}

var unboundListCmd = &cobra.Command{
	Use:   "list",
	Short: "List host overrides",
	RunE: func(cmd *cobra.Command, args []string) error {
		hosts, err := unbound.List(client)
		if err != nil {
			return err
		}
		if len(hosts) == 0 {
			cmd.Printf("no host overrides found", args[0])
			return nil
		}
		for _, host := range hosts {
			fmt.Printf("%+v\n", host)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(unboundCmd)
	connectionFlags(unboundCmd)
	unboundCmd.AddCommand(unboundCreateCmd)
	unboundCreateCmd.Flags().StringVar(&createDescription, "description", "", "host rule description")
	unboundCreateCmd.Flags().StringVar(&createAliases, "aliases", "", "host rule aliases")
	unboundCmd.AddCommand(unboundDeleteCmd)
	unboundCmd.AddCommand(unboundGetCmd)
	unboundCmd.AddCommand(unboundListCmd)
}
