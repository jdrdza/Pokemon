//Author: Juan Daniel Rodríguez Aguilar
//Description: Main app that creates a service to consume the APIs

package main

// import packages
import (
	c "Pokemon/Config"
	frame "Pokemon/framework"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
)

// Main function that deploys a service depending on the Framework inside the Server method
func main() {

	server("gin")

}

// Functions that return a service
func server(framework string) {

	switch strings.ToUpper(framework) {
	case "GIN":
		ginPokemon()
	case "ECHO":
		echoPokemon()
	case "GORILLA":
		gorillaPokemon()
	default:
		fmt.Printf("The %v framework is not implemented", framework)

	}
}

func ginPokemon() {
	router := gin.Default()
	config := c.Conf
	fmk := &frame.Gin{G: &gin.Context{}}
	defer router.Run(config.Environment)

	router.POST(config.EpCopyPkmn, fmk.ExternalAPI)
	router.POST(config.EpNewPkmn, fmk.NewPokemon)
	router.GET(config.EpAllPkmn, fmk.AllPokemon)
	router.GET(config.EpByRegion, fmk.PokemonByRegion)
	router.GET(config.EpById, fmk.SearchById)
	router.GET(config.EpByName, fmk.SearchByName)

}

func echoPokemon() {
	router := echo.New()
	config := c.Conf
	fmk := &frame.Echo{}

	router.POST(config.EpCopyPkmn, fmk.ExternalAPI)
	router.POST(config.EpNewPkmn, fmk.NewPokemon)
	router.GET(config.EpAllPkmn, fmk.AllPokemon)
	router.GET(config.EpByRegion, fmk.PokemonByRegion)
	router.GET(config.EpById, fmk.SearchById)
	router.GET(config.EpByName, fmk.SearchByName)

	router.Logger.Fatal(router.Start(config.Environment))

}

func gorillaPokemon() {
	router := mux.NewRouter()
	config := c.Conf
	fmk := &frame.Gorilla{R: &http.Request{}}
	defer http.ListenAndServe(config.Environment, router)

	router.HandleFunc(config.EpCopyPkmn, fmk.ExternalAPI).Methods("POST")
	router.HandleFunc(config.EpNewPkmn, fmk.NewPokemon).Methods("POST")
	router.HandleFunc(config.EpAllPkmn, fmk.AllPokemon).Methods("GET")
	router.HandleFunc(config.EpByRegion, fmk.PokemonByRegion).Methods("GET")
	router.HandleFunc(config.EpById, fmk.SearchById).Methods("GET")
	router.HandleFunc(config.EpByName, fmk.SearchByName).Methods("GET")

	fmt.Println("Gorilla is running")

}
