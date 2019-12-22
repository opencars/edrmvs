package config

import "github.com/BurntSushi/toml"

type Settings struct {
	DB  Database `toml:"database"`
	HSC HSC      `toml:"hsc"`
}

type Database struct {
	Network  string `toml:"network"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
	Pool     int    `toml:"pool"`
}

type HSC struct {
	BaseURL string `toml:"base_url"`
}

// New reads application configuration from specified file path.
func New(path string) (*Settings, error) {
	config := &Settings{}
	if _, err := toml.DecodeFile(path, config); err != nil {
		return nil, err
	}

	return config, nil
}
