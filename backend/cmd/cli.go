package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mobius <command> <subcommand> [flags]",
		Short: "🐦 Mobius is an open-source, lightweight, high-performance ai studio gateway",
		Long: heredoc.Doc(`
      Mobius is an open-source, lightweight, high-performance ai studio gateway.
    `),
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(ServerCommand())
	SetHelp(cmd)
	return cmd
}
