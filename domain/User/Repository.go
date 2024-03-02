package user

import (
	"context"

	db "github.com/resistance22/micorsales/db/sqlc"
)

type IUserRepository interface {
	createUser(user *User) (*User, error)
}

type UserRepository struct {
	dbtx db.DBTX
	ctx  context.Context
}

func NewUserRepository(dbtx db.DBTX, ctx context.Context) *UserRepository {
	return &UserRepository{dbtx, ctx}
}

func (r *UserRepository) createUser(user *User) (*User, error) {
	params := db.CreateUserParams{
		FirstName: user.firstName,
		LastName:  user.lastName,
		Email:     user.email,
		Password:  user.password,
		Role:      db.AppUserRoles(user.role),
	}
	queries := db.New(r.dbtx)
	createdUser, err := queries.CreateUser(r.ctx, params)

	if err != nil {
		return &User{}, err
	}

	return &User{
		id:        int(createdUser.ID),
		firstName: createdUser.FirstName,
		lastName:  createdUser.LastName,
		email:     createdUser.Email,
		password:  createdUser.Password,
		role:      UserRole(createdUser.Role),
	}, nil

}
