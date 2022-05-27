package framework

import (
	con "Pokemon/Config"
	m "Pokemon/methods"
	s "Pokemon/structs"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
)

type Echo struct {
	E *echo.Context
}

type Gin struct {
	G *gin.Context
}

type Gorilla struct {
	R *http.Request
}

// Copy pokemon by region from an external API
func (G Gin) ExternalAPI(c *gin.Context) {

	status, err := m.PokeAPI(c.Param("region"))
	if err != nil {
		G.errorResponse(c, status, err.Error())
		return
	}

	G.successResponse(c, status)

}

func (E Echo) ExternalAPI(c echo.Context) error {

	status, err := m.PokeAPI(c.Param("region"))
	if err != nil {
		E.errorResponse(c, status, err.Error())
		return err
	}

	E.successResponse(c, status)
	return err

}

func (Gor Gorilla) ExternalAPI(w http.ResponseWriter, r *http.Request) {

	status, err := m.PokeAPI(mux.Vars(r)["region"])
	if err != nil {
		Gor.errorResponse(w, status, err.Error())
		return
	}

	Gor.successResponse(w, status)

}

// Inserts a new pokemon if not exists
func (G Gin) NewPokemon(c *gin.Context) {
	var newPokemon s.Pokemon
	if err := c.Bind(&newPokemon); err != nil {
		log.Println(err.Error())
		status := http.StatusInternalServerError
		err = errors.New(con.Conf.ServerError)
		G.errorResponse(c, status, err.Error())
		return
	}

	status, err := m.PostNewPokemon(newPokemon)
	if err != nil {
		G.errorResponse(c, status, err.Error())
		return
	}

	G.successResponse(c, status)
}

func (E Echo) NewPokemon(c echo.Context) error {
	var newPokemon s.Pokemon
	if err := c.Bind(&newPokemon); err != nil {
		log.Println(err.Error())
		status := http.StatusInternalServerError
		err = errors.New(con.Conf.ServerError)
		E.errorResponse(c, status, err.Error())
		return err
	}

	status, err := m.PostNewPokemon(newPokemon)
	if err != nil {
		E.errorResponse(c, status, err.Error())
		return err
	}

	E.successResponse(c, status)
	return err
}

func (Gor Gorilla) NewPokemon(w http.ResponseWriter, r *http.Request) {
	var newPokemon s.Pokemon
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		status := http.StatusInternalServerError
		err = errors.New(con.Conf.ServerError)
		Gor.errorResponse(w, status, err.Error())
		return
	}
	json.Unmarshal(responseData, &newPokemon)

	status, err := m.PostNewPokemon(newPokemon)
	if err != nil {
		Gor.errorResponse(w, status, err.Error())
		return
	}

	Gor.successResponse(w, status)
}

// Get all Pokemon
func (G Gin) AllPokemon(c *gin.Context) {
	status, err := m.GetAllPokemon()
	if err != nil {
		G.errorResponse(c, status, err.Error())
		return
	}

	G.successResponse(c, status)
}

func (E Echo) AllPokemon(c echo.Context) error {
	status, err := m.GetAllPokemon()
	if err != nil {
		E.errorResponse(c, status, err.Error())
		return err
	}

	E.successResponse(c, status)
	return err
}

func (Gor Gorilla) AllPokemon(w http.ResponseWriter, r *http.Request) {
	status, err := m.GetAllPokemon()
	if err != nil {
		Gor.errorResponse(w, status, err.Error())
		return
	}

	Gor.successResponse(w, status)
}

// Get pokemon by region
func (G Gin) PokemonByRegion(c *gin.Context) {
	status, err := m.GetPokemonByRegion(c.Param("region"))
	if err != nil {
		G.errorResponse(c, status, err.Error())
		return
	}

	G.successResponse(c, status)
}

func (E Echo) PokemonByRegion(c echo.Context) error {
	status, err := m.GetPokemonByRegion(c.Param("region"))
	if err != nil {
		E.errorResponse(c, status, err.Error())
		return err
	}

	E.successResponse(c, status)
	return err
}

func (Gor Gorilla) PokemonByRegion(w http.ResponseWriter, r *http.Request) {
	status, err := m.GetPokemonByRegion(mux.Vars(r)["region"])
	if err != nil {
		Gor.errorResponse(w, status, err.Error())
		return
	}

	Gor.successResponse(w, status)
}

// Gets a pokemon by ID
func (G Gin) SearchById(c *gin.Context) {
	status, err := m.SearchById(c.Param("id"))
	if err != nil {
		G.errorResponse(c, status, err.Error())
		return
	}

	G.successResponse(c, status)
}

func (E Echo) SearchById(c echo.Context) error {
	status, err := m.SearchById(c.Param("id"))
	if err != nil {
		E.errorResponse(c, status, err.Error())
		return err
	}

	E.successResponse(c, status)
	return err
}

func (Gor Gorilla) SearchById(w http.ResponseWriter, r *http.Request) {
	status, err := m.SearchById(mux.Vars(r)["id"])
	if err != nil {
		Gor.errorResponse(w, status, err.Error())
		return
	}

	Gor.successResponse(w, status)
}

// Gets a pokemon by name
func (G Gin) SearchByName(c *gin.Context) {
	status, err := m.SearchByName(c.Param("name"))
	if err != nil {
		G.errorResponse(c, status, err.Error())
		return
	}

	G.successResponse(c, status)

}

func (E Echo) SearchByName(c echo.Context) error {
	status, err := m.SearchByName(c.Param("name"))
	if err != nil {
		E.errorResponse(c, status, err.Error())
		return err
	}

	E.successResponse(c, status)
	return err
}

func (Gor Gorilla) SearchByName(w http.ResponseWriter, r *http.Request) {
	status, err := m.SearchByName(mux.Vars(r)["name"])
	if err != nil {
		Gor.errorResponse(w, status, err.Error())
		return
	}

	Gor.successResponse(w, status)

}

// Response
func (G Gin) errorResponse(c *gin.Context, status int, errorResponse string) {
	var response s.ErrorResponse
	response.ErrorMessage = errorResponse

	c.IndentedJSON(status, response)
}

func (G Gin) successResponse(c *gin.Context, status int) {
	var response s.SuccessResponse
	response.Count = len(m.PokemonSlice)
	response.Pokemon = m.PokemonSlice

	c.IndentedJSON(status, response)
}

func (E Echo) errorResponse(c echo.Context, status int, errorResponse string) error {
	var response s.ErrorResponse
	response.ErrorMessage = errorResponse

	return c.JSON(status, response)
}

func (E Echo) successResponse(c echo.Context, status int) error {
	var response s.SuccessResponse
	response.Count = len(m.PokemonSlice)
	response.Pokemon = m.PokemonSlice

	return c.JSON(status, response)
}

func (Gor Gorilla) errorResponse(w http.ResponseWriter, status int, errorResponse string) {
	var response s.ErrorResponse
	response.ErrorMessage = errorResponse
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(status)
	w.Write(jsonResponse)
}

func (Gor Gorilla) successResponse(w http.ResponseWriter, status int) {
	var response s.SuccessResponse
	response.Count = len(m.PokemonSlice)
	response.Pokemon = m.PokemonSlice
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(status)
	w.Write(jsonResponse)
}
