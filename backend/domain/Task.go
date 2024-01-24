package domain

import "errors"

type Task struct {
	taskId      string
	userId      string
	Title       string
	Description string
	Status      string // NOT_STARTED, IN_PROGRESS, DONE
}

var STATUS = []string{"NOT_STARTED", "IN_PROGRESS", "DONE"}

type ITaskRepository interface {
	GetTask(taskId string) *Task
	SaveTask(task *Task)
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
		taskId:      taskId,
		userId:      userId,
		Title:       title,
		Description: description,
		Status:      status,
	}, nil
}
