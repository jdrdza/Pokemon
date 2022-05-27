//Author: Juan Daniel Rodríguez Aguilar
//Description: This module contains the structures used

package structs

type Conf struct {
	Environment string `yaml:"environment"`
	ExternalAPI string `yaml:"external-API"`
	CsvDirFile  string `yaml:"csv-dirFile"`
	CsvExt      string `yaml:"csv-ext"`
	CsvComma    rune   `yaml:"csv-comma"`
	EpCopyPkmn  string `yaml:"ep-copyPkmn"`
	EpNewPkmn   string `yaml:"ep-newPkmn"`
	EpAllPkmn   string `yaml:"ep-allPkmn"`
	EpByRegion  string `yaml:"ep-byRegion"`
	EpById      string `yaml:"ep-byId"`
	EpByName    string `yaml:"ep-byName"`
	ServerError string `yaml:"server-error"`
}

// Structures used to get the information from the Pokemon API
type ResponsePokeAPI struct {
	Resp []PokeAPI `json:"results"`
}

type PokeAPI struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

// Structures used to display the information from the CSV file
type Pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
}

type SuccessResponse struct {
	Count   int       `json:"count"`
	Pokemon []Pokemon `json:"pokemon"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"response"`
}
