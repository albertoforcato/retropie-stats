package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/mitchellh/cli"

	"github.com/albertoforcato/retropie-stats/logger"
	"github.com/albertoforcato/retropie-stats/models"
)

type EndGameCommand struct{}

func (e *EndGameCommand) Help() string {
	return "\nusage: retropie-stats end-game system emulator game-path emulator-command\n"
}

func (e *EndGameCommand) Run(args []string) int {
	logger.Log.Info("Saving ending game...")
	jsonData, err := json.Marshal(models.Entry{
		System:      args[0],
		Emulator:    args[1],
		GamePath:    args[2],
		EmulatorCmd: args[3],
	})
	if err != nil {
		logger.Log.Error(err)
		return 1
	}
	res, err := http.Post(
		"http://localhost:9512/api/v1/end-game-entry",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		logger.Log.Error(err)
		return 1
	}
	logger.Log.Info(res.Status)
	return 0
}

func (e *EndGameCommand) Synopsis() string {
	return "\nInsert a new end_game entry\n"
}

func endGameCommandFactory() (cli.Command, error) {
	return new(EndGameCommand), nil
}
