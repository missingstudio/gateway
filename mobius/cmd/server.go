package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/missingstudio/studio/backend/config"

	"github.com/spf13/cobra"
)

func ServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "Server management",
		Long:    "Server management commands.",
		Example: heredoc.Doc(`
			$ mobius server start
		`),
	}

	cmd.AddCommand(serverInitCommand())
	cmd.AddCommand(serverStartCommand())
	return cmd
}

func serverInitCommand() *cobra.Command {
	var configFile string
	c := &cobra.Command{
		Use:   "init",
		Short: "Initialize server",
		Long: heredoc.Doc(`
			Initializing server. Creating a sample of frontier server config.
			Default: ./config.yaml
		`),
		Example: "frontier server init",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.Init(configFile); err != nil {
				return err
			}

			fmt.Printf("server config created: %v\n", configFile)
			return nil
		},
	}

	c.Flags().StringVarP(&configFile, "output", "o", "./config.yaml", "Output config file path")
	return c
}

func serverStartCommand() *cobra.Command {
	var configFile string

	c := &cobra.Command{
		Use:     "start",
		Short:   "Start server and proxy default on port 8080",
		Example: "frontier server start",
		RunE: func(cmd *cobra.Command, args []string) error {
			appConfig, err := config.Load(configFile)
			if err != nil {
				panic(err)
			}
			return Serve(appConfig)
		},
	}

	c.Flags().StringVarP(&configFile, "config", "c", "", "config file path")
	return c
}
