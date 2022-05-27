package main

import (
	m "Pokemon/methods"
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
)

type Pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
}

/*
var PokemonSlice = []Pokemon{
	{4, "charmander", "kanto"}, {5, "charmeleon", "kanto"}, {6, "charizard", "kanto"},
	{158, "totodile", "johto"}, {159, "croconaw", "johto"}, {160, "feraligatr", "johto"},
	{7, "squirtle", "kanto"}, {8, "wartortle", "kanto"}, {9, "blastoise", "kanto"},
	{1, "bulbasaur", "kanto"}, {2, "ivysaur", "kanto"}, {3, "venusaur", "kanto"},
	{252, "treecko", "hoenn"}, {255, "torchic", "hoenn"}, {258, "mudkip", "hoenn"},
	{152, "chikorita", "johto"}, {153, "bayleef", "johto"}, {154, "meganium", "johto"},
	{155, "cyndaquil", "johto"}, {156, "quilava", "johto"}, {157, "typhlosion", "johto"},
}
*/
func TestAllPokemon(t *testing.T) {

	var testCases = []struct {
		name           string
		statusResponse int
		pokemonSlice   []Pokemon
		err            error
	}{
		{
			name:           "Get all pokemon from the csv file",
			statusResponse: 200,
			pokemonSlice: []Pokemon{
				{1, "bulbasaur", "kanto"}, {2, "ivysaur", "kanto"}, {3, "venusaur", "kanto"},
				{4, "charmander", "kanto"}, {5, "charmeleon", "kanto"}, {6, "charizard", "kanto"},
				{7, "squirtle", "kanto"}, {8, "wartortle", "kanto"}, {9, "blastoise", "kanto"},
				{152, "chikorita", "johto"}, {153, "bayleef", "johto"}, {154, "meganium", "johto"},
				{155, "cyndaquil", "johto"}, {156, "quilava", "johto"}, {157, "typhlosion", "johto"},
				{158, "totodile", "johto"}, {159, "croconaw", "johto"}, {160, "feraligatr", "johto"},
				{252, "treecko", "hoenn"}, {255, "torchic", "hoenn"}, {258, "mudkip", "hoenn"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, err := m.GetAllPokemon()
			assert.Equal(t, status, tc.statusResponse)
			assert.Equal(t, err, tc.err)
			assert.IsEqual(m.PokemonSlice, tc.pokemonSlice)

		})
	}
}

func TestPokemonByRegion(t *testing.T) {

	var testCases = []struct {
		name           string
		region         string
		statusResponse int
		pokemonSlice   []Pokemon
		err            error
	}{
		{
			name:           "Get all pokemon from the csv file by the region kanto",
			region:         "kanto",
			statusResponse: 200,
			pokemonSlice: []Pokemon{
				{1, "bulbasaur", "kanto"}, {2, "ivysaur", "kanto"}, {3, "venusaur", "kanto"},
				{4, "charmander", "kanto"}, {5, "charmeleon", "kanto"}, {6, "charizard", "kanto"},
				{7, "squirtle", "kanto"}, {8, "wartortle", "kanto"}, {9, "blastoise", "kanto"},
			},
		},
		{
			name:           "Get all pokemon from the csv file by the region Galardon",
			region:         "Galardon",
			statusResponse: 404,
			pokemonSlice:   []Pokemon{},
			err:            errors.New("There are no rows in the Galardon region"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, err := m.GetPokemonByRegion(tc.region)
			assert.Equal(t, status, tc.statusResponse)
			assert.Equal(t, err, tc.err)
			assert.IsEqual(m.PokemonSlice, tc.pokemonSlice)
		})
	}
}

func TestPokemonById(t *testing.T) {

	var testCases = []struct {
		name           string
		id             string
		statusResponse int
		pokemonSlice   []Pokemon
		err            error
	}{
		{
			name:           "Get all pokemon from the csv file by ID 6",
			id:             "6",
			statusResponse: 200,
			pokemonSlice:   []Pokemon{{6, "charizard", "kanto"}},
		},
		{
			name:           "Get all pokemon from the csv file by ID 150",
			id:             "150",
			statusResponse: 404,
			pokemonSlice:   []Pokemon{},
			err:            errors.New("The id 150 does not exist"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, err := m.SearchById(tc.id)
			assert.Equal(t, status, tc.statusResponse)
			assert.Equal(t, err, tc.err)
			assert.IsEqual(m.PokemonSlice, tc.pokemonSlice)
		})
	}
}

func TestPokemonByName(t *testing.T) {

	var testCases = []struct {
		name           string
		pokemonName    string
		statusResponse int
		pokemonSlice   []Pokemon
		err            error
	}{
		{
			name:           "Get all pokemon from the csv file by ID 6",
			pokemonName:    "charizard",
			statusResponse: 200,
			pokemonSlice:   []Pokemon{{6, "charizard", "kanto"}},
		},
		{
			name:           "Get all pokemon from the csv file by ID 150",
			pokemonName:    "charizandro",
			statusResponse: 404,
			pokemonSlice:   []Pokemon{},
			err:            errors.New("The name charizandro does not exist"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, err := m.SearchByName(tc.pokemonName)
			assert.Equal(t, status, tc.statusResponse)
			assert.Equal(t, err, tc.err)
			assert.IsEqual(m.PokemonSlice, tc.pokemonSlice)
		})
	}
}
