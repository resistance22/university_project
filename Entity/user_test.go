package entity

import (
	"testing"

	utils "github.com/resistance22/university_project/Utils"
	"github.com/stretchr/testify/require"
)

func TestUserEntity(t *testing.T) {
	user, err := NewUser(
		"Amin",
		"Foroutan",
		"amin_foroutan",
		"Password",
	)

	require.NoError(t, err)
	require.NotEmpty(t, user.CreatedAt)
	require.NotEmpty(t, user.ID)
	require.Equal(t, user.FirstName, "Amin")
	require.Equal(t, user.LastName, "Foroutan")
	require.Equal(t, user.UserName, "amin_foroutan")
	require.NotEqual(t, user.Password, "Password")
}

func TestUserEntityError(t *testing.T) {
	user, err := NewUser(
		"Amin",
		"Foroutan",
		"amin_foroutan",
		utils.GenerateRandomString(73),
	)

	require.Nil(t, user)
	require.Error(t, err)
	require.ErrorContains(t, err, ErrPasswordHashFailed.Error())

}
