package main

import (
	"fmt"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	"github.com/MaciejTe/twitter/api"
	"github.com/MaciejTe/twitter/pkg/config"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	log.Info("Reading application configuration")
	settings, err := config.NewConfig()
	if err != nil {
		log.Fatal("Fatal error during config creation: ", err)
	}

	app, err := api.NewServer(settings)
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
	app.Shutdown()

	log.Info("Shutting down server")
	os.Exit(0)
}
