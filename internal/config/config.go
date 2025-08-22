package config

import ()

type Config struct {
	DBURL    string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func Read(path string) Config {
	cfg := Config{}

	return cfg
}

func (c *Config) SetUser(user string) error {

	return nil
}
