package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config holds application configuration. Currently Config struct is populated from environment variables (12-factor app).
type Config struct {
	Server struct {
		Port string `envconfig:"API_PORT" default: "3000"`
	}
	Database struct {
		URI            string `required:"true" envconfig:"DB_URI"`
		DbName         string `envconfig:"DB_NAME" default: "twitter" `
		CollectionName string `envconfig:"DB_COLLECTION_NAME" default: "messages"`
	}
}

// NewConfig creates Config structure
func NewConfig() (*Config, error) {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
