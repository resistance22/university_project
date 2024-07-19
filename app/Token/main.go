package token

import (
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token provided")
	ErrExpiredToken = errors.New("token has been expired")
)

type TokenMaker interface {
	CreateToken(username string, userID string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
