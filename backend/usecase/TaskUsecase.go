package usecase

import (
	"errors"
	"todo_app/domain"

	"github.com/google/uuid"
)

type TaskUsecase struct {
	taskRepository domain.ITaskRepository
}

func (u *TaskUsecase) CreateTask(userId, title, description, status string) error {
	id := uuid.New().String()

	task, err := domain.NewTask(id, userId, title, description, status)
	if err != nil {
		return errors.New("タスクの作成に失敗しました")
	}

	u.taskRepository.SaveTask(task)
	return nil
}
