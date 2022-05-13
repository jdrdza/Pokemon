//Author: Juan Daniel Rodríguez Aguilar
//Description: This module contains the structures used

package structs

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
