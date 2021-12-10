package routers

import (
	"github.com/gin-gonic/gin"

	apiControllerV1 "github.com/albertoforcato/retropie-stats/controllers/api/v1"
	"github.com/albertoforcato/retropie-stats/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// API routes for version 1
	v1 := r.Group("/api/v1")
	v1.Use(middlewares.Middlewares())
	{
		v1.GET("game-list", apiControllerV1.GameList)
	}
	return r
}
