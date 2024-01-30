package cmd

import (
	"errors"

	"github.com/MakeNowJust/heredoc"
)

var (
	ErrClientConfigNotFound = errors.New(heredoc.Doc(`
		Mobius client config not found.

		Run "mobius config init" to initialize a new client config or
		Run "mobius help environment" for more information.
	`))
	ErrClientConfigHostNotFound = errors.New(heredoc.Doc(`
		Mobius client config "host" not found.

		Pass mobius server host with "--host" flag or 
		set host in mobius config.

		Run "mobius config <subcommand>" or
		"mobius help environment" for more information.
	`))
	ErrClientNotAuthorized = errors.New(heredoc.Doc(`
		Mobius auth error. Mobius requires an auth header.
		
		Run "mobius help auth" for more information.
	`))
)
