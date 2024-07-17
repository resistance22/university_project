package entity

import (
	"testing"

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
