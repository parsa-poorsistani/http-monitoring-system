package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server  ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
  HealthChecker HealthCheckConfig `yaml:"healthcheck"`
}

type HealthCheckConfig struct {
  Interval int `yaml:"interval"`
}
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

func LoadConfig(path string) (*Config, error) {
	filename := filepath.Join(path, "config.yaml")
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(buf, &cfg)
  if err != nil {
    return nil, err
  }
  return &cfg, nil
}


