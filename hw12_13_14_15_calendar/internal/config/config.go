package config

import (
	"fmt"
	"os"

	yamlv3 "gopkg.in/yaml.v3"
)

type Config struct {
	Logger  LoggerConf
	Storage StorageConf
	HTTP    HTTPConf
	GRPC    GRPCConf
}

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
)

type LoggerConf struct {
	Level    Level
	Filename string
}

type StorageType string

const (
	SQL      StorageType = "sql"
	InMemory StorageType = "in-memory"
)

type StorageConf struct {
	Type StorageType
	Dsn  string
}

type HTTPConf struct {
	Host string
	Port string
}

type GRPCConf struct {
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
	err = yamlv3.Unmarshal(configContent, &config)
	if err != nil {
		return nil, fmt.Errorf("invalid config file content %s: %w", path, err)
	}

	return &config, nil
}
