package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/soner3/evently/db/sqlc"
)

var Queries *sqlc.Queries

func InitDB() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/eventlydb?parseTime=true")
	if err != nil {
		log.Fatalln("Could not connect to db:", err)
	}

	initTables(db)

	Queries = sqlc.New(db)
}

func initTables(db *sql.DB) {

	userTableStmt := `CREATE TABLE IF NOT EXISTS user(
    user_id BINARY(16) UNIQUE NOT NULL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);`

	_, err := db.Exec(userTableStmt)
	if err != nil {
		log.Fatalln("Could not create user table:", err)
	}

	eventTableStmt := `CREATE TABLE IF NOT EXISTS event(
    event_id BINARY(16) UNIQUE NOT NULL PRIMARY KEY,
    user_id  BINARY(16) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    location VARCHAR(255),
    date_time DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(user_id)
        ON DELETE CASCADE
);`

	_, err = db.Exec(eventTableStmt)
	if err != nil {
		log.Fatalln("Could not create event table:", err)
	}
}
