//Author: Juan Daniel Rodríguez Aguilar
/*Description: This module contains functions that return the pokemon id and name from a CSV file
**using the Gorilla framework
 */

package GORILLA

// import packages
import (
	"Pokemon/csv"
	"Pokemon/rowsOperations"
	"Pokemon/structs"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Declaration of global variables
var pokemonSlice = []structs.Pokemon{}

// Function that gets all pokemon from a CSV file
func GetAllPokemon(w http.ResponseWriter, r *http.Request) {
	pokemonSlice = []structs.Pokemon{}

	status := http.StatusOK

	rows, err := csv.ReadCSV()
	if err != nil {
		errorOnServer(w)
		return

	}

	for _, row := range rows {
		appendSlice(row)
	}

	pokemonSlice = pokemonSlice[1:len(pokemonSlice)]

	if len(pokemonSlice) == 0 {
		status = http.StatusNotFound
		errorMessage := "There are no pokemon"
		errorResponse(w, status, errorMessage)
		return
	}

	pokemonSlice = rowsOperations.SortSlice(pokemonSlice)
	successResponse(w, status)

}

// Function that gets the pokemon by region from a CSV file
func GetPokemonByRegion(w http.ResponseWriter, r *http.Request) {
	pokemonSlice = []structs.Pokemon{}
	status := http.StatusOK

	rows, err := csv.ReadCSV()

	if err != nil {
		errorOnServer(w)
		return
	}
	region := mux.Vars(r)["region"]

	for _, row := range rows {
		if row[2] == region {
			appendSlice(row)
		}
	}

	if len(pokemonSlice) == 0 {
		status = http.StatusNotFound
		errorMessage := "There are no rows in the " + region + " region"
		errorResponse(w, status, errorMessage)
		return

	}

	pokemonSlice = rowsOperations.SortSlice(pokemonSlice)
	successResponse(w, status)
}

// Function that gets the pokemon by id from a CSV file
func SearchById(w http.ResponseWriter, r *http.Request) {
	pokemonSlice = []structs.Pokemon{}
	status := http.StatusOK

	rows, err := csv.ReadCSV()

	if err != nil {
		errorOnServer(w)
		return
	}

	id := mux.Vars(r)["id"]
	exists, index := rowsOperations.ExistsRow(rows, id, "id")

	if !exists {
		status = http.StatusNotFound
		errorMessage := "The id " + id + " does not exist"

		errorResponse(w, status, errorMessage)
		return

	}
	appendSlice(rows[index])
	successResponse(w, status)

}

// Function that gets the pokemon by name from a CSV file
func SearchByName(w http.ResponseWriter, r *http.Request) {
	pokemonSlice = []structs.Pokemon{}
	status := http.StatusOK

	rows, err := csv.ReadCSV()

	if err != nil {
		errorOnServer(w)
		return
	}

	name := mux.Vars(r)["name"]

	exists, index := rowsOperations.ExistsRow(rows, name, "name")

	if !exists {
		status = http.StatusNotFound
		errorMessage := "The name " + name + " does not exist"

		errorResponse(w, status, errorMessage)
		return

	}

	appendSlice(rows[index])
	successResponse(w, status)

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
func errorOnServer(w http.ResponseWriter) {
	status := http.StatusInternalServerError
	errorMessage := "There is an error on the server, contact the admin"

	errorResponse(w, status, errorMessage)
}

// Function that returns an message to the client
func errorResponse(w http.ResponseWriter, status int, errorResponse string) {
	var response structs.ErrorResponse
	response.ErrorMessage = errorResponse
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(status)
	w.Write(jsonResponse)
}

// Function that returns a pokemon slice
func successResponse(w http.ResponseWriter, status int) {
	var response structs.SuccessResponse
	response.Count = len(pokemonSlice)
	response.Pokemon = pokemonSlice
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(status)
	w.Write(jsonResponse)
}
