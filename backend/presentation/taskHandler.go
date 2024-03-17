package presentation

import (
	"fmt"
	"net/http"
	"todo_app/domain"
	"todo_app/usecase"

	"github.com/labstack/echo/v4"
)

type TaskHandler interface {
	GetTasks() echo.HandlerFunc
	CreateTask() echo.HandlerFunc
	UpdateTask() echo.HandlerFunc
	DeleteTask() echo.HandlerFunc
}

type taskHandler struct {
	tu usecase.TaskUsecase
}

func NewTaskHandler(tu usecase.TaskUsecase) TaskHandler {
	return &taskHandler{tu: tu}
}

func (th *taskHandler) GetTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")

		tasks, err := th.tu.GetUserTasks(userId)
		if err != nil {
			return fmt.Errorf("ユーザーの全タスク取得APIでエラーが発生しました。 userId: %v", userId)
		}

		return c.JSON(http.StatusOK, tasks)
	}

}

func (th *taskHandler) CreateTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		var task domain.NoIdTask
		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"status": err.Error()})
		}

		userId, e := GetUserIdFromToken(c)
		if e != nil {
			fmt.Print(e)
			return fmt.Errorf("トークンからのuserId取得に失敗しました task: %v", task)
		}

		if userId != task.UserId {
			return fmt.Errorf("トークンのuserIdと新規作成するタスクのuserIdが異なります: %v != %v", userId, task.UserId)
		}

		err := th.tu.CreateTask(task)
		if err != nil {
			return fmt.Errorf("タスクの新規作成APIでエラーが発生しました。 task: %v", task)

		}

		return c.JSON(http.StatusOK, "success")
	}
}

func (th *taskHandler) DeleteTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		taskId := c.Param("taskId")

		err := th.tu.DeleteTask(taskId)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("ユーザータスク削除APIでエラーが発生しました。 userId: %v", taskId)
		}

		return c.JSON(http.StatusOK, "success")
	}
}

func (th *taskHandler) UpdateTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		var notValidatedTask domain.NotValidatedTask
		if err := c.Bind(&notValidatedTask); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"status": err.Error()})
		}

		task, err := domain.NewTask(notValidatedTask)
		if err != nil {
			return fmt.Errorf("ドメインのタスク型への変換に失敗しました。task: %v", task)
		}

		err = th.tu.UpdateTask(task)
		if err != nil {
			return fmt.Errorf("タスク更新APIでエラーが発生しました。task: %v", task)
		}

		return c.JSON(http.StatusOK, "success")
	}
}
