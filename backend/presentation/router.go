package presentation

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouting(e *echo.Echo, taskHandler TaskHandler, userHandler UserHandler) {
	// ミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS()) // 必要に応じて細かいCORSの設定を行う

	//タスク関連のAPI
	e.GET("/:userId/tasks", taskHandler.GetTasks())
	e.POST("/:userId/tasks", taskHandler.CreateTask())
	e.PUT("/:userId/tasks/:taskId", taskHandler.UpdateTask())
	e.DELETE("/:userId/tasks/:taskId", taskHandler.DeleteTask())

	//ユーザー関連のAPI
	e.GET("/users", userHandler.GetUsers())
	e.GET("/users/:userId", userHandler.GetUser())
	e.POST("/users", userHandler.CreateUser())
	e.PUT("/users/:userId", userHandler.UpdateUser())
	e.DELETE("/users/:userId", userHandler.DeleteUser())
}
