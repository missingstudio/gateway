package main

import (
	"log"

	"github.com/missingstudio/studio/backend/cmd"
)

func main() {
	cli := cmd.New()

	if err := cli.Execute(); err != nil {
		log.Fatalf("mobius finished with error: %v", err)
	}
}
