package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	apiControllerV1 "github.com/albertoforcato/retropie-stats/controllers/api/v1"
	"github.com/albertoforcato/retropie-stats/middlewares"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	pCtrl := apiControllerV1.NewPublicController(db)

	// API routes for version 1
	v1 := r.Group("/api/v1")
	v1.Use(middlewares.Middlewares())
	{
		v1.GET("game-list", pCtrl.GameList)
		v1.POST("start-game-entry", pCtrl.InsertStartGameEntry)
		v1.POST("end-game-entry", pCtrl.InsertEndGameEntry)
	}
	return r
}
