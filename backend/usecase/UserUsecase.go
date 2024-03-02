package usecase

import (
	"errors"
	"fmt"
	"todo_app/domain"
)

type UserUsecase struct {
	userRepository domain.IUserRepository
}

type NewUser struct {
	Name  string
	Email string
}

func NewUserUsecase(r domain.IUserRepository) UserUsecase {
	return UserUsecase{userRepository: r}
}

// func (u *UserUsecase) CreateUser(name, email string) error {
// 	id := uuid.New().String()

// 	user, err := domain.NewUser(domain.NotValidatedUser{UserId: id, Name: name, Email: email})
// 	if err != nil {
// 		return errors.New("ユーザーの作成に失敗しました")
// 	}

// 	u.userRepository.SaveUser(user)
// 	return nil
// }

func (u *UserUsecase) GetUsers() ([]*domain.User, error) {
	users, err := u.userRepository.GetUsers()
	if err != nil {
		fmt.Print(err)
		return nil, errors.New(fmt.Sprintf("ユーザー一覧を取得するのに失敗しました"))
	}

	return users, nil
}
