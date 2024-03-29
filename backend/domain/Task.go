package domain

import "database/sql"

/** バリデーション済みのタスク*/
type Task struct {
	TaskId      string `db:"task_id" json:"taskId"`
	UserId      string `db:"user_id"  json:"userId"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Status      string `db:"status" json:"status"` // NOT_STARTED, IN_PROGRESS, DONE
}

/** バリデーション前のタスク*/
type NotValidatedTask struct {
	TaskId      string       `db:"task_id" `
	UserId      string       `db:"user_id"`
	Title       string       `db:"title"`
	Description string       `db:"description"`
	Status      string       `db:"status"` // NOT_STARTED, IN_PROGRESS, DONE
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

type NoIdTask struct {
	UserId      string
	Title       string
	Description string
	Status      string
}

type TaskValidationError struct {
	Message string
}

func (te *TaskValidationError) Error() string {
	return "タスクとして適切な値ではありません。" + te.Message
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
		return nil, &TaskValidationError{Message: "ステータスとして不適な値"}
	}

	return &Task{
		TaskId:      task.TaskId,
		UserId:      task.UserId,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}, nil
}
