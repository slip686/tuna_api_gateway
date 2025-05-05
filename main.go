package main

import (
	"TunaAPIGateway/config"
	"TunaAPIGateway/internal/api"
	"fmt"
	"log"
)

func main() {
	server := api.NewAPIServer(fmt.Sprintf(":%d", config.Config.API.APIPort))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
