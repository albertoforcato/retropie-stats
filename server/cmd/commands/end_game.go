package main

import (
	"github.com/mitchellh/cli"
)

type EndGameCommand struct{}

func (e EndGameCommand) Help() string {
	//TODO implement me
	panic("implement me")
}

func (e EndGameCommand) Run(args []string) int {
	//TODO implement me
	panic("implement me")
}

func (e EndGameCommand) Synopsis() string {
	//TODO implement me
	panic("implement me")
}

func endGameCommandFactory() (cli.Command, error) {
	return new(EndGameCommand), nil
}
