package main

import (
	"net/http"

	"github.com/mitchellh/cli"
)

type StartGameCommand struct{}

func (s StartGameCommand) Help() string {
	//TODO implement me
	panic("implement me")
}

func (s StartGameCommand) Run(args []string) int {
	//TODO implement me
	_, err := http.Post("http://localhost:9512/api/v1/insert-new-entry", "application/json", nil)
	if err != nil {
		return 1
	}
	panic("implement me")
}

func (s StartGameCommand) Synopsis() string {
	//TODO implement me
	panic("implement me")
}

func startGameCommandFactory() (cli.Command, error) {
	return new(StartGameCommand), nil
}
