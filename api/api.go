package api

import (
	endpoints "github.com/MaciejTe/twitter/api/enpoints"
	"github.com/MaciejTe/twitter/pkg/config"
	"github.com/MaciejTe/twitter/pkg/db"
	"github.com/MaciejTe/twitter/pkg/messenger"
	"github.com/gofiber/fiber/v2"
)

// NewServer inititiates database connection and setups API endpoints
func NewServer(settings *config.Config) (*fiber.App, error) {
	connector := db.NewMongoConnector(*settings)
	err := connector.Connect()
	if err != nil {
		return nil, err
	}
	database := connector.Client.Database(settings.Database.DbName)
	defer connector.Disconnect()
	twitter := messenger.NewTwitter(database, settings.Database.CollectionName)

	app := fiber.New()

	endpoints.MessagesEndpoint(app, twitter)
	return app, nil
}
