package infra

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
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

func (r UserRepository) GetUser(userId string) (*domain.User, error) {
	var u domain.NotValidatedUser
	err := r.db.Get(&u, "SELECT * FROM users WHERE user_id=$1 AND deleted_at IS NULL", userId)
	if err == sql.ErrNoRows {
		fmt.Printf("%vのユーザーが見つかりませんでした", userId)
		return nil, nil
	} else if err != nil {
		fmt.Print(err)
		return nil, errors.New(fmt.Sprintf("%vのユーザー取得時にエラーが発生しました", userId))
	}

	user, err := domain.NewUser(u)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%vドメインのユーザー型に変換時にエラーが発生しました", u))
	}
	return user, nil
}

func (r UserRepository) SaveUser(user *domain.User) error {
	_, err := r.db.NamedExec(`INSERT INTO users (user_id, name, email) VALUES (:user_id, :name, :email)`, user)

	if err != nil {
		return errors.New(fmt.Sprintf("%vユーザーのDB登録時にエラーが発生しました", user))
	}

	return nil
}

func (r UserRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	_, err := r.db.NamedExec(`UPDATE users SET user_id = :user_id, name = :name, email = :email WHERE user_id = :user_id`, user)

	if err != nil {
		fmt.Print(err)
		return nil, errors.New(fmt.Sprintf("%vユーザーの更新時にエラーが発生しました", user))
	}

	// 更新後のユーザーレコードを取得
	var updatedUser domain.NotValidatedUser
	err = r.db.Get(&updatedUser, "SELECT * FROM users WHERE user_id =$1", user.UserId)
	if err != nil {
		fmt.Print(err)
		return nil, errors.New(fmt.Sprintf("更新後のユーザーレコードの取得時にエラーが発生しました"))
	}

	uu, err := domain.NewUser(updatedUser)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%vドメインのユーザー型に変換時にエラーが発生しました", updatedUser))
	}
	return uu, nil
}

func (r UserRepository) DeleteUser(userId string) error {
	// deletedAt に現在時刻を設定
	deletedAt := time.Now()

	_, err := r.db.Exec(`UPDATE users SET deleted_at = $1 WHERE user_id = $2`, deletedAt, userId)

	if err != nil {
		return errors.New(fmt.Sprintf("%vユーザーの削除時にエラーが発生しました", userId))
	}

	return nil
}

func (r UserRepository) GetUserByEmailAndPassword(email string, password string) (*domain.User, error) {
	var u domain.NotValidatedUser
	err := r.db.Get(&u, "SELECT * FROM users WHERE email=$1 AND password=$2 AND deleted_at IS NULL", email, password)
	if err == sql.ErrNoRows {
		fmt.Printf("%vのユーザーが見つかりませんでした", email)
		return nil, nil
	} else if err != nil {
		fmt.Print(err)
		return nil, errors.New(fmt.Sprintf("%vのユーザー取得時にエラーが発生しました", email))
	}

	user, err := domain.NewUser(u)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%vドメインのユーザー型に変換時にエラーが発生しました", u))
	}
	return user, nil
}
