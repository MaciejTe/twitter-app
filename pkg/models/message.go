package models

import (
	"time"
)

// Message represents how Twitter message is composed
type Message struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Text      string    `json:"text,omitempty" bson:"text,omitempty"`
	Tags      []string  `json:"tags,omitempty" bson:"tags,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
