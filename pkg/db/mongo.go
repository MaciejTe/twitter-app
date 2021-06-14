package db

import (
	"context"
	"log"
	"time"

	"github.com/MaciejTe/twitter/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoConnector implements Connector interface; supports connection to mongoDB
type MongoConnector struct {
	Client *mongo.Client
	config config.Config
}

// NewMongoConnector creates MongoConnector interface
func NewMongoConnector(appConfig config.Config) *MongoConnector {
	client, err := mongo.NewClient(options.Client().ApplyURI(appConfig.Database.URI))
	if err != nil {
		log.Fatal(err)
	}
	var connector MongoConnector
	connector.Client = client
	connector.config = appConfig
	return &connector
}

// Connect initializes connection to the database
func (m *MongoConnector) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := m.Client.Connect(ctx)
	if err != nil {
		return err
	}
	err = m.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}

// Disconnect breaks connection with database
func (m *MongoConnector) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	m.Client.Disconnect(ctx)
}
