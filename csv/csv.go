//Author: Juan Daniel Rodríguez Aguilar
//Description: This module contains all of methods used to read and write a CSV file

package csv

// import packages
import (
	c "Pokemon/Config"
	"Pokemon/rowsOperations"
	"Pokemon/structs"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strings"
)

// Declaration of global variables
var header = []string{"ID", "Name", "Region"}
var dirFile = c.Conf.CsvDirFile
var extFile = c.Conf.CsvExt
var comma = c.Conf.CsvComma

// Function that reads and csv file
func ReadCSV() ([][]string, error) {

	file, err := os.Open(dirFile)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer file.Close()

	split := strings.Split(file.Name(), ".")
	ext := split[len(split)-1]

	if ext != extFile {
		err = errors.New("The " + ext + " extension is not correct")
		log.Println(err.Error())
		return nil, err
	}
	reader := csv.NewReader(file)

	reader.Comma = comma
	rows, err := reader.ReadAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return rows, nil
}

// Function that writes a new or an existing file
func WriteCSV(newRows [][]string) ([]structs.Pokemon, error) {
	var (
		rows          [][]string
		createPokemon = [][]string{}
		pokemonSlice  = []structs.Pokemon{}
	)

	_, err := os.Open(dirFile)
	if err == nil {
		rows, err = ReadCSV()
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	} else {
		rows = append(rows, header)
	}

	for i := 0; i < len(newRows); i++ {
		row := newRows[i]
		id := row[0]

		if exists, _ := rowsOperations.ExistsRow(rows, id, "id"); !exists {
			createPokemon = append(createPokemon, rowsOperations.AddRowToFile(row)...)
			pokemonSlice = append(pokemonSlice, rowsOperations.AddRowToSlice(row))
		}

	}

	file, err := os.Create(dirFile)
	defer file.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	file.WriteString("\xEF\xBB\xBF")

	newFile := csv.NewWriter(file)

	rows = append(rows, createPokemon...)

	newFile.Comma = comma
	newFile.WriteAll(rows)
	newFile.Flush()

	return pokemonSlice, nil
}
