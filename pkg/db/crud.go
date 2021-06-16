package db

import (
	"github.com/MaciejTe/twitter/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CRUD is data access interface; usually contains Create, Read, Update and Delete methods,
// but I removed unnecessary methods for sake of Twitter app simplicity
type CRUD interface {
	Create(models.Message) (models.Message, error)
	Count(filters bson.M) (int64, error)
}
