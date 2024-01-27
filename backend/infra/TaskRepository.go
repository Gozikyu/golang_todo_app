package infra

import (
	"database/sql"
	"errors"
	"fmt"
	"todo_app/domain"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TaskRepository struct {
	db *sqlx.DB
}

type NotValidatedTask struct {
	TaskId      string `db:"task_id"`
	UserId      string `db:"user_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Status      string `db:"status"` // NOT_STARTED, IN_PROGRESS, DONE
}

func NewTaskRepository(db *sqlx.DB) TaskRepository {
	return TaskRepository{db: db}
}

func (r TaskRepository) GetTask(taskId string) (*domain.Task, error) {
	var t NotValidatedTask
	err := r.db.Get(&t, "SELECT * FROM tasks WHERE task_id=$1", taskId)
	if err == sql.ErrNoRows {
		fmt.Printf("%vのタスクが見つかりませんでした", taskId)
		return nil, nil
	} else if err != nil {
		return nil, errors.New(fmt.Sprintf("%vのタスク取得時にエラーが発生しました", taskId))
	}

	task, err := domain.NewTask(t.TaskId, t.UserId, t.Title, t.Description, t.Status)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("%vドメインのタスク型に変換時にエラーが発生しました", t))
	}

	return task, nil
}

func (r TaskRepository) GetTasks(userId string) ([]*domain.Task, error) {
	var t []NotValidatedTask
	err := r.db.Select(&t, "SELECT * FROM tasks WHERE user_id=$1", userId)
	if err == sql.ErrNoRows {
		fmt.Printf("%vのタスクが見つかりませんでした", userId)
		return nil, nil
	} else if err != nil {
		return nil, errors.New(fmt.Sprintf("%vのタスク取得時にエラーが発生しました", userId))
	}

	tasks := []*domain.Task{}
	for _, v := range t {
		task, err := domain.NewTask(v.TaskId, v.UserId, v.Title, v.Description, v.Status)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%vドメインのタスク型に変換時にエラーが発生しました", t))
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r TaskRepository) SaveTask(task *domain.Task) error {
	_, err := r.db.NamedExec(`INSERT INTO tasks (task_id, user_id, title, description, status) VALUES (:task_id, :user_id, :title, :description, :status)`, task)

	if err != nil {
		return errors.New(fmt.Sprintf("%vタスクのDB登録時にエラーが発生しました", task))
	}

	return nil
}
