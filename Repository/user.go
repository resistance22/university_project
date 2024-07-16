package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	entity "github.com/resistance22/university_project/Entity"
	db "github.com/resistance22/university_project/db/sqlc"
)

type UserRepository struct {
	store *db.Store
}

func createUserParams(user *entity.User) *db.CreateUserParams {
	return &db.CreateUserParams{
		ID: pgtype.UUID{
			Bytes: user.ID,
			Valid: true,
		},
		CreatedAt: pgtype.Date{
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

func (repo *UserRepository) Register(ctx context.Context, user *entity.User) error {
	params := createUserParams(user)
	_, err := repo.store.Queries.CreateUser(ctx, *params)
	return err
}

func NewUserRepository(store *db.Store) *UserRepository {
	return &UserRepository{
		store: store,
	}
}
