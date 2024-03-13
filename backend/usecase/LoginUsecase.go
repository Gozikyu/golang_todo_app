package usecase

import (
	"errors"
	"fmt"
	"todo_app/domain"
)

type LoginUsecase struct {
	userRepository domain.IUserRepository
}

func NewLoginUsecase(r domain.IUserRepository) LoginUsecase {
	return LoginUsecase{userRepository: r}
}

func (u *LoginUsecase) FindUser(email string, password string) (*domain.User, error) {
	user, err := u.userRepository.GetUserByEmailAndPassword(email, password)
	if err != nil {
		fmt.Print(err)
		return nil, errors.New(fmt.Sprintf("ユーザーを取得するのに失敗しました"))
	}

	return user, nil
}
