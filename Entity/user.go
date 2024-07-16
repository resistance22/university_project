package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/resistance22/university_project/utils"
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
		return nil, err
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
