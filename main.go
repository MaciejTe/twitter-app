package main

import (
	"fmt"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	"github.com/MaciejTe/twitter/api"
	"github.com/MaciejTe/twitter/pkg/config"
	"github.com/MaciejTe/twitter/pkg/db"
	"github.com/MaciejTe/twitter/pkg/messenger"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	log.Info("Reading application configuration")
	settings, err := config.NewConfig()
	if err != nil {
		log.Fatal("Fatal error during config creation: ", err)
	}

	connector := db.NewMongoConnector(*settings)
	err = connector.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database. Details: ", err)
	}
	database := connector.Client.Database(settings.Database.DbName)
	defer connector.Disconnect()

	twitter := messenger.NewTwitter(database, settings.Database.CollectionName)

	app, err := api.NewServer(settings, database, twitter)
	if err != nil {
		log.Fatal("Fatal error during application startup: ", err)
	}
	log.Info("Starting API")

	go func() {
		log.Fatal(app.Listen(fmt.Sprintf(":%s", settings.Server.Port)))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	if err = app.Shutdown(); err != nil {
		log.Error("Failed to shutdown server. Details: ", err)
	}

	log.Info("Shutting down server")
	os.Exit(0)
}
