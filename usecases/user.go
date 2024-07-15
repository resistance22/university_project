package usecase

import (
	"context"

	entity "github.com/resistance22/university_project/Entity"
)

type UserRepository interface {
	Register(c context.Context, u *entity.User) error
}

type UserUseCase struct {
	repository UserRepository
}

func NewUserUseCase(repo UserRepository) *UserUseCase {
	return &UserUseCase{
		repository: repo,
	}
}

func (useCase *UserUseCase) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := useCase.repository.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
