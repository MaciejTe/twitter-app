package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config holds application configuration. Currently Config struct is populated from environment variables (12-factor app).
type Config struct {
	Server struct {
		Port string `envconfig:"API_PORT"`
	}
	Database struct {
		URI            string `envconfig:"DB_URI"`
		DbName         string `envconfig:"DB_NAME"`
		CollectionName string `envconfig:"DB_COLLECTION_NAME"`
	}
}

// NewConfig creates Config structure
func NewConfig() *Config {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
