package v1

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/albertoforcato/retropie-stats/models"
)

func Test_InsertStartGameEntry(t *testing.T) {
	entry := models.Entry{
		System:      "test_system",
		Emulator:    "test_emulator",
		EmulatorCmd: "test_emulator_cmd",
		Event:       models.StartGameEvent,
		GamePath:    "test_game_path",
	}

	db := models.NewDbInstance()
	err := InsertStartGameEntry(db, entry)

	require.NoError(t, err)
}
