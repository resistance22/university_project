package usecase

import (
	"context"

	entity "github.com/resistance22/university_project/Entity"
	validator "github.com/resistance22/university_project/Validator"
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

func (useCase *UserUseCase) Register(ctx context.Context, user *validator.RegisterBody) (*entity.User, error) {
	u, e := entity.NewUser(
		user.FirstName,
		user.LastName,
		user.UserName,
		user.Password,
	)

	if e != nil {
		return nil, e
	}

	err := useCase.repository.Register(ctx, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
