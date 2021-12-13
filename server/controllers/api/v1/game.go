package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/albertoforcato/retropie-stats/api_helpers"
	"github.com/albertoforcato/retropie-stats/models"
	services "github.com/albertoforcato/retropie-stats/services/api/v1"
)

//GameList returns a list of games
func (ctrl *PublicController) GameList(c *gin.Context) {
	gameList, err := services.GameList(ctrl.db)
	if err != nil {
		api_helpers.ResponseWithError(c, http.StatusInternalServerError, err)
		return
	}

	api_helpers.Response(c, http.StatusOK, gameList)
}

//InsertNewGame inserts a new game into the database
func (ctrl *PublicController) InsertNewGame(c *gin.Context) {
	var game models.Game
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		api_helpers.ResponseWithError(c, http.StatusInternalServerError, err)
		return
	}

	if err = json.Unmarshal(data, &game); err != nil {
		api_helpers.ResponseWithError(c, http.StatusInternalServerError, err)
		return
	}

	if err = services.InsertNewGame(ctrl.db, game); err != nil {
		api_helpers.ResponseWithError(c, http.StatusInternalServerError, err)
		return
	}

	api_helpers.Response(c, http.StatusOK, nil)
}
