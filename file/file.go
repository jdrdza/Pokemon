package file

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strings"
)

type csvFile struct {
	Rows [][]string
}

type CSV interface {
	GetRows() [][]string
	WriteCSV(rows [][]string) (err error)
}

func (c *csvFile) WriteCSV(rows [][]string) (err error) {
	file, err := os.Create("data\\pokemon.csv")

	if err != nil {
		log.Println(err.Error())
		return err
	}

	file.WriteString("\xEF\xBB\xBF")

	newFile := csv.NewWriter(file)
	rows = append(rows)

	newFile.Comma = ';'
	newFile.WriteAll(rows)
	newFile.Flush()
	file.Close()

	return nil
}

func (c *csvFile) GetRows() (rows [][]string) {
	return c.Rows
}

func NewFile() (c *csvFile, err error) {
	File, err := os.Open("data/pokemon.csv")
	if err != nil {
		log.Println(err.Error())
		err = errors.New("The File does not exist")
		return nil, err
	}

	split := strings.Split(File.Name(), ".")
	ext := split[len(split)-1]

	if ext != "csv" {
		err = errors.New("The " + ext + " extension is not correct")
		log.Println(err.Error())
		return nil, err
	}

	reader := csv.NewReader(File)

	reader.Comma = ';'
	rows, err := reader.ReadAll()
	if err != nil {
		log.Println(err.Error())
		err = errors.New("The File could not be reading")
		return nil, err
	}

	return &csvFile{Rows: rows}, nil
}
