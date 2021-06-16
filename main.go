package main

import (
	"fmt"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	"github.com/MaciejTe/twitter/api"
	"github.com/MaciejTe/twitter/pkg/config"
	mongodb "github.com/MaciejTe/twitter/pkg/db/mongodb"
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

	connector := mongodb.NewMongoConnector(*settings)
	err = connector.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database. Details: ", err)
	}

	crud := mongodb.NewMongoCRUD(connector, settings.Database.DbName, settings.Database.CollectionName)
	defer connector.Disconnect()

	twitter := messenger.NewTwitter(crud, settings.Database.CollectionName)

	app, err := api.NewServer(settings, crud.Database, twitter)
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
