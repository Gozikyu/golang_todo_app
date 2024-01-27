package domain

import (
	"errors"
)

type Task struct {
	TaskId      string `db:"task_id"`
	UserId      string `db:"user_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Status      string `db:"status"` // NOT_STARTED, IN_PROGRESS, DONE
}

var STATUS = []string{"NOT_STARTED", "IN_PROGRESS", "DONE"}

type ITaskRepository interface {
	GetTask(taskId string) (*Task, error)
	GetTasks(userId string) ([]*Task, error)
	SaveTask(task *Task) error
}

func NewTask(taskId string, userId string, title string, description string, status string) (*Task, error) {
	found := false
	for _, v := range STATUS {
		if status == v {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("ステータスとして定義されていない文字列です")
	}

	return &Task{
		TaskId:      taskId,
		UserId:      userId,
		Title:       title,
		Description: description,
		Status:      status,
	}, nil
}
