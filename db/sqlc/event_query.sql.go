// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: event_query.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createEvent = `-- name: CreateEvent :exec
INSERT INTO 
    event (
        event_id,
        user_id,
        name,
        description,
        location,
        date_time
    )
VALUES 
    (?, ?, ?, ?, ?, ?)
`

type CreateEventParams struct {
	EventID     []byte
	UserID      []byte
	Name        string
	Description sql.NullString
	Location    sql.NullString
	DateTime    time.Time
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) error {
	_, err := q.db.ExecContext(ctx, createEvent,
		arg.EventID,
		arg.UserID,
		arg.Name,
		arg.Description,
		arg.Location,
		arg.DateTime,
	)
	return err
}

const deleteEventById = `-- name: DeleteEventById :exec
DELETE FROM
    event
WHERE
    1 = 1
    AND event_id = ?
`

func (q *Queries) DeleteEventById(ctx context.Context, eventID []byte) error {
	_, err := q.db.ExecContext(ctx, deleteEventById, eventID)
	return err
}

const eventExistsById = `-- name: EventExistsById :one
SELECT EXISTS(
    SELECT
        event_id, user_id, name, description, location, date_time
    FROM 
        event
    WHERE
        1 = 1
        AND event_id = ?
    LIMIT 1
)
`

func (q *Queries) EventExistsById(ctx context.Context, eventID []byte) (bool, error) {
	row := q.db.QueryRowContext(ctx, eventExistsById, eventID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const findEventById = `-- name: FindEventById :one
SELECT
    event_id, user_id, name, description, location, date_time
FROM 
    event
WHERE
    1 = 1
    AND event_id = ?
LIMIT 1
`

func (q *Queries) FindEventById(ctx context.Context, eventID []byte) (Event, error) {
	row := q.db.QueryRowContext(ctx, findEventById, eventID)
	var i Event
	err := row.Scan(
		&i.EventID,
		&i.UserID,
		&i.Name,
		&i.Description,
		&i.Location,
		&i.DateTime,
	)
	return i, err
}

const listEvents = `-- name: ListEvents :many
SELECT 
    event_id, user_id, name, description, location, date_time 
FROM 
    event 
LIMIT 10
`

func (q *Queries) ListEvents(ctx context.Context) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, listEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.EventID,
			&i.UserID,
			&i.Name,
			&i.Description,
			&i.Location,
			&i.DateTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEventById = `-- name: UpdateEventById :exec
UPDATE
    event
SET
    name = ?,
    description = ?,
    location = ?,
    date_time = ?
WHERE
    1 = 1
    AND event_id = ?
`

type UpdateEventByIdParams struct {
	Name        string
	Description sql.NullString
	Location    sql.NullString
	DateTime    time.Time
	EventID     []byte
}

func (q *Queries) UpdateEventById(ctx context.Context, arg UpdateEventByIdParams) error {
	_, err := q.db.ExecContext(ctx, updateEventById,
		arg.Name,
		arg.Description,
		arg.Location,
		arg.DateTime,
		arg.EventID,
	)
	return err
}
