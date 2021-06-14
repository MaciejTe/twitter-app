package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	endpoints "github.com/MaciejTe/twitter/api/enpoints"
	"github.com/MaciejTe/twitter/pkg/config"
	"github.com/MaciejTe/twitter/pkg/db"
	"github.com/MaciejTe/twitter/pkg/messenger"
)

func main() {
	log.Info("Reading application configuration")
	settings := config.NewConfig()

	// db connection here
	connector := db.NewMongoConnector(*settings)
	err := connector.Connect()
	if err != nil {
		log.Fatal(err)
	}
	database := connector.Client.Database(settings.Database.DbName)
	defer connector.Disconnect()

	app := fiber.New()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	twitter := messenger.NewTwitter(database, settings.Database.CollectionName)

	endpoints.MessagesEndpoint(app, twitter)
	log.Info("Starting API")

	log.Fatal(app.Listen(fmt.Sprintf(":%s", settings.Server.Port)))
}
