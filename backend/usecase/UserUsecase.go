package usecase

import (
	"errors"
	"todo_app/domain"

	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepository domain.IUserRepository
}

func (u *UserUsecase) CreateUser(name, email string) error {
	id := uuid.New().String()

	user, err := domain.NewUser(id, name, email)
	if err != nil {
		return errors.New("ユーザーの作成に失敗しました")
	}

	u.userRepository.SaveUser(user)
	return nil
}
