package presentation

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	"todo_app/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type LoginHandler interface {
	Login() echo.HandlerFunc
}

type loginHandler struct {
	lu usecase.LoginUsecase
}

func NewLoginHandler(lu usecase.LoginUsecase) LoginHandler {
	return &loginHandler{lu: lu}
}

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func (lh *loginHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		type authInfo struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var ai authInfo
		if err := c.Bind(&ai); err != nil {
			fmt.Print("ログインに失敗しました")
			return echo.ErrUnauthorized
		}

		user, err := lh.lu.FindUser(ai.Email, ai.Password)

		if err != nil || user == nil {
			return echo.ErrUnauthorized
		}

		// Set custom claims
		claims := &jwtCustomClaims{
			user.UserId,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return errors.New("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return errors.New("failed to cast claims as jwt.MapClaims")
	}
	return c.JSON(http.StatusOK, claims)
}
