package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	entity "github.com/resistance22/university_project/Entity"
	db "github.com/resistance22/university_project/db/sqlc"
)

var (
	ErrNotFound = errors.New("no entry found")
)

type IUserRepository interface {
	Register(c context.Context, u *entity.User) error
	FindUserByUserName(c context.Context, username string) (*entity.User, error)
}

type UserRepository struct {
	store *db.Store
}

func createUserParams(user *entity.User) *db.CreateUserParams {
	return &db.CreateUserParams{
		ID: pgtype.UUID{
			Bytes: user.ID,
			Valid: true,
		},
		CreatedAt: pgtype.Timestamp{
			Time:  user.CreatedAt,
			Valid: true,
		},
		FirstName: pgtype.Text{
			String: user.FirstName,
			Valid:  true,
		},
		LastName: pgtype.Text{
			String: user.LastName,
			Valid:  true,
		},
		UserName: pgtype.Text{
			String: user.UserName,
			Valid:  true,
		},
		Password: pgtype.Text{
			String: user.Password,
			Valid:  true,
		},
	}
}

func createEntityFromDbUser(user *db.AppUser) (*entity.User, error) {
	userUUID, err := uuid.FromBytes(user.ID.Bytes[:])
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:        userUUID,
		CreatedAt: user.CreatedAt.Time,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		UserName:  user.UserName.String,
		Password:  user.Password.String,
	}, nil
}

func (repo *UserRepository) Register(ctx context.Context, user *entity.User) error {
	params := createUserParams(user)
	_, err := repo.store.Queries.CreateUser(ctx, *params)
	return err
}

func (repo *UserRepository) FindUserByUserName(c context.Context, username string) (*entity.User, error) {
	user, err := repo.store.Queries.GetUserByUserName(c, pgtype.Text{
		String: username,
		Valid:  true,
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, ErrNotFound
	}

	return createEntityFromDbUser(&user)
}

func NewUserRepository(store *db.Store) IUserRepository {
	return &UserRepository{
		store: store,
	}
}
