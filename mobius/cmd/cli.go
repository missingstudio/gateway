package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

const Description = `


███▄ ▄███▓ ▒█████   ▄▄▄▄    ██▓ █    ██   ██████ 
▓██▒▀█▀ ██▒▒██▒  ██▒▓█████▄ ▓██▒ ██  ▓██▒▒██    ▒ 
▓██    ▓██░▒██░  ██▒▒██▒ ▄██▒██▒▓██  ▒██░░ ▓██▄   
▒██    ▒██ ▒██   ██░▒██░█▀  ░██░▓▓█  ░██░  ▒   ██▒
▒██▒   ░██▒░ ████▓▒░░▓█  ▀█▓░██░▒▒█████▓ ▒██████▒▒
░ ▒░   ░  ░░ ▒░▒░▒░ ░▒▓███▀▒░▓  ░▒▓▒ ▒ ▒ ▒ ▒▓▒ ▒ ░
░  ░      ░  ░ ▒ ▒░ ▒░▒   ░  ▒ ░░░▒░ ░ ░ ░ ░▒  ░ ░
░      ░   ░ ░ ░ ▒   ░    ░  ▒ ░ ░░░ ░ ░ ░  ░  ░  
       ░       ░ ░   ░       ░     ░           ░  
                          ░                       
                                                  
                                                     
🌈 Mobius is an open-source, lightweight, high-performance ai studio gateway
`

func New(cliConfig *Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "mobius <command> <subcommand> [flags]",
		Short:         "🌈 Mobius is an open-source, lightweight, high-performance ai studio gateway",
		Long:          heredoc.Doc(Description),
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(ServerCommand())
	cmd.AddCommand(ConfigCommand())
	SetHelp(cmd)
	return cmd
}
