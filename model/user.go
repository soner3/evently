package model

import "github.com/google/uuid"

type User struct {
	UserId   uuid.UUID
	Email    string
	Password string
}
