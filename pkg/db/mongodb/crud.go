package mongodb

import (
	"context"
	"time"

	"github.com/MaciejTe/twitter/pkg/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoCRUD implements CRUD interface, allows to operate on mongoDB database
type MongoCRUD struct {
	Database   *mongo.Database
	collection *mongo.Collection
}

// NewMongoCRUD creates MongoCRUD structure, ready to work on given database and collection
func NewMongoCRUD(connector *MongoConnector, dbName, collectionName string) *MongoCRUD {
	var crud MongoCRUD
	database := connector.Client.Database(dbName)
	crud.Database = database
	collection := crud.Database.Collection(collectionName)
	crud.collection = collection
	return &crud
}

// Create inserts document into mongoDB
func (m *MongoCRUD) Create(message models.Message) (models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	result, err := m.collection.InsertOne(ctx, message)
	if err != nil {
		log.Error("Failed to insert document into mongoDB. Details: ", err)
		return models.Message{}, err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		message.ID = oid.Hex()
	}
	log.Info("Inserted ID: ", message.ID)
	return message, nil
}

// Count returns number of found documents according to given filters
func (m *MongoCRUD) Count(filters bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	count, err := m.collection.CountDocuments(
		ctx,
		filters,
	)
	if err != nil {
		log.Error("Failed to fetch document count from mongoDB. Details: ", err)
		return -1, err
	}
	return count, nil
}
