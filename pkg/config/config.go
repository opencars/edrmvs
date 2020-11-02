package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Server    Server    `yaml:"server"`
	Log       Log       `yaml:"log"`
	DB        Database  `yaml:"database"`
	HSC       HSC       `yaml:"hsc"`
	Extractor Extractor `yaml:"extractor"`
}

// Server represents settings for creating http server.
type Server struct {
	ShutdownTimeout Duration `yaml:"shutdown_timeout"`
	ReadTimeout     Duration `yaml:"read_timeout"`
	WriteTimeout    Duration `yaml:"write_timeout"`
	IdleTimeout     Duration `yaml:"idle_timeout"`
}

// Log represents settings for application logger.
type Log struct {
	Level string `yaml:"level"`
	Mode  string `yaml:"mode"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"ssl_mode"`
}

type HSC struct {
	BaseURL  string `yaml:"base_url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Extractor struct {
	Delay   Duration `yaml:"delay"`
	BackOff Duration `yaml:"backoff"`
}

// New reads application configuration from specified file path.
func New(path string) (*Settings, error) {
	var config Settings

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
