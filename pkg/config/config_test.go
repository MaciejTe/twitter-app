package config

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigDefaultValues(t *testing.T) {
	os.Setenv("DB_URI", "test_db_uri")
	settings, err := NewConfig()

	assert.Equal(t, nil, err, "Improper DB URI port found")
	assert.Equal(t, "3000", settings.Server.Port, "Improper API port found")
	assert.Equal(t, "test_db_uri", settings.Database.URI, "Improper DB URI port found")
	assert.Equal(t, "twitter", settings.Database.DbName, "Improper API port found")
	assert.Equal(t, "messages", settings.Database.CollectionName, "Improper API port found")
}

func TestProperConfig(t *testing.T) {
	os.Setenv("API_PORT", "4000")
	os.Setenv("DB_URI", "test_db_uri")
	os.Setenv("DB_NAME", "test_db_name")
	os.Setenv("DB_COLLECTION_NAME", "test_collection")
	settings, err := NewConfig()

	assert.Equal(t, nil, err, "Improper DB URI port found")
	assert.Equal(t, "4000", settings.Server.Port, "Improper API port found")
	assert.Equal(t, "test_db_uri", settings.Database.URI, "Improper DB URI port found")
	assert.Equal(t, "test_db_name", settings.Database.DbName, "Improper API port found")
	assert.Equal(t, "test_collection", settings.Database.CollectionName, "Improper API port found")
}

func TestConfigLackOfDbURI(t *testing.T) {
	os.Unsetenv("API_PORT")
	os.Unsetenv("DB_URI")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("DB_COLLECTION_NAME")

	settings, err := NewConfig()
	assert.Equal(t, (*Config)(nil), settings, "Config should not be created when DB URI is not provided in environment")
	assert.Equal(t, errors.New("required key DB_URI missing value"), err, "Improper DB URI port found")
}
