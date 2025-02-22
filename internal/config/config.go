package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFileName() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

func write(config Config) error {
	fname, err := getConfigFileName()
	if err != nil {
		return err
	}
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile(fname, data, 0644)
	return err
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

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	return write(*c)
}
