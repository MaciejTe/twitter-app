package models

import (
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// MessageFilter represents raw query parameters for GET messages/ endpoint
type MessageFilter struct {
	Tags []string
	From time.Time
	To   time.Time
}

// NewMessageFilter creates raw filters structure, which needs to be modified later on to fit mongo query system
func NewMessageFilter(tags, from, to string) (*MessageFilter, error) {
	var filters MessageFilter
	if tags != "" {
		filters.Tags = strings.Split(tags, ",")
	}
	var err error // nil by default
	if len(from) != 0 {
		filters.From, err = time.Parse(time.RFC3339, from)
		if err != nil {
			log.Error("Failed to convert start filtering date to time.Time. Details: ", err)
			return nil, err
		}
	} else {
		filters.From = time.Time{}
	}
	if len(to) != 0 {
		filters.To, err = time.Parse(time.RFC3339, to)
		if err != nil {
			log.Error("Failed to convert stop filtering date to time.Time. Details: ", err)
			return nil, err
		}
	} else {
		filters.To, err = time.Parse(time.RFC3339, "3000-01-01T00:00:00Z")
		if err != nil {
			log.Error("Failed to convert max date to time.Time. Details: ", err)
			return nil, err
		}
	}
	return &filters, nil
}
