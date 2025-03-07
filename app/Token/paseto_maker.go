package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoTokenMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasteoTokenMaker(symmetricKey []byte) (TokenMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoTokenMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: symmetricKey,
	}

	return maker, nil
}

func (maker *PasetoTokenMaker) CreateToken(username string, userID string, duration time.Duration) (string, error) {
	payload, err := NewPayload(
		username,
		userID,
		duration,
	)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}
func (maker *PasetoTokenMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
