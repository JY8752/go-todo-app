package test

import (
	"testing"
	"time"
	"todo-app/ent"

	"github.com/stretchr/testify/assert"
)

func AssertUser(t *testing.T, expected *ent.User, actual *ent.User) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Email, actual.Email)
	assert.Equal(t, expected.Age, actual.Age)
}

func AssertTodo(t *testing.T, expected, actual *ent.Todo, isValidTime bool) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.UserID, actual.UserID)
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.Detail, actual.Detail)

	if isValidTime {
		AssertTime(t, expected.CreatedAt, actual.CreatedAt)
		AssertTime(t, expected.UpdatedAt, actual.UpdatedAt)
		AssertTime(t, expected.CompletedAt, actual.CompletedAt)
	}
}

func AssertTime(t *testing.T, expected, actual time.Time) {
	type AssertTime struct {
		year   int
		mounth int
		day    int
		hour   int
		minute int
	}
	expectedTime := AssertTime{
		year:   expected.Year(),
		mounth: int(expected.Month()),
		day:    expected.Day(),
		hour:   expected.Hour(),
		minute: expected.Minute(),
	}
	actualTime := AssertTime{
		year:   actual.Year(),
		mounth: int(actual.Month()),
		day:    actual.Day(),
		hour:   actual.Hour(),
		minute: actual.Minute(),
	}
	assert.Equal(t, expectedTime, actualTime)
}
