package domain

import (
	"errors"
)

/** バリデーション済みのタスク*/
type Task struct {
	TaskId      string `db:"task_id"`
	UserId      string `db:"user_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Status      string `db:"status"` // NOT_STARTED, IN_PROGRESS, DONE
}

/** バリデーション前のタスク*/
type NotValidatedTask struct {
	TaskId      string `db:"task_id"`
	UserId      string `db:"user_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Status      string `db:"status"` // NOT_STARTED, IN_PROGRESS, DONE
}

// TODO: もっと上手く定義できそう
var STATUS = []string{"NOT_STARTED", "IN_PROGRESS", "DONE"}

type ITaskRepository interface {
	GetTask(taskId string) (*Task, error)
	GetTasks(userId string) ([]*Task, error)
	SaveTask(task *Task) error
	UpdateTask(task *Task) error
	DeleteTask(taskId string) error
}

func NewTask(task NotValidatedTask) (*Task, error) {
	found := false
	for _, v := range STATUS {
		if task.Status == v {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("ステータスとして定義されていない文字列です")
	}

	return &Task{
		TaskId:      task.TaskId,
		UserId:      task.UserId,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}, nil
}
