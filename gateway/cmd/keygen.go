package cmd

import (
	"encoding/base64"
	"fmt"
	"log/slog"

	"github.com/MakeNowJust/heredoc"
	"github.com/gtank/cryptopasta"
	"github.com/spf13/cobra"
)

func GenEncryptionKeyCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "keygen",
		Short: "Generate encryption key",
		Long: heredoc.Doc(`
			Generate encryption key encoded as base64 for encrypting/decrypting connection config`),
		Example: heredoc.Doc(`
			$ conman keygen`),
		RunE: func(c *cobra.Command, args []string) error {
			key := cryptopasta.NewEncryptionKey()
			slog.Info(fmt.Sprintf("Encryption key: %s", base64.RawStdEncoding.EncodeToString(key[:])))
			return nil
		},
	}
	return c
}
