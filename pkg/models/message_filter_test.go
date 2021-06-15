package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMessageFilterPositiveCases(t *testing.T) {
	testCaseTable := []struct {
		name           string
		inputTags      string
		inputFrom      string
		inputTo        string
		expectedResult *MessageFilter
		expectedError  error
	}{
		{
			name:      "Create raw filters structure from proper parameters",
			inputTags: "tag1,tag2",
			inputFrom: "2021-06-13T15:00:05Z",
			inputTo:   "2022-06-13T15:00:05Z",
			expectedResult: &MessageFilter{
				Tags: []string{"tag1", "tag2"},
				From: time.Time{},
				To:   time.Time{},
			},
		},
		{
			name:      "Insert empty string into from field",
			inputTags: "",
			inputFrom: "",
			inputTo:   "2021-06-13T15:00:05Z",
			expectedResult: &MessageFilter{
				Tags: nil,
				From: time.Time{},
				To:   time.Time{},
			},
			expectedError: nil,
		},
		{
			name:      "Insert empty string into to field",
			inputTags: "",
			inputFrom: "2021-06-13T15:00:05Z",
			inputTo:   "",
			expectedResult: &MessageFilter{
				Tags: nil,
				From: time.Time{},
				To:   time.Time{},
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCaseTable {
		t.Run(tc.name, func(t *testing.T) {
			filter, err := NewMessageFilter(tc.inputTags, tc.inputFrom, tc.inputTo)
			assert.Equal(t, tc.expectedResult.Tags, filter.Tags, "Result tags are different than expected ones")
			assert.IsType(t, tc.expectedResult.From, filter.From, "Result's start time is not time.Time type")
			assert.IsType(t, tc.expectedResult.To, filter.To, "Result's stop time is not time.Time type")
			assert.Equal(t, tc.expectedError, err, "Message filter error is different than expected one")
		})
	}
}

func TestNewMessageFilterNegativeCases(t *testing.T) {
	testCaseTable := []struct {
		name           string
		inputTags      string
		inputFrom      string
		inputTo        string
		expectedResult *MessageFilter
		expectedError  error
	}{
		{
			name:           "Insert improper RFC-3339 string into from field",
			inputTags:      "tag1,tag2",
			inputFrom:      "2021-06-13",
			inputTo:        "2022-06-13T15:00:05Z",
			expectedResult: nil,
			expectedError: &time.ParseError{
				Layout:     "2006-01-02T15:04:05Z07:00",
				Value:      "2021-06-13",
				LayoutElem: "T",
				ValueElem:  "",
				Message:    "",
			},
		},
		{
			name:           "Insert improper RFC-3339 string into to field",
			inputTags:      "tag1,tag2",
			inputFrom:      "",
			inputTo:        "2022-06-13T15:012q5Z",
			expectedResult: nil,
			expectedError: &time.ParseError{
				Layout:     "2006-01-02T15:04:05Z07:00",
				Value:      "2022-06-13T15:012q5Z",
				LayoutElem: ":",
				ValueElem:  "2q5Z",
				Message:    "",
			},
		},
	}

	for _, tc := range testCaseTable {
		t.Run(tc.name, func(t *testing.T) {
			filter, err := NewMessageFilter(tc.inputTags, tc.inputFrom, tc.inputTo)
			assert.Equal(t, tc.expectedResult, filter, "Result tags are different than expected ones")
			assert.Equal(t, tc.expectedError, err, "Message filter error is different than expected one")
		})
	}
}
