//Author: Juan Daniel Rodríguez Aguilar
/*Description: This module contains functions that return the pokemon id and name from a CSV file
**using the ECHO framework
 */

package ECHO

// import packages
import (
	"Pokemon/csv"
	"Pokemon/rowsOperations"
	"Pokemon/structs"
	"net/http"

	"github.com/labstack/echo"
)

// Declaration of global variables
var pokemonSlice = []structs.Pokemon{}

// Function that gets all pokemon from a CSV file
func GetAllPokemon(c echo.Context) error {
	pokemonSlice = []structs.Pokemon{}

	status := http.StatusOK

	rows, err := csv.ReadCSV()
	if err != nil {
		errorOnServer(c)
		return err

	}

	for _, row := range rows {
		appendSlice(row)
	}

	pokemonSlice = pokemonSlice[1:len(pokemonSlice)]

	if len(pokemonSlice) == 0 {
		status = http.StatusNotFound
		errorMessage := "There are no pokemon"
		errorResponse(c, status, errorMessage)
		return err
	}

	pokemonSlice = rowsOperations.SortSlice(pokemonSlice)
	return successResponse(c, status)

}

// Function that gets the pokemon by region from a CSV file
func GetPokemonByRegion(c echo.Context) error {
	pokemonSlice = []structs.Pokemon{}
	status := http.StatusOK

	rows, err := csv.ReadCSV()

	if err != nil {
		errorOnServer(c)
		return err
	}
	region := c.Param("region")

	for _, row := range rows {
		if row[2] == region {
			appendSlice(row)
		}
	}

	if len(pokemonSlice) == 0 {
		status = http.StatusNotFound
		errorMessage := "There are no rows in the " + region + " region"
		errorResponse(c, status, errorMessage)
		return err

	}

	pokemonSlice = rowsOperations.SortSlice(pokemonSlice)
	return successResponse(c, status)
}

// Function that gets the pokemon by id from a CSV file
func SearchById(c echo.Context) error {
	pokemonSlice = []structs.Pokemon{}
	status := http.StatusOK

	rows, err := csv.ReadCSV()

	if err != nil {
		errorOnServer(c)
		return err
	}

	id := c.Param("id")
	exists, index := rowsOperations.ExistsRow(rows, id, "id")

	if !exists {
		status = http.StatusNotFound
		errorMessage := "The id " + id + " does not exist"

		errorResponse(c, status, errorMessage)
		return err

	}
	appendSlice(rows[index])
	return successResponse(c, status)

}

// Function that gets the pokemon by name from a CSV file
func SearchByName(c echo.Context) error {
	pokemonSlice = []structs.Pokemon{}
	status := http.StatusOK

	rows, err := csv.ReadCSV()

	if err != nil {
		errorOnServer(c)
		return err
	}

	name := c.Param("name")

	exists, index := rowsOperations.ExistsRow(rows, name, "name")

	if !exists {
		status = http.StatusNotFound
		errorMessage := "The name " + name + " does not exist"

		errorResponse(c, status, errorMessage)
		return err

	}

	appendSlice(rows[index])
	return successResponse(c, status)

}

// Function that gets a specific region
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

func appendSlice(row []string) {
	pokemon := rowsOperations.AddRowToSlice(row)
	pokemonSlice = append(pokemonSlice, pokemon)

}

// Function that returns an 500 error
func errorOnServer(c echo.Context) {
	status := http.StatusInternalServerError
	errorMessage := "There is an error on the server, contact the admin"
	errorResponse(c, status, errorMessage)

}

// Function that returns an message to the client
func errorResponse(c echo.Context, status int, errorResponse string) error {
	var response structs.ErrorResponse
	response.ErrorMessage = errorResponse

	return c.JSON(status, response)
}

// Function that returns a pokemon slice
func successResponse(c echo.Context, status int) error {
	var response structs.SuccessResponse
	response.Count = len(pokemonSlice)
	response.Pokemon = pokemonSlice

	return c.JSON(status, response)
}
