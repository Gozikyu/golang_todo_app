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

	/**以下は認証が必要なAPI*/

	//タスク関連のAPI
	r.GET("/:userId/tasks", taskHandler.GetTasks())
	r.POST("/:userId/tasks", taskHandler.CreateTask())
	r.PUT("/:userId/tasks/:taskId", taskHandler.UpdateTask())
	r.DELETE("/:userId/tasks/:taskId", taskHandler.DeleteTask())

	//ユーザー関連のAPI
	r.GET("/users", userHandler.GetUsers())
	r.GET("/users/:userId", userHandler.GetUser())
	r.POST("/users", userHandler.CreateUser())
	r.PUT("/users/:userId", userHandler.UpdateUser())
	r.DELETE("/users/:userId", userHandler.DeleteUser())
}
