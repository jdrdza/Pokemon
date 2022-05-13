//Author: Juan Daniel Rodríguez Aguilar
//Description: This module contains all of methods used to read from a CSV file

package csv

// import packages
import (
	"Pokemon/constants"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strings"
)

// Declaration of global variables
var header = []string{"ID", "Name", "Region"}

// Function that reads and csv file
func ReadCSV() ([][]string, error) {

	file, err := os.Open(constants.DirFile())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer file.Close()

	split := strings.Split(file.Name(), ".")
	ext := split[len(split)-1]

	if ext != constants.ExtFile() {
		err = errors.New("The " + ext + " extension is not correct")
		log.Println(err.Error())
		return nil, err
	}
	reader := csv.NewReader(file)

	reader.Comma = constants.Comma()
	rows, err := reader.ReadAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return rows, nil
}
