package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger  LoggerConf
	Storage StorageConf
	HTTP    HttpConf
}

type Level string

const (
	Info  Level = "info"
	Error       = "error"
)

type LoggerConf struct {
	Level    Level
	Filename string
}

type StorageType string

const (
	SQL      StorageType = "sql"
	InMemory             = "in-memory"
)

type StorageConf struct {
	Type StorageType
	Dsn  string
}

type HttpConf struct {
	Host string
	Port string
}

func NewConfig() Config {
	return Config{}
}

func LoadConfig(path string) (*Config, error) {
	configContent, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("invalid config file %s: %w", path, err)
	}

	config := NewConfig()
	err = yaml.Unmarshal(configContent, &config)
	if err != nil {
		return nil, fmt.Errorf("invalid config file content %s: %w", path, err)
	}

	return &config, nil
}
