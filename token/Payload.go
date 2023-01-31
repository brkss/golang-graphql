package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken = errors.New("this token has been expired !")
var ErrInvalidToken = errors.New("this token is invalid !")

type Payload struct {
	ID string
	UserID string
	ExpireAt time.Time
	IssuedAt time.Time
}

func NewPayload(userID string, duration time.Duration)(*Payload) {
	return &Payload{
		ID: uuid.New().String(),
		UserID: userID,
		ExpireAt: time.Now().Add(duration),
		IssuedAt: time.Now(),
	}
}

func (p *Payload)Valid()(bool, error){

	if time.Now().After(p.ExpireAt){
		return false, ErrExpiredToken
	}

	return true, nil;
}
