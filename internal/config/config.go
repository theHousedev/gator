package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	cfgFile = ".gatorconfig.json"
)

type Config struct {
	DBURL    string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func (cfg *Config) SetUser(user string) (Config, error) {
	cfg.Username = user

	err := write(cfg)
	if err != nil {
		return *cfg, fmt.Errorf("Failed to write config: %w", err)
	}

	return *cfg, nil
}

func Read() (Config, error) {
	cfg := Config{}

	cfgPath, err := getConfigPath()
	if err != nil {
		return cfg, fmt.Errorf("Failed to find config path: %w", err)
	}

	jsonData, err := os.ReadFile(cfgPath)
	if err != nil {
		return cfg, fmt.Errorf("Failed to read config file: %w", err)
	}

	err = json.Unmarshal(jsonData, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("Failed to unmarshal config: %w", err)
	}

	return cfg, nil
}

func getConfigPath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Unable to get home dir: %w", err)
	}

	return filepath.Join(homedir, cfgFile), nil
}

func write(cfg *Config) error {
	path, err := getConfigPath()
	if err != nil {
		return fmt.Errorf("Failed to find config path: %w", err)
	}

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Failed to marshal config: %w", err)
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write config: %w", err)
	}

	return nil
}
