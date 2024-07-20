package usecase

import (
	"context"
	"time"

	entity "github.com/resistance22/university_project/Entity"
	repository "github.com/resistance22/university_project/Repository"
	token "github.com/resistance22/university_project/Token"
	validator "github.com/resistance22/university_project/Validator"
)

type IUserUseCase interface {
	Register(ctx context.Context, user *validator.RegisterBody) (*entity.User, error)
	Login(ctx context.Context, user *validator.LoginBody) (string, error)
}

type UserUseCase struct {
	repository repository.IUserRepository
	tokenMaker token.TokenMaker
}

func NewUserUseCase(repo repository.IUserRepository, tokenMaker token.TokenMaker) IUserUseCase {
	return &UserUseCase{
		repo,
		tokenMaker,
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

func (useCase *UserUseCase) Login(ctx context.Context, dto *validator.LoginBody) (string, error) {
	user, err := useCase.repository.FindUserByUserName(ctx, dto.UserName)
	if err != nil {
		return "", err
	}

	token, err := useCase.tokenMaker.CreateToken(user.UserName, user.ID.String(), 15*time.Minute)
	if err != nil {
		return "", err
	}

	return token, nil

}
