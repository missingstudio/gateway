package main

import (
	"log"

	"github.com/missingstudio/ai/gateway/cmd"
)

func main() {
	// Load CLI configuration to connect with GRPC server
	cliConfig, err := cmd.LoadConfig()
	if err != nil {
		cliConfig = &cmd.Config{}
	}

	if err := cmd.New(cliConfig).Execute(); err != nil {
		log.Fatalf("AI Gateway finished with error: %v", err)
	}
}
