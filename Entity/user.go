package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/resistance22/university_project/utils"
)

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	FirstName string
	LastName  string
	UserName  string
	Password  string
}

func NewUser(
	ID uuid.UUID,
	CreatedAt time.Time,
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
		ID,
		CreatedAt,
		FirstName,
		LastName,
		UserName,
		hashedPassword,
	}, nil
}
