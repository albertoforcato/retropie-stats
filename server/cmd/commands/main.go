package main

import (
	"os"

	"github.com/mitchellh/cli"

	"github.com/albertoforcato/retropie-stats/logger"
)

func main() {
	c := cli.NewCLI("retropie-stats", "0.1.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"start-game": startGameCommandFactory,
		"end-game":   endGameCommandFactory,
	}
	exitStatus, err := c.Run()
	if err != nil {
		logger.Log.Errorf("Error: %s", err)
	}

	os.Exit(exitStatus)
}
