package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"todo_app/domain"
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
	e.Use(middleware.CORS()) // 必要に応じて細かいCORSの設定を行う

	db, err := sqlx.Connect("postgres", "user=app password=password dbname=app_db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	tr := infra.NewTaskRepository(db)

	taskUsecase := usecase.NewTaskUsecase(tr)

	/**
	ユーザーの全タスク取得API

	リクエスト例
	curl  http://localhost:8888/1/tasks
	*/
	e.GET("/:userId/tasks", func(c echo.Context) error {
		userId := c.Param("userId")

		tasks, err := taskUsecase.GetUserTasks(userId)
		if err != nil {
			return errors.New(fmt.Sprintf("ユーザーの全タスク取得APIでエラーが発生しました。 userId: %v", userId))
		}

		return c.JSON(http.StatusOK, tasks)
	})

	/**
	タスク作成API

	リクエスト例
	curl -X POST \
	http://localhost:8888/your_user_id/tasks \
	-H 'Content-Type: application/json' \
	-d '{
		"UserId": "1",
		"Title": "Sample Task",
		"Description": "This is a sample task.",
		"Status": "NOT_STARTED"
		}'
	*/
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

	/**
	タスク削除API

	リクエストサンプル
	curl -X DELETE http://localhost:8888/1/tasks/101
	*/
	e.DELETE("/:userId/tasks/:taskId", func(c echo.Context) error {
		taskId := c.Param("taskId")

		err := taskUsecase.DeleteTask(taskId)
		if err != nil {
			fmt.Println(err)
			return errors.New(fmt.Sprintf("ユーザータスク削除APIでエラーが発生しました。 userId: %v", taskId))
		}

		return c.JSON(http.StatusOK, "success")

	})

	/**
	タスク更新API

	リクエストサンプル
	curl -X PUT \
	http://localhost:8888/your_user_id/tasks/your_task_id \
	-H 'Content-Type: application/json' \
	-d '{
		"TaskId": "102",
		"UserId": "2",
		"Title": "Updated Task",
		"Description": "This task has been updated.",
		"Status": "IN_PROGRESS"
		}'
	*/
	e.PUT("/:userId/tasks/:taskId", func(c echo.Context) error {
		var notValidatedTask domain.NotValidatedTask
		if err := c.Bind(&notValidatedTask); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"status": err.Error()})
		}

		task, err := domain.NewTask(notValidatedTask)
		if err != nil {
			return errors.New(fmt.Sprintf("ドメインのタスク型への変換に失敗しました。task: %v", task))
		}

		err = taskUsecase.UpdateTask(task)
		if err != nil {
			return errors.New(fmt.Sprintf("タスク更新APIでエラーが発生しました。task: %v", task))
		}

		return c.JSON(http.StatusOK, "success")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
