package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	PayloadID uuid.UUID `json:"payload_id"`
	UserName  string    `json:"user_name"`
	UserID    string    `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(
	UserName string,
	UserID string,
	duration time.Duration,
) (*Payload, error) {
	payloadID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	payload := &Payload{
		PayloadID: payloadID,
		UserName:  UserName,
		UserID:    UserID,
		IssuedAt:  now,
		ExpiredAt: now.Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
