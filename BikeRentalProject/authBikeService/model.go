package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const (
	signingKey = "afho457iewnflfp2452oivjnsqojr"
)

type tokenClaims struct {
	Role string
	Id   int
	jwt.StandardClaims
}

type authentication interface {
	decode(string) (*tokenClaims, error)
}
type AuthenticationService struct {
}

func (a *AuthenticationService) decode(token string) (*tokenClaims, error) {
	tokenType, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenType.Claims.(*tokenClaims); ok && tokenType.Valid {
		fmt.Println(claims)
		return claims, nil
	} else {

		return nil, err
	}
}
