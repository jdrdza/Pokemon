package method

import (
	"errors"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/jdrdza/pokemon/file"
	"github.com/jdrdza/pokemon/worker"
)

type method struct {
	CSV file.CSV
}

type Method interface {
	PokeAPI(rows [][]string) (pokemonSlice []pokemon, err error)
	PostNewPokemon(id int, name string, region string) (pokemonSlice []pokemon, err error)
	GetAllPokemon() (pokemonSlice []pokemon)
	GetPokemonByRegion(region string) (pokemonSlice []pokemon)
	GetPokemonById(id int) (pokemonSlice []pokemon)
	GetPokemonByName(name string) (pokemonSlice []pokemon)
	GetTypes(types string, items int, items_per_worker int) (pokemonSlice []pokemon, err error)
	DeletePokemonByRegion(region string) (pokemonSlice []pokemon, err error)
	DeletePokemonById(id int) (pokemonSlice []pokemon, err error)
	DeletePokemonByName(name string) (pokemonSlice []pokemon, err error)
}

type pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
}

type Task struct {
	TaskProcessor func()
}

func (t Task) Run() {
	t.TaskProcessor()
}

func NewMethod(r file.CSV) (c *method) {
	return &method{CSV: r}
}

func (m method) PokeAPI(rows [][]string) (pokemonSlice []pokemon, err error) {
	var newPokemon []pokemon
	var newRow []string
	var newRows [][]string

	for _, row := range rows {

		strId := row[0]
		name := row[1]
		region := row[2]

		id, _ := strconv.Atoi(strId)

		newPokemon = m.GetPokemonById(id)
		if len(newPokemon) != 0 {
			continue
		}

		newPokemon = m.GetPokemonByName(name)
		if len(newPokemon) != 0 {
			continue
		}

		newRow = []string{strId, name, region}
		newRows = append(newRows, newRow)
	}
	rows = m.CSV.GetRows()
	rows = append(rows, newRows...)

	err = m.CSV.WriteCSV(rows)
	if err != nil {
		err = errors.New("The file could not be saved")
		return nil, err
	}

	for _, newRow := range newRows {
		pokemonSlice = appendSlice(newRow, pokemonSlice)
	}

	return pokemonSlice, nil
}

func (m method) PostNewPokemon(id int, name string, region string) (pokemonSlice []pokemon, err error) {
	pokemonSlice = m.GetPokemonById(id)
	strId := strconv.Itoa(id)

	if len(pokemonSlice) != 0 {
		err = errors.New("The pokemon id " + strId + " already exists")
		return nil, err
	}

	pokemonSlice = m.GetPokemonByName(name)

	if len(pokemonSlice) != 0 {
		err = errors.New("The pokemon name " + name + " already exists")
		return nil, err
	}

	newRow := []string{strId, name, region}
	rows := m.CSV.GetRows()
	rows = append(rows, newRow)
	err = m.CSV.WriteCSV(rows)
	if err != nil {
		err = errors.New("The file could not be saved")
		return nil, err
	}
	pokemonSlice = appendSlice(newRow, pokemonSlice)
	return pokemonSlice, nil

}

func (m method) GetAllPokemon() (pokemonSlice []pokemon) {
	pokemonSlice = []pokemon{}

	for _, row := range m.CSV.GetRows() {
		pokemonSlice = appendSlice(row, pokemonSlice)
	}

	pokemonSlice = sortSlice(pokemonSlice)
	return pokemonSlice
}

func (m method) GetPokemonByRegion(region string) (pokemonSlice []pokemon) {
	pokemonSlice = []pokemon{}

	for _, row := range m.CSV.GetRows() {
		if row[2] == region {
			pokemonSlice = appendSlice(row, pokemonSlice)
		}
	}

	pokemonSlice = sortSlice(pokemonSlice)
	return pokemonSlice
}

func (m method) GetPokemonById(id int) (pokemonSlice []pokemon) {
	pokemonSlice = []pokemon{}
	strId := strconv.Itoa(id)

	for _, row := range m.CSV.GetRows() {

		if row[0] == strId {
			pokemonSlice = appendSlice(row, pokemonSlice)
			break
		}
	}

	return pokemonSlice
}

func (m method) GetPokemonByName(name string) (pokemonSlice []pokemon) {
	pokemonSlice = []pokemon{}

	for _, row := range m.CSV.GetRows() {
		if row[1] == name {
			pokemonSlice = appendSlice(row, pokemonSlice)
			break
		}
	}

	return pokemonSlice
}

