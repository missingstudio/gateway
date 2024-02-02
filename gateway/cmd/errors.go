package cmd

import (
	"errors"

	"github.com/MakeNowJust/heredoc"
)

var (
	ErrClientConfigNotFound = errors.New(heredoc.Doc(`
		gateway client config not found.

		Run "gateway config init" to initialize a new client config or
		Run "gateway help environment" for more information.
	`))
	ErrClientConfigHostNotFound = errors.New(heredoc.Doc(`
		gateway client config "host" not found.

		Pass gateway server host with "--host" flag or 
		set host in gateway config.

		Run "gateway config <subcommand>" or
		"gateway help environment" for more information.
	`))
	ErrClientNotAuthorized = errors.New(heredoc.Doc(`
		gateway auth error. gateway requires an auth header.
		
		Run "gateway help auth" for more information.
	`))
)
