package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port string `yaml:"port"`
}

func NewConfig() *Config {
	// Load .env configuration file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// got path config from .env
	path := os.Getenv("PATH_CONFIG")
	// read yaml configuration file with all variables
	file, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("could not open config file: %v", err)
	}

	if len(file) == 0 {
		log.Fatalf("config is empty: %v", err)
	}

	var conf *Config

	if err := yaml.Unmarshal(file, &conf); err != nil {
		log.Fatalf("could not unmarshal config: %v", err)
	}

	return conf
}
