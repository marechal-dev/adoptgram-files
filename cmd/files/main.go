package main

import (
	"fmt"
	"log"

	"github.com/marechal-dev/adoptgram-files/api"
	"github.com/marechal-dev/adoptgram-files/config"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	server := api.NewServer()
	serverAddressWithPort := fmt.Sprintf("%s:%s", api.ServerDefaultAddress, config.Port)

	server.SpinUp(serverAddressWithPort)
}
