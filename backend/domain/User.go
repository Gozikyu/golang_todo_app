package domain

import (
	"errors"
	"strings"
)

type User struct {
	userId string
	Name   string
	Email  string
}

type IUserRepository interface {
	GetUser(userId int) *User
	SaveUser(user *User)
}

func NewUser(userId string, name string, email string) (*User, error) {
	if !(strings.Contains(email, "@")) {
		return nil, errors.New("メールアドレスには@が含まれている必要があります")
	}

	return &User{
		userId: userId,
		Name:   name,
		Email:  email,
	}, nil
}
