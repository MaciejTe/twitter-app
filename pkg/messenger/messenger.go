package messenger

import (
	"github.com/MaciejTe/twitter/pkg/models"
)

// Messenger interface allows to have more than one messaging implementation
type Messenger interface {
	InsertMessage(models.Message) (*models.Message, error)
	FetchMessagesCount(models.MessageFilter) (int64, error)
}
