package config

import (
	"os"

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

func LoadConfig() *Config {
  return &Config{
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: DatabaseConfig{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Name:     os.Getenv("DATABASE_NAME"),
		},
		HealthChecker: HealthCheckConfig{
			Interval: 5, // This could also be read from an environment variable or remain as a constant.
		}, 
	}	
}


