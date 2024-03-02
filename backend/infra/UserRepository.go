package infra

import (
	"database/sql"
	"errors"
	"fmt"
	"todo_app/domain"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{db: db}
}

func (r UserRepository) GetUsers() ([]*domain.User, error) {
	var u []domain.NotValidatedUser
	err := r.db.Select(&u, "SELECT * FROM users WHERE deleted_at IS NULL")
	if err == sql.ErrNoRows {
		fmt.Printf("ユーザーが見つかりませんでした")
		return nil, nil
	} else if err != nil {
		fmt.Print(err)
		return nil, errors.New(fmt.Sprintf("ユーザー一覧の取得時にエラーが発生しました"))
	}

	users := []*domain.User{}
	for _, v := range u {
		user, err := domain.NewUser(v)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%vドメインのユーザー型に変換時にエラーが発生しました", u))
		}
		users = append(users, user)
	}

	return users, nil
}
