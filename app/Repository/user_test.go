package repository

import (
	"context"
	"testing"
	"time"

	entity "github.com/resistance22/university_project/Entity"
	utils "github.com/resistance22/university_project/Utils"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	user, err := entity.NewUser(
		"Amin",
		"Foroutan",
		utils.GenerateRandomString(10),
		"ThePassword",
	)

	require.NoError(t, err)

	err = UserRepo.Register(context.Background(), user)
	require.NoError(t, err)
}

func TestFindUserByUserName(t *testing.T) {
	userName := utils.GenerateRandomString(10)
	password := "complexPass"
	user, err := entity.NewUser(
		"Amin",
		"Foroutan",
		userName,
		password,
	)

	require.NoError(t, err)
	err = UserRepo.Register(context.Background(), user)
	require.NoError(t, err)

	foundUser, err := UserRepo.FindUserByUserName(context.Background(), userName)
	require.NoError(t, err)
	require.NotEmpty(t, foundUser)
	require.Equal(t, user.ID, foundUser.ID)
	require.WithinDuration(t, user.CreatedAt, foundUser.CreatedAt, time.Second)
	require.Equal(t, user.FirstName, foundUser.FirstName)
	require.Equal(t, user.LastName, foundUser.LastName)
	require.Equal(t, user.Password, foundUser.Password)
	require.Equal(t, user.UserName, foundUser.UserName)

}
