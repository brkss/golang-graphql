package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)


type PasetoMaker struct {
	paseto 		*paseto.V2
	symtricKey 	string
}


func NewPasetoMaker(symtrictKey string)(Maker, error) {
	if len(symtrictKey) < chacha20poly1305.KeySize{
		return nil, fmt.Errorf("Invalid symtrickey size must be atleast %d chars", chacha20poly1305.KeySize)
	}	
	maker := &PasetoMaker{
		paseto: paseto.NewV2(),
		symtricKey: symtrictKey,
	}
	return maker, nil
}

func (p *PasetoMaker)CreateToken(userId string, duration time.Duration)(string, error){
	
	payload := NewPayload(userId, duration)
	token, err := p.paseto.Encrypt([]byte(p.symtricKey), payload, nil)
	if err != nil {
		return "", err;
	}

	return token, nil 	
}

func (p *PasetoMaker)VerifyToken(token string)(*Payload, error){
	var payload Payload
	err := p.paseto.Decrypt(token, []byte(p.symtricKey), &payload, nil)

	if err != nil {
		return nil, ErrInvalidToken 
	} 

	valid, err := payload.Valid()
	if !valid {
		return nil, err
	}

	return &payload, nil

}
