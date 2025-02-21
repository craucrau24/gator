package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl string `json:"db_url"`
}

func Read() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	fname := filepath.Join(home, configFileName)
	data, err := os.ReadFile(fname)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
