//Author: Juan Daniel Rodríguez Aguilar
//Description: This module contains all validations such as if a row exists, add rows and sort slice

package rowsOperations

// import packages
import (
	s "Pokemon/structs"
	"sort"
	"strconv"
)

// Function that returns true and the index of the slice if a row exists
func ExistsRow(rows [][]string, value string, columnName string) (bool, int) {
	exists := false
	indexToSearch := 0
	index := 0

	switch columnName {
	case "id":
		indexToSearch = 0
	case "name":
		indexToSearch = 1

	}

	for id, row := range rows {
		if row[indexToSearch] == value {
			exists = true
			index = id
			break
		}
	}
	return exists, index
}

// Function that adds the new row in the Pokemon structure
func AddRowToSlice(row []string) s.Pokemon {
	var pokemon s.Pokemon

	pokemon.ID, _ = strconv.Atoi(row[0])
	pokemon.Name = row[1]
	pokemon.Region = row[2]

	return pokemon
}

// Function that adds the new row in the CSV file
func AddRowToFile(row []string) [][]string {
	id := row[0]
	name := row[1]
	region := row[2]

	newPokemon := [][]string{{id, name, region}}

	return newPokemon
}

// Function that sorts the slice by ID in ascending order
func SortSlice(pokemonSlice []s.Pokemon) []s.Pokemon {

	sort.Slice(pokemonSlice, func(i, j int) bool {
		return pokemonSlice[i].ID < pokemonSlice[j].ID
	})

	return pokemonSlice
}
