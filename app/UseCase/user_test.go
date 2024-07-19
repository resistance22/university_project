package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/aead/chacha20poly1305"
	repository "github.com/resistance22/university_project/Repository"
	token "github.com/resistance22/university_project/Token"
	utils "github.com/resistance22/university_project/Utils"
	validator "github.com/resistance22/university_project/Validator"
	"github.com/stretchr/testify/require"
)

func TestUserRegisterUseCase(t *testing.T) {
	repository := repository.NewUserMockRepository()
	tokenMaker, err := token.NewPasteoTokenMaker([]byte(utils.GenerateRandomString(chacha20poly1305.KeySize)))
	require.NoError(t, err)
	now := time.Now()
	useCase := NewUserUseCase(repository, tokenMaker)
	body := &validator.RegisterBody{
		FirstName: "Amin",
		LastName:  "Foroutan",
		UserName:  "amin_f",
		Password:  "password",
	}
	user, err := useCase.Register(context.Background(), body)

	require.NoError(t, err)
	require.NotEmpty(t, user, user.ID)
	require.NotEmpty(t, user, user.CreatedAt)
	require.WithinDuration(t, user.CreatedAt, now, time.Second)
	require.NotEqual(t, user.Password, body.Password)
	require.Equal(t, user.UserName, body.UserName)
	require.Equal(t, user.FirstName, body.FirstName)
	require.Equal(t, user.LastName, body.LastName)
}
