package domain

import (
	"database/sql"
	"errors"
	"strings"
)

type User struct {
	UserId   string `db:"user_id" json:"id"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

/** バリデーション前のユーザー*/
type NotValidatedUser struct {
	UserId    string       `db:"user_id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type NoIdUser struct {
	Name     string
	Email    string
	Password string
}

type IUserRepository interface {
	GetUsers() ([]*User, error)
	GetUser(userId string) (*User, error)
	SaveUser(user *User) error
	UpdateUser(user *User) (*User, error)
	DeleteUser(userId string) error
	GetUserByEmailAndPassword(email string, password string) (*User, error)
}

func NewUser(user NotValidatedUser) (*User, error) {
	if !(strings.Contains(user.Email, "@")) {
		return nil, errors.New("メールアドレスには@が含まれている必要があります")
	}

	return &User{
		UserId:   user.UserId,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
