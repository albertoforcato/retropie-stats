package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/albertoforcato/retropie-stats/api_helpers"
	services "github.com/albertoforcato/retropie-stats/services/api/v1"
)

func GameList(c *gin.Context) {
	var gameServices services.GameService

	gameList, err := gameServices.GameList()
	if err != nil {
		api_helpers.ResponseWithError(c, http.StatusInternalServerError, err)
		return
	}

	api_helpers.Response(c, http.StatusOK, gameList)
}
