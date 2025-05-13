package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/soner3/evently/db"
	"github.com/soner3/evently/db/sqlc"
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
		Name:        name,
		Description: description,
		Location:    location,
		DateTime:    dateTime,
	}
}

func (e *Event) Save() error {
	ctx := context.Background()
	exists, err := db.Queries.EventExistsById(ctx, e.EventId[:])
	if err != nil {
		return err
	}

	if exists {
		return db.Queries.UpdateEventById(ctx, sqlc.UpdateEventByIdParams{
			EventID:     e.EventId[:],
			Name:        e.Name,
			Description: sql.NullString{String: e.Description, Valid: true},
			Location:    sql.NullString{String: e.Location, Valid: true},
			DateTime:    e.DateTime,
		})
	}
	e.EventId = uuid.New()
	e.UserId = uuid.New()
	return db.Queries.CreateEvent(ctx, sqlc.CreateEventParams{
		EventID:     e.EventId[:],
		Name:        e.Name,
		Description: sql.NullString{String: e.Description, Valid: true},
		Location:    sql.NullString{String: e.Location, Valid: true},
		DateTime:    e.DateTime,
		UserID:      e.UserId[:],
	})
}

func (Event) FindById(eventId uuid.UUID) (*Event, error) {
	event, err := db.Queries.FindEventById(context.Background(), eventId[:])
	if err != nil {
		return nil, err
	}

	return &Event{
		EventId:     uuid.UUID(event.EventID),
		Name:        event.Name,
		Description: event.Description.String,
		Location:    event.Location.String,
		DateTime:    event.DateTime,
		UserId:      uuid.UUID(event.UserID),
	}, nil

}

func (Event) ListEvents() (*[]Event, error) {
	events, err := db.Queries.ListEvents(context.Background())
	if err != nil {
		return nil, err
	}
	resEvents := make([]Event, len(events))

	for i, e := range events {
		resEvents[i] = Event{
			EventId:     uuid.UUID(e.EventID),
			Name:        e.Name,
			Description: e.Description.String,
			Location:    e.Location.String,
			DateTime:    e.DateTime,
			UserId:      uuid.UUID(e.UserID),
		}
	}

	return &resEvents, nil
}

func (e *Event) DeleteById() error {
	return db.Queries.DeleteEventById(context.Background(), e.EventId[:])
}
