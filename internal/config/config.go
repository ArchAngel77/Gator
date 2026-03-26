
package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbUrl		string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	fullPath := homeDir + "/" + configFileName
	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	return write(*cfg)
}

func write(cfg Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	fullPath := homeDir + "/" + configFileName
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
        err = encoder.Encode(cfg)
	return err
}
