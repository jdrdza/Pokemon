package main

import (
	"testing"

	"github.com/jdrdza/pokemon/file"
	"github.com/jdrdza/pokemon/method"

	"github.com/go-playground/assert/v2"
)

type Pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
}

type controller struct {
	method method.Method
}

var rows, err = file.NewFile()
var meth = method.NewMethod(rows)

func TestPokemonByRegion(t *testing.T) {
	var testCases = []struct {
		name           string
		region         string
		statusResponse int
		pokemonSlice   []Pokemon
		err            error
	}{
		{
			name:         "Get all pokemon from the csv file by the region hoenn",
			region:       "hoenn",
			pokemonSlice: []Pokemon{{252, "treecko", "hoenn"}},
		},
		{
			name:         "Get all pokemon from the csv file by the region Galardon",
			region:       "Galardon",
			pokemonSlice: []Pokemon{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pokemonSlice := meth.GetPokemonByRegion(tc.region)
			assert.IsEqual(pokemonSlice, tc.pokemonSlice)
		})
	}
}

func TestPokemonById(t *testing.T) {

	var testCases = []struct {
		name         string
		id           int
		pokemonSlice []Pokemon
	}{
		{
			name:         "Get the pokemon with the ID 6",
			id:           6,
			pokemonSlice: []Pokemon{{6, "charizard", "kanto"}},
		},
		{
			name:         "Get the pokemon with the ID 150",
			id:           150,
			pokemonSlice: []Pokemon{{150, "mewtwo", "kanto"}},
		},
		{
			name:         "Get the pokemon with the ID 158",
			id:           158,
			pokemonSlice: []Pokemon{},
		},
		{
			name:         "Get the pokemon with the ID 252",
			id:           252,
			pokemonSlice: []Pokemon{{252, "treecko", "hoenn"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pokemonSlice := meth.GetPokemonById(tc.id)
			assert.IsEqual(pokemonSlice, tc.pokemonSlice)
		})
	}
}

func TestPokemonByName(t *testing.T) {

	var testCases = []struct {
		name         string
		pokemonName  string
		pokemonSlice []Pokemon
	}{
		{
			name:         "Get the pokemon with the name bulbasaur",
			pokemonName:  "bulbasaur",
			pokemonSlice: []Pokemon{{1, "bulbasaur", "kanto"}},
		},
		{
			name:         "Get the pokemon with the name charizandro",
			pokemonName:  "charizandro",
			pokemonSlice: []Pokemon{},
		},
		{
			name:         "Get the pokemon with the name Mewtwo",
			pokemonName:  "Mewtwo",
			pokemonSlice: []Pokemon{},
		},
		{
			name:         "Get the pokemon with the name mewtwo",
			pokemonName:  "mewtwo",
			pokemonSlice: []Pokemon{{150, "mewtwo", "kanto"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pokemonSlice := meth.GetPokemonByName(tc.pokemonName)
			assert.IsEqual(pokemonSlice, tc.pokemonSlice)
		})
	}
}
