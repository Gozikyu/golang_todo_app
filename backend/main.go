package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"todo_app/infra"
	"todo_app/usecase"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := sqlx.Connect("postgres", "user=app password=password dbname=app_db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	tr := infra.NewTaskRepository(db)

	taskUsecase := usecase.NewTaskUsecase(tr)

	/**
	ユーザーの全タスクを取得する
	*/
	e.GET("/:userId/tasks", func(c echo.Context) error {
		userId := c.Param("userId")

		fmt.Println(userId)

		// var t []infra.NotValidatedTask
		// db.Select(&t, "SELECT * FROM tasks WHERE user_id=$1", userId)

		tasks, err := taskUsecase.GetUserTasks(userId)
		if err != nil {
			return errors.New(fmt.Sprintf("ユーザーの全タスク取得APIでエラーが発生しました。 userId: %v", userId))
		}

		return c.JSON(http.StatusOK, tasks)
	})

	/**
	タスクを追加する
	リクエスト例: curl -X POST http://localhost:8888/your_user_id/tasks -H 'Content-Type: application/json' -d '{"UserId": "1", "Title": "Sample Task", "Description": "This is a sample task.", "Status": "NOT_STARTED"}'
	// */
	e.POST("/:userId/tasks", func(c echo.Context) error {
		var task usecase.NewTask
		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"status": err.Error()})
		}

		err := taskUsecase.CreateTask(task)
		if err != nil {
			return errors.New(fmt.Sprintf("タスクの新規作成APIでエラーが発生しました。 task: %v", task))
		}

		return c.JSON(http.StatusOK, "success")
	})

	// /**
	// タスクを削除する
	// */
	// e.POST("/:userId/tasks/:taskId", func(c echo.Context) error {
	// 	return ""
	// })

	// /**
	// タスクを更新する
	// */
	// e.PUT("/:userId/tasks/:taskId", func(c echo.Context) error {
	// 	return ""
	// })

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
