package token

import (
	"testing"
	"time"

	"github.com/aead/chacha20poly1305"
	utils "github.com/resistance22/university_project/Utils"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasteoTokenMaker([]byte(utils.GenerateRandomString(chacha20poly1305.KeySize)))

	require.NoError(t, err)

	username := "username"
	userID := utils.GenerateRandomString(10)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(time.Minute)
	token, err := maker.CreateToken(username, userID, duration)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotEmpty(t, payload.PayloadID)
	require.Equal(t, username, payload.UserName)
	require.Equal(t, userID, payload.UserID)
	require.WithinDuration(t, payload.IssuedAt, issuedAt, time.Second)
	require.WithinDuration(t, payload.ExpiredAt, expiredAt, time.Second)

}

func TestInvalidKeySize(t *testing.T) {
	maker, err := NewPasteoTokenMaker([]byte(utils.GenerateRandomString(33)))
	require.Error(t, err)
	require.Nil(t, maker)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasteoTokenMaker([]byte(utils.GenerateRandomString(chacha20poly1305.KeySize)))
	require.NoError(t, err)

	token, err := maker.CreateToken("username", "2131231", -time.Minute)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidPasetoToken(t *testing.T) {
	maker, err := NewPasteoTokenMaker([]byte(utils.GenerateRandomString(chacha20poly1305.KeySize)))
	require.NoError(t, err)

	payload, err := maker.VerifyToken(utils.GenerateRandomString(10))

	require.Error(t, err)
	require.ErrorContains(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
