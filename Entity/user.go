package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	utils "github.com/resistance22/university_project/Utils"
)

var (
	ErrPasswordHashFailed = errors.New("password hash failed")
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
}

func NewUser(
	FirstName string,
	LastName string,
	UserName string,
	Password string,
) (*User, error) {
	hashedPassword, err := utils.HashString(Password)

	if err != nil {
		return nil, ErrPasswordHashFailed
	}

	return &User{
		uuid.New(),
		time.Now().UTC(),
		FirstName,
		LastName,
		UserName,
		hashedPassword,
	}, nil
}
