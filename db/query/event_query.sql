-- name: ListEvents :many
SELECT 
    * 
FROM 
    event 
LIMIT 10;

-- name: FindEventById :one
SELECT
    *
FROM 
    event
WHERE
    1 = 1
    AND event_id = ?
LIMIT 1;

-- name: CreateEvent :exec
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
    (?, ?, ?, ?, ?, ?);

-- name: UpdateEventById :exec
UPDATE
    event
SET
    name = ?,
    description = ?,
    location = ?,
    date_time = ?
WHERE
    1 = 1
    AND event_id = ?;

-- name: DeleteEventById :exec
DELETE FROM
    event
WHERE
    1 = 1
    AND event_id = ?;

-- name: EventExistsById :one
SELECT EXISTS(
    SELECT
        *
    FROM 
        event
    WHERE
        1 = 1
        AND event_id = ?
    LIMIT 1
);




