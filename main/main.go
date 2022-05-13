//Author: Juan Daniel Rodríguez Aguilar
//Description: Main app that creates a service to consume the APIs

package main

// import packages
import (
	"Pokemon/ECHO"
	"Pokemon/GIN"
	"Pokemon/GORILLA"
	"Pokemon/constants"
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

func echoPokemon() {
	router := echo.New()

	router.GET(constants.AllPokemon(), ECHO.GetAllPokemon)
	router.GET(constants.ByRegion(), ECHO.GetPokemonByRegion)
	router.GET(constants.ById(), ECHO.SearchById)
	router.GET(constants.ByName(), ECHO.SearchByName)

	router.Logger.Fatal(router.Start(constants.Domain()))

}

func ginPokemon() {
	router := gin.Default()
	defer router.Run(constants.Domain())

	router.GET(constants.AllPokemon(), GIN.GetAllPokemon)
	router.GET(constants.ByRegion(), GIN.GetPokemonByRegion)
	router.GET(constants.ById(), GIN.SearchById)
	router.GET(constants.ByName(), GIN.SearchByName)

}

func gorillaPokemon() {

	const (
		byRegion = "/pokemonByRegion/{region}"
		byId     = "/pokemonById/{id}"
		byName   = "/pokemonByName/{name}"
	)
	router := mux.NewRouter()
	defer http.ListenAndServe(constants.Domain(), router)

	router.HandleFunc(constants.AllPokemon(), GORILLA.GetAllPokemon).Methods("GET")
	router.HandleFunc(byRegion, GORILLA.GetPokemonByRegion).Methods("GET")
	router.HandleFunc(byId, GORILLA.SearchById).Methods("GET")
	router.HandleFunc(byName, GORILLA.SearchByName).Methods("GET")

	fmt.Println("Gorilla is running")
}
