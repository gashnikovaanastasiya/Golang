package main

import "golang.org/x/crypto/bcrypt"

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

type withHashing interface {
	CreateUser(*User) (string, error)
	validate(*User) error
}

type UserCreatorService struct {
	repository
}

func (u *UserCreatorService) CreateUser(user *User) (string, error) {
	var err error
	if err = u.validate(user); err == nil {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return "Hashing error", err
		}

		user.Password = string(hashedPass)
		str, err := u.repository.CreateUser(user)
		return str, err
	}
	return "Creation error", err
}

func (u *UserCreatorService) validate(user *User) error {
	if user.Name == "" || user.Password == "" {
		return &ValidationError{Message: "Missing fields\n"}
	}
	if _, err := u.repository.GetByName(user.Name); err == nil {

		return &ValidationError{Message: "User with this name already exists.Please change your name"}
	}

	return nil
}
