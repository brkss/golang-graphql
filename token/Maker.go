package token

import "time"

type Maker interface {
	CreateToken(userId string, duration time.Duration)(string, error);
	VerifyToken(token string)(*Payload, error)
}
