package config

import (
	s "Pokemon/structs"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Conf = Config()

func Config() s.Conf {
	data, err := ioutil.ReadFile("data\\config.yaml")

	if err != nil {

		log.Fatal(err)
	}
	y := s.Conf{}

	err = yaml.Unmarshal([]byte(data), &y)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return y
}
