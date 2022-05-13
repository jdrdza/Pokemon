//Author: Juan Daniel Rodríguez Aguilar
//Description: This module contains all the constants used

package constants

// Declaration of constants
const (
	dirFile    = "..\\data\\pokemon.csv"
	extFile    = "csv"
	comma      = ';'
	message    = "response"
	domain     = "localhost:8080"
	allPokemon = "/allPokemon"
	byRegion   = "/pokemonByRegion/:region"
	byId       = "/pokemonById/:id"
	byName     = "/pokemonByName/:name"
)

// Methods returning a constant
func DirFile() string {
	return dirFile
}

func ExtFile() string {
	return extFile
}

func Comma() rune {
	return comma
}

func Message() string {
	return message
}

func Domain() string {
	return domain
}

func AllPokemon() string {
	return allPokemon
}

func ByRegion() string {
	return byRegion
}

func ById() string {
	return byId
}

func ByName() string {
	return byName
}
