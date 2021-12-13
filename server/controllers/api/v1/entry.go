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

//InsertStartGameEntry inserts a new "start game" entry
func (ctrl *PublicController) InsertStartGameEntry(c *gin.Context) {
	var entry models.Entry
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		api_helpers.ResponseWithError(c, http.StatusInternalServerError, err)
		return
	}

	if err = json.Unmarshal(data, &entry); err != nil {
		api_helpers.ResponseWithError(c, http.StatusInternalServerError, err)
		return
	}

	if err = services.InsertStartGameEntry(ctrl.db, entry); err != nil {
		api_helpers.ResponseWithError(c, http.StatusInternalServerError, err)
		return
	}

	api_helpers.Response(c, http.StatusOK, nil)
}

//InsertEndGameEntry inserts a new "end game" entry
func (ctrl *PublicController) InsertEndGameEntry(c *gin.Context) {
	var entry models.Entry
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		api_helpers.ResponseWithError(c, http.StatusInternalServerError, err)
		return
	}

	if err = json.Unmarshal(data, &entry); err != nil {
		api_helpers.ResponseWithError(c, http.StatusInternalServerError, err)
		return
	}

	if err = services.InsertEndGameEntry(ctrl.db, entry); err != nil {
		api_helpers.ResponseWithError(c, http.StatusInternalServerError, err)
		return
	}

	api_helpers.Response(c, http.StatusOK, nil)
}
