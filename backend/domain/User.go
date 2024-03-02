package domain

import (
	"database/sql"
	"errors"
	"strings"
)

type User struct {
	UserId string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

/** バリデーション前のユーザー*/
type NotValidatedUser struct {
	UserId    string       `db:"user_id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type IUserRepository interface {
	GetUsers() ([]*User, error)
	// SaveUser(user *User) error
	// UpdateUser(user *User) error
	// DeleteUser(userId int) error
}

func NewUser(user NotValidatedUser) (*User, error) {
	if !(strings.Contains(user.Email, "@")) {
		return nil, errors.New("メールアドレスには@が含まれている必要があります")
	}

	return &User{
		UserId: user.UserId,
		Name:   user.Name,
		Email:  user.Email,
	}, nil
}
