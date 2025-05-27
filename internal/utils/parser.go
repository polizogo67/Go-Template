package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Appdata            string
	Areas_lookup_table string
	Roster             struct {
		Timezone   string `yaml:"timezone"`
		csv        string
		db         string
		table      string
		sharepoint string
	}
	Mqtt struct {
		Broker   string
		Port     int
		User     string
		Password string
	}
}

func Config_parser(path string) Config {
	// Load the file; returns []byte
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Create an empty Car to be are target of unmarshalling
	var cfg Config

	// Unmarshal our input YAML file into empty Car (var c)
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
