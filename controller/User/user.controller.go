package user

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/resistance22/university_project/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type UserControllers struct {
	store *db.Store
}

type UserRegisterArgs struct {
	FirstName string
	LastName  string
	UserName  string
	Password  string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func generateNewUserArgs(u *UserRegisterArgs) *db.CreateUserParams {
	passwordHash, err := HashPassword(u.Password)
	if err != nil {
		log.Fatal("Can't Hash Password")
	}

	return &db.CreateUserParams{
		ID: pgtype.UUID{
			Bytes: uuid.New(),
			Valid: true,
		},
		CreatedAt: pgtype.Date{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		FirstName: pgtype.Text{
			String: u.FirstName,
			Valid:  true,
		},
		LastName: pgtype.Text{
			String: u.LastName,
			Valid:  true,
		},
		UserName: pgtype.Text{
			String: u.UserName,
			Valid:  true,
		},
		Password: pgtype.Text{
			String: passwordHash,
			Valid:  true,
		},
	}
}

func (c *UserControllers) Register(args *UserRegisterArgs) (db.AppUser, error) {
	params := generateNewUserArgs(args)
	user, err := c.store.Queries.CreateUser(context.Background(), *params)
	return user, err
}
