package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	entity "github.com/resistance22/university_project/Entity"
)

type UserMockRepository struct{}

func (repo *UserMockRepository) Register(ctx context.Context, user *entity.User) error {
	return nil
}

func (repo *UserMockRepository) FindUserByUserName(c context.Context, username string) (*entity.User, error) {
	uuid, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:        uuid,
		CreatedAt: time.Now(),
		FirstName: "amin",
		LastName:  "foroutan",
		UserName:  username,
		Password:  "1232331123123",
	}, nil
}

func NewUserMockRepository() IUserRepository {
	return &UserMockRepository{}
}
