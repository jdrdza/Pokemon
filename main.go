package main

import (
	"log"

	"github.com/jdrdza/pokemon/router"
)

func main() {
	router, err := router.Initialise()
	if err != nil {
		log.Println("The server could not be started")
		return
	}

	router.Routers()

}
