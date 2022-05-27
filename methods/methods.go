package methods

import (
	con "Pokemon/Config"
	"Pokemon/csv"
	ro "Pokemon/rowsOperations"
	s "Pokemon/structs"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var PokemonSlice = []s.Pokemon{}

func PokeAPI(region string) (status int, err error) {

	var newPokemon s.ResponsePokeAPI
	newRows := [][]string{}
	PokemonSlice = []s.Pokemon{}
	status = http.StatusCreated
	offset, limit := getRegion(region)

	if offset > -1 || limit > -1 {

		response, err := http.Get(con.Conf.ExternalAPI + "?offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit))
		if err != nil {
			log.Println(err.Error())
			status = http.StatusInternalServerError
			err = errors.New(con.Conf.ServerError)
			return status, err
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err.Error())
			status = http.StatusInternalServerError
			err = errors.New(con.Conf.ServerError)
			return status, err
		}

		json.Unmarshal(responseData, &newPokemon)

		for _, rows := range newPokemon.Resp {

			url := strings.Split(rows.Url, "/")
			id := url[len(url)-2]

			row := [][]string{{id, rows.Name, region}}
			newRows = append(newRows, row...)
		}

		PokemonSlice, err = csv.WriteCSV(newRows)
		if err != nil {
			status = http.StatusInternalServerError
			err = errors.New(con.Conf.ServerError)
			return status, err
		}

		if len(PokemonSlice) == 0 {
			status = http.StatusOK
			err = errors.New("There are no new pokemon")
			return status, err
		}

	} else {
		status = http.StatusNotFound
		err = errors.New("the " + region + " region does not exists")
		return status, err
	}
	return status, err

}

func PostNewPokemon(newPokemon s.Pokemon) (status int, err error) {
	PokemonSlice = []s.Pokemon{}
	status = http.StatusCreated

	newRow := [][]string{{strconv.Itoa(newPokemon.ID), newPokemon.Name, newPokemon.Region}}
	PokemonSlice, err = csv.WriteCSV(newRow)

	if err != nil {
		log.Println(err.Error())
		status = http.StatusInternalServerError
		err = errors.New(con.Conf.ServerError)
		return status, err
	}

	if len(PokemonSlice) == 0 {
		status = http.StatusOK
		err = errors.New("The pokemon already exists")
		return status, err

	}

	return status, err

}

func GetAllPokemon() (status int, err error) {
	PokemonSlice = []s.Pokemon{}

	status = http.StatusOK

	rows, err := csv.ReadCSV()
	if err != nil {
		status = http.StatusInternalServerError
		err = errors.New(con.Conf.ServerError)
		return status, err

	}

	for _, row := range rows {
		appendSlice(row)
	}

	PokemonSlice = PokemonSlice[1:len(PokemonSlice)]

	if len(PokemonSlice) == 0 {
		status = http.StatusNotFound
		err = errors.New("There are no pokemon")
		return status, err
	}

	PokemonSlice = ro.SortSlice(PokemonSlice)

	return status, err
}

func GetPokemonByRegion(region string) (status int, err error) {
	PokemonSlice = []s.Pokemon{}
	status = http.StatusOK

	rows, err := csv.ReadCSV()

	if err != nil {
		status = http.StatusInternalServerError
		err = errors.New(con.Conf.ServerError)
		//errorOnServer(c)
		return status, err
	}

	for _, row := range rows {
		if row[2] == region {
			appendSlice(row)
		}
	}

	if len(PokemonSlice) == 0 {
		status = http.StatusNotFound
		err = errors.New("There are no rows in the " + region + " region")
		return status, err

	}

	PokemonSlice = ro.SortSlice(PokemonSlice)
	return status, err
}

func SearchById(id string) (status int, err error) {
	PokemonSlice = []s.Pokemon{}
	status = http.StatusOK

	rows, err := csv.ReadCSV()

	if err != nil {
		status = http.StatusInternalServerError
		err = errors.New(con.Conf.ServerError)
		//errorOnServer(c)
		return status, err
	}

	exists, index := ro.ExistsRow(rows, id, "id")

	if !exists {
		status = http.StatusNotFound
		err = errors.New("The id " + id + " does not exist")

		//errorResponse(c, status, errorMessage)
		return status, err

	}
	appendSlice(rows[index])
	//successResponse(c, status)

	return status, err
}

func SearchByName(name string) (status int, err error) {
	PokemonSlice = []s.Pokemon{}
	status = http.StatusOK

	rows, err := csv.ReadCSV()

	if err != nil {
		status = http.StatusInternalServerError
		err = errors.New(con.Conf.ServerError)
		//errorOnServer(c)
		return status, err
	}

	exists, index := ro.ExistsRow(rows, name, "name")

	if !exists {
		status = http.StatusNotFound
		err = errors.New("The name " + name + " does not exist")
		return status, err

	}
	appendSlice(rows[index])
	return status, err

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
	pokemon := ro.AddRowToSlice(row)
	PokemonSlice = append(PokemonSlice, pokemon)

}
