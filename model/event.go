package model

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	EventId     uuid.UUID
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserId      uuid.UUID
}

var Events = make(map[string]Event)

func NewEvent(name, description, location string, dateTime time.Time) *Event {
	return &Event{
		Name:        name,
		Description: description,
		Location:    location,
		DateTime:    dateTime,
	}
}
