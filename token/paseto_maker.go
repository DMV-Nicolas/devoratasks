package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"

	"github.com/o1egl/paseto"
)

// PasetoMaker is a Paseto maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetrickey []byte
}

// NewPasetoMaker crates a new PasetoMaker
func NewPasetoMaker(symmetrickey string) (Maker, error) {
	if len(symmetrickey) != chacha20poly1305.KeySize {
		err := fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
		return nil, err
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetrickey: []byte(symmetrickey),
	}

	return maker, nil
}

// CreateToken creates a new token for the specific username and duration
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", nil
	}

	return maker.paseto.Encrypt(maker.symmetrickey, payload, nil)
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetrickey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
