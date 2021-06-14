package messenger

import (
	"context"
	"time"

	"github.com/MaciejTe/twitter/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	log "github.com/sirupsen/logrus"
)

// Twitter implements Messager interface, is example of how different messaging implementations can be used
type Twitter struct {
	collection *mongo.Collection
}

// NewTwitter creates ready to work Twitter struct
func NewTwitter(database *mongo.Database, collectionName string) *Twitter {
	messagesCollection := database.Collection(collectionName)

	var twitter Twitter
	twitter.collection = messagesCollection
	return &twitter
}

// InsertMessage inserts sent message to database
func (t *Twitter) InsertMessage(message *models.Message) (*models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	now := time.Now()
	message.CreatedAt = now
	result, err := t.collection.InsertOne(ctx, message)
	if err != nil {
		log.Error("Failed to insert document into mongoDB. Details: %v", err)
		return nil, err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		message.ID = oid.Hex()
	}
	log.Info("Inserted ID: %v", message.ID)
	return message, nil
}

// FetchMessagesCount gets the count of messages, filtering them using provided message filters
func (t *Twitter) FetchMessagesCount(rawFilters models.MessageFilter) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	filters := prepareFilters(rawFilters)
	count, err := t.collection.CountDocuments(
		ctx,
		filters,
	)
	if err != nil {
		log.Error("Failed to fetch document count from mongoDB. Details: %v", err)
		return -1, err
	}
	return count, nil
}

// prepareFilters creates bson.M filter which can be used in mongo queries
func prepareFilters(filtersRaw models.MessageFilter) bson.M {
	timeFilters := bson.M{}
	tagFilter := bson.M{}
	filters := bson.M{}
	if filtersRaw.Tags != nil {
		tagFilter["$in"] = filtersRaw.Tags
		filters["tags"] = tagFilter
	}
	timeFilters["$gte"] = filtersRaw.From
	timeFilters["$lte"] = filtersRaw.To
	filters["created_at"] = timeFilters
	return filters
}
