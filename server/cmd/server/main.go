package main

import (
	"github.com/albertoforcato/retropie-stats/routers"
)

const serverPort = ":9512"

func main() {
	// Initialize the server.
	r := routers.SetupRouter()
	r.Run(serverPort)
}
