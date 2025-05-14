package model

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/soner3/evently/db"
	"github.com/soner3/evently/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId   uuid.UUID
	Email    string
	Password string
	Events   *[]Event
}

func (u *User) Save() error {
	ctx := context.Background()
	exists, err := db.Queries.UserExistsByEmail(ctx, u.Email)
	if err != nil {
		return nil
	}
	if exists {
		return errors.New("User with this email already exists")
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.UserId = uuid.New()
	db.Queries.CreateUser(ctx, sqlc.CreateUserParams{
		UserID:   u.UserId[:],
		Email:    u.Email,
		Password: string(pwd),
	})
	return nil
}

func (u *User) FindUserByEmail(getPwd bool) error {
	user, err := db.Queries.FindUserByEmail(context.Background(), u.Email)
	if err != nil {
		return nil
	}

	if getPwd {
		u.Password = user.Password
	}
	u.Email = user.Email
	u.UserId = uuid.UUID(user.UserID)
	return nil

}

func (u *User) DeleteUserById() error {
	return db.Queries.DeleteUserById(context.Background(), u.UserId[:])
}

func (u *User) FindUserByIdWithEvents() error {
	rows, err := db.Queries.FindUserByIdWithReferences(context.Background(), u.UserId[:])
	if err != nil {
		return nil
	}

	u.Email = rows[0].Email

	events := make([]Event, len(rows))

	for _, r := range rows {
		events = append(events, Event{
			EventId:     uuid.MustParse(r.EventID.String),
			Name:        r.Name.String,
			Description: r.Description.String,
			Location:    r.Location.String,
			DateTime:    r.DateTime.Time,
			User:        User{UserId: uuid.UUID(r.UserID)},
		})
	}
	u.Events = &events
	return nil
}
