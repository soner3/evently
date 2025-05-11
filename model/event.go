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

func NewEvent(name, description, location string, dateTime time.Time) *Event {
	return &Event{
		EventId:     uuid.New(),
		UserId:      uuid.New(),
		Name:        name,
		Description: description,
		Location:    location,
		DateTime:    dateTime,
	}
}
