package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RequestEmail struct {
	Email string `validate:"required,min=1,max=140"`
}

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/**
	ユーザーの全タスクを取得する
	*/
	e.GET("/:userId/tasks", func(c echo.Context) error {
		return ""
	})

	/**
	タスクを追加する
	*/
	e.POST("/:userId/tasks", func(c echo.Context) error {
		return ""
	})

	/**
	タスクを削除する
	*/
	e.POST("/:userId/tasks/:taskId", func(c echo.Context) error {
		return ""
	})

	/**
	タスクを更新する
	*/
	e.PUT("/:userId/tasks/:taskId", func(c echo.Context) error {
		return ""
	})

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
