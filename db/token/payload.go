package token

import (
	"time"

	"github.com/google/uuid"
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    int64     `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific userId and duration
func NewPayload(userID int64, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
