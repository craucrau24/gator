package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func getConfigFileName() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

type Config struct {
	DbUrl string `json:"db_url"`
}

func Read() (Config, error) {
	fname, err := getConfigFileName()
	if err != nil {
		return Config{}, err
	}

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
