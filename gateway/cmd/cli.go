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
                                                  
                                                     
🌈 AI gateway is an open-source, lightweight and high-performance gateway
`

func New(cliConfig *Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "gateway <command> <subcommand> [flags]",
		Short:         "🌈 AI gateway is an open-source, lightweight and high-performance gateway",
		Long:          heredoc.Doc(Description),
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(ServerCommand())
	cmd.AddCommand(ConfigCommand())
	SetHelp(cmd)
	return cmd
}
