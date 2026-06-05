package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	RabbitMQ struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
	} `yaml:"rabbitmq"`
	Postgres struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		DB       string `yaml:"db"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
}

func (gC *Config) GetConfig() {
	cfg, err := gC.Load(filepath.Join("config", "development.yaml"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.RabbitMQ.Username)
	fmt.Println(cfg.Postgres.Host)
}
func (gC *Config) Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
