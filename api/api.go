package api

import (
	endpoints "github.com/MaciejTe/twitter/api/endpoints"
	"github.com/MaciejTe/twitter/pkg/config"
	"github.com/MaciejTe/twitter/pkg/messenger"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewServer inititiates database connection and setups API endpoints
func NewServer(settings *config.Config, database *mongo.Database, messenger messenger.Messenger) (*fiber.App, error) {
	app := fiber.New()

	endpoints.MessagesEndpoint(app, messenger)
	return app, nil
}
