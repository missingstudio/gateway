package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	mscmd "github.com/missingstudio/studio/common/cmd"

	"github.com/spf13/cobra"
)

type Config struct {
	Host string `mapstructure:"host"`
}

func LoadConfig() (*Config, error) {
	var config Config

	cfg := mscmd.SetConfig("mobius")
	err := cfg.Load(&config)

	return &config, err
}

func ConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config <command>",
		Short: "Manage client configurations",
		Example: heredoc.Doc(`
			$ mobius config init
			$ mobius config list`),
	}

	cmd.AddCommand(configInitCommand())
	cmd.AddCommand(configListCommand())

	return cmd
}

func configInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize a new client configuration",
		Example: heredoc.Doc(`
			$ mobius config init
		`),
		Annotations: map[string]string{
			"group": "core",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := mscmd.SetConfig("mobius")

			if err := cfg.Init(&Config{}); err != nil {
				return err
			}

			fmt.Printf("config created: %v\n", cfg.File())
			return nil
		},
	}
}

func configListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List client configuration settings",
		Example: heredoc.Doc(`
			$ mobius config list
		`),
		Annotations: map[string]string{
			"group": "core",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := mscmd.SetConfig("mobius")

			data, err := cfg.Read()
			if err != nil {
				return ErrClientConfigNotFound
			}

			fmt.Println(data)
			return nil
		},
	}
	return cmd
}
