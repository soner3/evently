CREATE TABLE IF NOT EXISTS event(
    event_id BINARY(16) UNIQUE NOT NULL PRIMARY KEY,
    user_id  BINARY(16) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    location VARCHAR(255),
    date_time DATETIME NOT NULL
);