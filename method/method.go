package method

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"

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
	GetPokemonById(id string) (pokemonSlice []pokemon)
	GetPokemonByName(name string) (pokemonSlice []pokemon)
	GetOdd() (pokemonSlice []pokemon)
	GetEven() (pokemonSlice []pokemon)
}

type pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
}

type Task struct {
	Row           []string
	TaskProcessor func([]string)
}

func (t Task) Run() {
	t.TaskProcessor(t.Row)
}

func NewMethod(r file.CSV) (c *method) {
	return &method{CSV: r}
}

func (m method) PokeAPI(rows [][]string) (pokemonSlice []pokemon, err error) {
	var newPokemon []pokemon
	var newRow []string
	var newRows [][]string

	for _, row := range rows {

		id := row[0]
		name := row[1]
		region := row[2]

		newPokemon = m.GetPokemonById(id)
		if len(newPokemon) != 0 {
			continue
		}

		newPokemon = m.GetPokemonByName(name)
		if len(newPokemon) != 0 {
			continue
		}

		newRow = []string{id, name, region}
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
	newId := strconv.Itoa(id)
	pokemonSlice = m.GetPokemonById(newId)

	if len(pokemonSlice) != 0 {
		err = errors.New("The pokemon id " + newId + " already exists")
		return nil, err
	}

	pokemonSlice = m.GetPokemonByName(name)

	if len(pokemonSlice) != 0 {
		err = errors.New("The pokemon name " + name + " already exists")
		return nil, err
	}

	newRow := []string{newId, name, region}
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

func (m method) GetPokemonById(id string) (pokemonSlice []pokemon) {
	pokemonSlice = []pokemon{}

	for _, row := range m.CSV.GetRows() {

		if row[0] == id {
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

func (m method) GetOdd() (pokemonSlice []pokemon) {
	pokemonSlice = []pokemon{}

	rows := m.CSV.GetRows()
	pool := worker.NewGoroutinePool(2)
	taskSize := len(rows)
	taskCounter := 0

	wg := &sync.WaitGroup{}
	wg.Add(taskSize)
	Odd := func(rows []string) {

		time.Sleep(time.Millisecond)

		e, _ := strconv.Atoi(rows[0])

		if e%2 != 0 {
			fmt.Printf("Finished %s\n", rows[0])
			pokemonSlice = appendSlice(rows, pokemonSlice)
		}

		taskCounter++
		wg.Done()
	}

	var tasks []Task

	for v := 0; v < taskSize; v++ {
		tasks = append(tasks, Task{
			Row:           rows[v],
			TaskProcessor: Odd,
		})

	}

	for _, task := range tasks {
		pool.ScheduleWork(task)

	}

	pool.Close()
	wg.Wait()

	return sortSlice(pokemonSlice)
}

func (m method) GetEven() (pokemonSlice []pokemon) {
	pokemonSlice = []pokemon{}

	rows := m.CSV.GetRows()
	pool := worker.NewGoroutinePool(100)
	taskSize := len(rows)
	taskCounter := 0

	wg := &sync.WaitGroup{}
	wg.Add(taskSize)
	Odd := func(rows []string) {

		time.Sleep(time.Millisecond)

		e, _ := strconv.Atoi(rows[0])

		if e%2 == 0 {
			fmt.Printf("Finished %s\n", rows[0])
			pokemonSlice = appendSlice(rows, pokemonSlice)
		}

		taskCounter++
		wg.Done()
	}

	var tasks []Task

	for v := 0; v < taskSize; v++ {
		tasks = append(tasks, Task{
			Row:           rows[v],
			TaskProcessor: Odd,
		})

	}

	for _, task := range tasks {
		pool.ScheduleWork(task)

	}

	pool.Close()
	wg.Wait()

	return sortSlice(pokemonSlice)
}
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
