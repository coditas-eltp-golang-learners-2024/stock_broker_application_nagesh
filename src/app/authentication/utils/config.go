package utils

import (
	"Stock_broker_application/src/app/authentication/models"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Define a struct to represent the YAML data structure

func FetchDatabaseConfig() models.DBConfig {
	yamlFile, err := ioutil.ReadFile("application.yml")
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}

	var config models.DBConfig

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error unmarshalling YAML: %v", err)
	}

	return config
}
