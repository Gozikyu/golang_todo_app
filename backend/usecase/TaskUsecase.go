package usecase

import (
	"errors"
	"fmt"
	"todo_app/domain"

	"github.com/google/uuid"
)

type TaskUsecase struct {
	taskRepository domain.ITaskRepository
}

type NewTask struct {
	UserId      string
	Title       string
	Description string
	Status      string
}

func NewTaskUsecase(r domain.ITaskRepository) *TaskUsecase {
	return &TaskUsecase{taskRepository: r}
}

func (u *TaskUsecase) CreateTask(newTask NewTask) error {
	//uuidをアプリ側で発行する
	id := uuid.New().String()

	task, err := domain.NewTask(id, newTask.UserId, newTask.Title, newTask.Description, newTask.Status)
	if err != nil {
		return errors.New("タスクの作成に失敗しました")
	}

	err = u.taskRepository.SaveTask(task)
	if err != nil {
		return errors.New("タスクの登録に失敗しました")
	}
	return nil
}

func (u *TaskUsecase) GetUserTasks(userId string) ([]*domain.Task, error) {
	tasks, err := u.taskRepository.GetTasks(userId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%vのタスク一覧を取得するのに失敗しました", userId))
	}

	return tasks, nil
}
