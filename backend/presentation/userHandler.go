package presentation

import (
	"fmt"
	"net/http"
	"strconv"
	"todo_app/domain"
	"todo_app/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetUsers() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	CreateUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
}

type userHandler struct {
	uu usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{uu: uu}
}

// TODO: ページネーションに対応させる
func (uh *userHandler) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

		users, err := uh.uu.GetUsers()
		if err != nil {
			return fmt.Errorf("全ユーザー取得APIでエラーが発生しました。")
		}

		totalCount := len(users)

		c.Response().Header().Set("X-Total-Count", strconv.Itoa(totalCount))
		return c.JSON(http.StatusOK, users)
	}
}

func (uh *userHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")

		user, err := uh.uu.GetUser(userId)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("ユーザー取得APIでエラーが発生しました。 userId: %v", userId)
		}

		return c.JSON(http.StatusOK, user)
	}
}

func (uh *userHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user domain.NoIdUser
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"status": err.Error()})
		}

		err := uh.uu.CreateUser(user)
		if err != nil {
			return fmt.Errorf("ユーザーの新規作成APIでエラーが発生しました。 user: %v", user)
		}

		return c.JSON(http.StatusOK, "success")
	}
}

func (uh *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")

		err := uh.uu.DeleteUser(userId)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("ユーザータスク削除APIでエラーが発生しました。 userId: %v", userId)
		}

		return c.JSON(http.StatusOK, "success")

	}
}

func (uh *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")
		var newUser domain.NoIdUser
		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"status": err.Error()})
		}

		user, err := domain.NewUser(domain.NotValidatedUser{UserId: userId, Name: newUser.Name, Email: newUser.Email})
		if err != nil {
			return fmt.Errorf("ドメインのユーザー型への変換に失敗しました。task: %v", user)
		}

		uu, err := uh.uu.UpdateUser(user)
		if err != nil {
			return fmt.Errorf("ユーザー更新APIでエラーが発生しました。task: %v", user)
		}

		return c.JSON(http.StatusOK, uu)
	}
}