func (m method) GetTypes(types string, items int, items_per_worker int) (pokemonSlice []pokemon, err error) {
	pokemonSlice = []pokemon{}

	rows := m.CSV.GetRows()
	rows = rows[1:]
	size := len(rows)

	if items <= 0 || items_per_worker <= 0 {
		err = errors.New("items and items_per_worker must be greater than 0")
		return nil, err
	}

	if items < items_per_worker {
		err = errors.New("items cannot be less than items_per_worker")
		return nil, err
	}

	if size == 0 {
		err = errors.New("there are no pokemon")
		return nil, err
	}

	numWorkers := items / items_per_worker
	pool := worker.NewGoroutinePool(numWorkers, items_per_worker)
	taskSize := numWorkers * items_per_worker

	types = strings.ToUpper(types)
	var tasks []Task

	for v := 0; v < taskSize; v++ {
		tasks = append(tasks, Task{
			TaskProcessor: func() {},
		})
	}

	taskCounter := 0
	counter := 0

	for taskCounter < taskSize {
		id, err := strconv.Atoi(rows[counter][0])
		if err != nil {
			log.Println("The value " + rows[counter][0] + " is not a number")
		}

		if types == "ODD" {
			if id%2 != 0 {
				pool.ScheduleWork(tasks[taskCounter])
				pokemonSlice = appendSlice(rows[counter], pokemonSlice)
				taskCounter++
			}

		}
		if types == "EVEN" {
			if id%2 == 0 {
				pool.ScheduleWork(tasks[taskCounter])
				pokemonSlice = appendSlice(rows[counter], pokemonSlice)
				taskCounter++
			}

		}

		counter++

		if counter == len(rows) {
			break
		}
	}

	pool.Close()

	return sortSlice(pokemonSlice), nil
}

func (m method) DeletePokemonByRegion(region string) (pokemonSlice []pokemon, err error) {
	pokemonSlice = []pokemon{}
	rows := [][]string{}

	for _, row := range m.CSV.GetRows() {
		if row[2] == region {
			pokemonSlice = appendSlice(row, pokemonSlice)
		} else {
			rows = append(rows, row)
		}
	}

	err = m.CSV.WriteCSV(rows)
	if err != nil {
		err = errors.New("The file could not be saved")
		return nil, err
	}
	pokemonSlice = sortSlice(pokemonSlice)
	return pokemonSlice, nil
}

func (m method) DeletePokemonById(id int) (pokemonSlice []pokemon, err error) {
	pokemonSlice = []pokemon{}
	strId := strconv.Itoa(id)
	rows := [][]string{}

	for _, row := range m.CSV.GetRows() {

		if row[0] == strId {
			pokemonSlice = appendSlice(row, pokemonSlice)
		} else {
			rows = append(rows, row)
		}
	}

	err = m.CSV.WriteCSV(rows)
	if err != nil {
		err = errors.New("The file could not be saved")
		return nil, err
	}

	return pokemonSlice, nil
}

func (m method) DeletePokemonByName(name string) (pokemonSlice []pokemon, err error) {
	pokemonSlice = []pokemon{}
	rows := [][]string{}

	for _, row := range m.CSV.GetRows() {
		if row[1] == name {
			pokemonSlice = appendSlice(row, pokemonSlice)
		} else {
			rows = append(rows, row)
		}
	}

	err = m.CSV.WriteCSV(rows)
	if err != nil {
		err = errors.New("The file could not be saved")
		return nil, err
	}

	return pokemonSlice, nil
}

/*
func (m method)write(newRows [][]string) (pokemonSlice []pokemon, err error){
	rows := m.CSV.GetRows()
	rows = append(rows, newRows...)

	err = m.CSV.WriteCSV(rows)
	if err != nil {
		err = errors.New("The file could not be saved")
		return nil, err
	}

	for _, newRow := range newRows {
		pokemonSlice = appendSlice(newRow, pokemonSlice)
	}

	return pokemonSlice, nil
}
*/
func appendSlice(row []string, pokemonSlice []pokemon) []pokemon {
	var pokemon pokemon

	pokemon.ID, _ = strconv.Atoi(row[0])

	pokemon.Name = row[1]
	pokemon.Region = row[2]

	pokemonSlice = append(pokemonSlice, pokemon)
	return pokemonSlice
}

func sortSlice(pokemonSlice []pokemon) []pokemon {

	sort.Slice(pokemonSlice, func(i, j int) bool {
		return pokemonSlice[i].ID < pokemonSlice[j].ID
	})

	return pokemonSlice
}
