package main

import (
	"github.com/albertoforcato/retropie-stats/models"
	"github.com/albertoforcato/retropie-stats/routers"
)

const serverPort = ":9512"

func main() {
	//Initialize database
	db := models.NewDbInstance()
	//Initialize the server.
	r := routers.SetupRouter(db)
	r.Run(serverPort)
}
