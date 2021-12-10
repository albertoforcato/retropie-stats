package main

import "github.com/mitchellh/cli"

type StartGameCommand struct{}

func (s StartGameCommand) Help() string {
	//TODO implement me
	panic("implement me")
}

func (s StartGameCommand) Run(args []string) int {
	//TODO implement me
	panic("implement me")
}

func (s StartGameCommand) Synopsis() string {
	//TODO implement me
	panic("implement me")
}

func startGameCommandFactory() (cli.Command, error) {
	return new(StartGameCommand), nil
}
