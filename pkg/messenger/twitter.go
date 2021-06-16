package messenger

import (
	"time"

	"github.com/MaciejTe/twitter/pkg/db"
	"github.com/MaciejTe/twitter/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Twitter implements Messager interface, is example of how different messaging implementations can be used
type Twitter struct {
	crud db.CRUD
}

// NewTwitter creates ready to work Twitter struct
func NewTwitter(crud db.CRUD, collectionName string) *Twitter {
	var twitter Twitter
	twitter.crud = crud
	return &twitter
}

// InsertMessage inserts sent message to database
func (t *Twitter) InsertMessage(message models.Message) (*models.Message, error) {
	now := time.Now()
	message.CreatedAt = now
	updatedMessage, err := t.crud.Create(message)
	if err != nil {
		return nil, err
	}
	return &updatedMessage, nil
}

// FetchMessagesCount gets the count of messages, filtering them using provided message filters
func (t *Twitter) FetchMessagesCount(rawFilters models.MessageFilter) (int64, error) {
	filters := prepareFilters(rawFilters)
	count, err := t.crud.Count(filters)
	if err != nil {
		return count, err
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
