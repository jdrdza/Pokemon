package control

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jdrdza/pokemon/method"

	"github.com/gin-gonic/gin"
)

type ResponsePokeAPI struct {
	Resp []PokeAPI `json:"results"`
}

type PokeAPI struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
}
type error struct {
	Message string `json:"response"`
}

type success struct {
	Count   int         `json:"count"`
	Pokemon interface{} `json:"pokemon"`
}

type controller struct {
	M method.Method
}

type Controller interface {
	PokeAPI(c *gin.Context)
	NewPokemon(c *gin.Context)
	AllPokemon(c *gin.Context)
	PokemonByRegion(c *gin.Context)
	PokemonById(c *gin.Context)
	PokemonByName(c *gin.Context)
	Items(c *gin.Context)
}

func NewController(M method.Method) *controller {
	return &controller{M: M}
}

func (con controller) PokeAPI(c *gin.Context) {
	region := c.Query("region")
	var newPokemon ResponsePokeAPI
	var Rows [][]string

	offset, limit := getRegion(region)

	if offset == -1 || limit == -1 {
		response := error{Message: "the " + region + " region does not exists"}
		c.IndentedJSON(http.StatusNotFound, response)
		return
	}

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/?offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit))
	if err != nil {
		log.Println(err.Error())
		response := error{Message: "There was an internal error, contact the admin"}
		c.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		response := error{Message: "There was an internal error, contact the admin"}
		c.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	json.Unmarshal(responseData, &newPokemon)

	for _, rows := range newPokemon.Resp {

		url := strings.Split(rows.Url, "/")
		id := url[len(url)-2]

		row := [][]string{{id, rows.Name, region}}
		Rows = append(Rows, row...)
	}
	pokemonSlice, err := con.M.PokeAPI(Rows)
	if err != nil {
		response := error{Message: "There is an internal error, Contact the admin"}
		c.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	if len(pokemonSlice) == 0 {
		response := error{Message: "There are no new pokemon"}
		c.IndentedJSON(http.StatusNotFound, response)
		return
	}
	response := success{Count: len(pokemonSlice), Pokemon: pokemonSlice}
	c.IndentedJSON(http.StatusOK, response)

}

func (con controller) NewPokemon(c *gin.Context) {
	var newPokemon pokemon
	if err := c.Bind(&newPokemon); err != nil {
		log.Println(err.Error())
		response := error{Message: "There was an internal error, contact the admin"}
		c.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	pokemonSlice, err := con.M.PostNewPokemon(newPokemon.ID, newPokemon.Name, newPokemon.Region)
	if err != nil {
		if err.Error() == "The file could not be saved" {
			response := error{Message: "There is an internal error, Contact the admin"}
			c.IndentedJSON(http.StatusInternalServerError, response)
			return
		}
		response := error{Message: err.Error()}
		c.IndentedJSON(http.StatusOK, response)
		return
	}
	response := success{Count: len(pokemonSlice), Pokemon: pokemonSlice}
	c.IndentedJSON(http.StatusOK, response)
}

func (con controller) AllPokemon(c *gin.Context) {
	pokemonSlice := con.M.GetAllPokemon()

	pokemonSlice = pokemonSlice[1:]

	if len(pokemonSlice) == 0 {
		response := error{Message: "There are no Pokemon"}
		c.IndentedJSON(http.StatusNotFound, response)
		return
	}

	response := success{Count: len(pokemonSlice), Pokemon: pokemonSlice}
	c.IndentedJSON(http.StatusOK, response)

}

func (con controller) PokemonByRegion(c *gin.Context) {
	region := c.Query("region")
	pokemonSlice := con.M.GetPokemonByRegion(region)

	if len(pokemonSlice) == 0 {
		response := error{Message: "There are no rows in the " + region + " region"}
		c.IndentedJSON(http.StatusNotFound, response)
		return
	}

	response := success{Count: len(pokemonSlice), Pokemon: pokemonSlice}
	c.IndentedJSON(http.StatusOK, response)

}

func (con controller) PokemonById(c *gin.Context) {
	id := c.Query("id")
	pokemonSlice := con.M.GetPokemonById(id)

	if len(pokemonSlice) == 0 {
		response := error{Message: "The id " + id + " does not exist"}
		c.IndentedJSON(http.StatusNotFound, response)
		return
	}

	response := success{Count: len(pokemonSlice), Pokemon: pokemonSlice}
	c.IndentedJSON(http.StatusOK, response)

}

func (con controller) PokemonByName(c *gin.Context) {
	name := c.Query("name")
	pokemonSlice := con.M.GetPokemonByName(name)

	if len(pokemonSlice) == 0 {
		response := error{Message: "The name " + name + " does not exist"}
		c.IndentedJSON(http.StatusNotFound, response)
		return
	}

	response := success{Count: len(pokemonSlice), Pokemon: pokemonSlice}
	c.IndentedJSON(http.StatusOK, response)

}

func (con controller) Items(c *gin.Context) {
	types := c.Query("type")

	if types == "odd" {
		pokemonSlice := con.M.GetOdd()

		if len(pokemonSlice) == 0 {
			response := error{Message: "There are no odd items"}
			c.IndentedJSON(http.StatusNotFound, response)
			return
		}

		response := success{Count: len(pokemonSlice), Pokemon: pokemonSlice}
		c.IndentedJSON(http.StatusOK, response)
		return
	}

	if types == "even" {
		pokemonSlice := con.M.GetEven()
		pokemonSlice = pokemonSlice[1:]

		if len(pokemonSlice) == 0 {
			response := error{Message: "There are no even items"}
			c.IndentedJSON(http.StatusNotFound, response)
			return
		}

		response := success{Count: len(pokemonSlice), Pokemon: pokemonSlice}
		c.IndentedJSON(http.StatusOK, response)
		return
	}

}

func getRegion(region string) (int, int) {
	var offset = -1
	var limit = -1

	switch region {
	case "kanto":
		offset = 0
		limit = 151

	case "johto":
		offset = 151
		limit = 100

	case "hoenn":
		offset = 251
		limit = 135

	case "sinnoh":
		offset = 386
		limit = 107

	case "teselia":
		offset = 493
		limit = 156

	case "kalos":
		offset = 649
		limit = 81

	case "alola":
		offset = 721
		limit = 88

	case "galar":
		offset = 809
		limit = 89
	}

	return offset, limit
}
