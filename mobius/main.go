package main

import (
	"log"

	"github.com/missingstudio/studio/backend/cmd"
)

func main() {
	cliConfig, err := cmd.LoadConfig()
	if err != nil {
		cliConfig = &cmd.Config{}
	}

	if err := cmd.New(cliConfig).Execute(); err != nil {
		log.Fatalf("mobius finished with error: %v", err)
	}
}
