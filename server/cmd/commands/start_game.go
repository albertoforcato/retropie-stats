package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/mitchellh/cli"

	"github.com/albertoforcato/retropie-stats/logger"
	"github.com/albertoforcato/retropie-stats/models"
)

type StartGameCommand struct{}

func (s *StartGameCommand) Help() string {
	return "\nusage: retropie-stats start-game system emulator game-path emulator-command\n"
}

func (s *StartGameCommand) Run(args []string) int {
	if len(args) != 4 {
		logger.Log.Errorf("Invalid number of arguments")
		return 1
	}
	logger.Log.Info("Saving starting game...")
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
		"http://localhost:9512/api/v1/start-game-entry",
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

func (s *StartGameCommand) Synopsis() string {
	return "\nInsert a new start_game entry\n"
}

func startGameCommandFactory() (cli.Command, error) {
	return new(StartGameCommand), nil
}
