package usecase

import (
	"errors"
	"fmt"
	"todo_app/domain"

	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepository domain.IUserRepository
}

func NewUserUsecase(r domain.IUserRepository) UserUsecase {
	return UserUsecase{userRepository: r}
}

func (u *UserUsecase) CreateUser(newUser domain.NoIdUser) error {
	id := uuid.New().String()

	user, err := domain.NewUser(domain.NotValidatedUser{UserId: id, Name: newUser.Name, Email: newUser.Email})
	if err != nil {
		return errors.New("ユーザーの作成に失敗しました")
	}

	u.userRepository.SaveUser(user)
	return nil
}

func (u *UserUsecase) GetUsers() ([]*domain.User, error) {
	users, err := u.userRepository.GetUsers()
	if err != nil {
		fmt.Print(err)
		return nil, errors.New(fmt.Sprintf("ユーザー一覧を取得するのに失敗しました"))
	}

	return users, nil
}

func (u *UserUsecase) GetUser(userId string) (*domain.User, error) {
	user, err := u.userRepository.GetUser(userId)
	if err != nil {
		fmt.Print(err)
		return nil, errors.New(fmt.Sprintf("ユーザーを取得するのに失敗しました"))
	}

	return user, nil
}

func (u *UserUsecase) UpdateUser(user *domain.User) (*domain.User, error) {
	uu, err := u.userRepository.UpdateUser(user)
	if err != nil {
		fmt.Print(err)
		return nil, errors.New(fmt.Sprintf("ユーザー一覧を取得するのに失敗しました"))
	}

	return uu, nil
}

func (u UserUsecase) DeleteUser(userId string) error {
	err := u.userRepository.DeleteUser(userId)
	if err != nil {
		fmt.Println(err)
		return errors.New(fmt.Sprintf("%vのユーザー削除に失敗しました", userId))
	}

	return nil
}
