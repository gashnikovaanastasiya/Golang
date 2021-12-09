package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "afho457iewnflfp2452oivjnsqojr"
	takenTTL   = 1 * time.Hour
)

type tokenClaims struct {
	Role string
	Id   int
	jwt.StandardClaims
}
type tokenMaker interface {
	CreateToken(int, string) (string, error)
}

type TokenMaker struct {
	repository
}

func (t *TokenMaker) CreateToken(Id int, Role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(takenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, Role: Role, Id: Id,
	})

	return token.SignedString([]byte(signingKey))
}
