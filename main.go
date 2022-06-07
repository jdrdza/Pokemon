package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jdrdza/pokemon/control"
	"github.com/jdrdza/pokemon/file"
	"github.com/jdrdza/pokemon/method"
)

var cont *control.Controller

type pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
}

type Task struct {
	Row           []string
	TaskProcessor func([]string)
}

func (t Task) Run() {
	t.TaskProcessor(t.Row)
}

func main() {
	cont, err := initialise()
	if err != nil {
		log.Println("The server could not be started")
		return
	}

	router := gin.Default()

	defer router.Run("localhost:8080")

	router.POST("/pokemon", func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		if len(query) > 1 {
			ctx.IndentedJSON(http.StatusNotFound, "The API only accpets one value per query")
			return
		}

		if query.Get("region") != "" {
			cont.PokeAPI(ctx)
		} else {
			cont.NewPokemon(ctx)
		}
		cont, err = initialise()
		if err != nil {
			log.Println("There was an error: " + err.Error())
			return
		}
	})

	router.GET("/pokemon", func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()

		switch {
		case query.Get("region") != "":
			cont.PokemonByRegion(ctx)

		case query.Get("name") != "":
			cont.PokemonByName(ctx)

		case query.Get("id") != "":
			cont.PokemonById(ctx)

		case query.Get("type") != "":
			cont.Items(ctx)
		default:
			cont.AllPokemon(ctx)
		}

	})

}

func initialise() (cont control.Controller, err error) {
	rows, err := file.NewFile()
	if err != nil {
		return nil, err
	}
	meth := method.NewMethod(rows)
	cont = control.NewController(meth)
	return cont, err
}

func appendSlice(row []string, pokemonSlice []pokemon) []pokemon {
	var pokemon pokemon

	pokemon.ID, _ = strconv.Atoi(row[0])
	pokemon.Name = row[1]
	pokemon.Region = row[2]

	pokemonSlice = append(pokemonSlice, pokemon)
	return pokemonSlice
}
