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

func NewTaskUsecase(r domain.ITaskRepository) TaskUsecase {
	return TaskUsecase{taskRepository: r}
}

func (u *TaskUsecase) CreateTask(newTask domain.NoIdTask) error {
	//uuidをアプリ側で発行する
	id := uuid.New().String()

	task, err := domain.NewTask(domain.NotValidatedTask{TaskId: id, UserId: newTask.UserId, Title: newTask.Title, Description: newTask.Description, Status: newTask.Status})
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
		return nil, fmt.Errorf("%vのタスク一覧を取得するのに失敗しました", userId)
	}

	return tasks, nil
}

func (u *TaskUsecase) UpdateTask(task *domain.Task) error {
	err := u.taskRepository.UpdateTask(task)
	if err != nil {
		return fmt.Errorf("%vのタスク更新に失敗しました", task)
	}

	return nil
}

func (u *TaskUsecase) DeleteTask(taskId string) error {
	err := u.taskRepository.DeleteTask(taskId)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("%vのタスク削除に失敗しました", taskId)
	}

	return nil
}
