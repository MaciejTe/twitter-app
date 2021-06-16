package messenger

import (
	"errors"
	"testing"
	"time"

	"github.com/MaciejTe/twitter/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type MockedCRUD struct {
	mock.Mock
}

func (m *MockedCRUD) Create(model models.Message) (models.Message, error) {
	args := m.Called(model)
	return args.Get(0).(models.Message), args.Error(1)
}

func (m *MockedCRUD) Count(filters bson.M) (int64, error) {
	args := m.Called()
	return int64(args.Int(0)), args.Error(1)
}

func TestFetchMessagesCountProper(t *testing.T) {
	crud := new(MockedCRUD)

	twitter := NewTwitter(crud, "test_collection")
	rawFilters := models.MessageFilter{
		Tags: []string{"tag1", "tag2"},
		From: time.Now(),
		To:   time.Now(),
	}
	crud.On("Count").Return(2013, nil)
	count, err := twitter.FetchMessagesCount(rawFilters)
	assert.Equal(t, int64(2013), count, "Expected document count is different than actual one")
	assert.Equal(t, nil, err, "Expected error is different than actual one")
}

func TestFetchMessagesCountError(t *testing.T) {
	crud := new(MockedCRUD)

	twitter := NewTwitter(crud, "test_collection")
	rawFilters := models.MessageFilter{
		Tags: []string{"tag1", "tag2"},
		From: time.Now(),
		To:   time.Now(),
	}
	crud.On("Count").Return(-1, errors.New("Database client disconnected"))
	count, err := twitter.FetchMessagesCount(rawFilters)
	assert.Equal(t, int64(-1), count, "Expected document count is different than actual one")
	assert.Equal(t, errors.New("Database client disconnected"), err, "Expected error is different than actual one")
}

func TestPrepareFilters(t *testing.T) {
	rawFilters := models.MessageFilter{
		Tags: []string{"tag1", "tag2"},
		From: time.Time{},
		To:   time.Time{},
	}
	expectedResult := bson.M{
		"tags":       bson.M{"$in": []string{"tag1", "tag2"}},
		"created_at": bson.M{"$gte": time.Time{}, "$lte": time.Time{}},
	}
	result := prepareFilters(rawFilters)
	assert.Equal(t, expectedResult, result, "Expected document count is different than actual one")
}
