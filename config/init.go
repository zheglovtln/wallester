package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type DBConfig struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func LoadConfig(path string) (*DBConfig, error) {
	configFile, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Unable to read configuration file %s", path)
	}

	config := new(DBConfig)

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse configuration file %s", path)
	}

	return config, nil
}
