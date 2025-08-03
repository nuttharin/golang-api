package config

import (
	"encoding/json"
	"golang-api/database"
	"os"
)

type ServerConfig struct {
	Host  string `json:"host"`
	Port  string `json:"port"`
	Debug string `json:"debug"`
}

type DBConnConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

type Config struct {
	Environment string                      `json:"environment"`
	Server      ServerConfig                `json:"server"`
	Database    database.DatabaseConnection `json:"database"`
}

func LoadConfig(filepath string) (*Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}

	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
