package main

import (
	"golang.org/x/crypto/bcrypt"
)

type authorization interface {
	Authorize(string, string) (string, int, error)
}
type AuthorizationService struct {
	repository
}

func (a *AuthorizationService) Authorize(name, password string) (string, int, error) {
	u, err := a.repository.GetByName(name)
	if err != nil {
		return "", -1, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return "", -1, err
	}
	return u.Role, u.UserId, nil
}
