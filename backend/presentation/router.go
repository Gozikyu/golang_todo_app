package presentation

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouting(e *echo.Echo, taskHandler TaskHandler, userHandler UserHandler, loginHandler LoginHandler) {
	// ミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS()) // 必要に応じて細かいCORSの設定を行う

	r := e.Group("/restricted")

	r.Use(echojwt.JWT([]byte("secret")))

	//ログイン
	e.POST("/login", loginHandler.Login())

	//タスク関連のAPI
	e.GET("/:userId/tasks", taskHandler.GetTasks())
	r.POST("/:userId/tasks", taskHandler.CreateTask())
	e.PUT("/:userId/tasks/:taskId", taskHandler.UpdateTask())
	e.DELETE("/:userId/tasks/:taskId", taskHandler.DeleteTask())

	//ユーザー関連のAPI
	e.GET("/users", userHandler.GetUsers())
	e.GET("/users/:userId", userHandler.GetUser())
	e.POST("/users", userHandler.CreateUser())
	e.PUT("/users/:userId", userHandler.UpdateUser())
	e.DELETE("/users/:userId", userHandler.DeleteUser())
}
