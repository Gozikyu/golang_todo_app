package infra

import (
	"database/sql"
	"fmt"
	"time"
	"todo_app/domain"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) TaskRepository {
	return TaskRepository{db: db}
}

func (r TaskRepository) GetTask(taskId string) (*domain.Task, error) {
	var t domain.NotValidatedTask
	err := r.db.Get(&t, "SELECT * FROM tasks WHERE task_id=$1 AND deleted_at IS NULL", taskId)
	if err == sql.ErrNoRows {
		fmt.Printf("%vのタスクが見つかりませんでした", taskId)
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("%vのタスク取得時にエラーが発生しました", taskId)
	}

	task, err := domain.NewTask(t)

	if err != nil {
		return nil, fmt.Errorf("%vドメインのタスク型に変換時にエラーが発生しました", t)
	}

	return task, nil
}

func (r TaskRepository) GetTasks(userId string) ([]*domain.Task, error) {
	var t []domain.NotValidatedTask
	err := r.db.Select(&t, "SELECT * FROM tasks WHERE user_id=$1 AND deleted_at IS NULL", userId)
	if err == sql.ErrNoRows {
		fmt.Printf("%vのタスクが見つかりませんでした", userId)
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("%vのタスク取得時にエラーが発生しました", userId)
	}

	tasks := []*domain.Task{}
	for _, v := range t {
		task, err := domain.NewTask(v)
		if err != nil {
			return nil, fmt.Errorf("%vドメインのタスク型に変換時にエラーが発生しました", t)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r TaskRepository) SaveTask(task *domain.Task) error {
	_, err := r.db.NamedExec(`INSERT INTO tasks (task_id, user_id, title, description, status) VALUES (:task_id, :user_id, :title, :description, :status)`, task)

	if err != nil {
		return fmt.Errorf("%vタスクのDB登録時にエラーが発生しました", task)
	}

	return nil
}

func (r TaskRepository) UpdateTask(task *domain.Task) error {
	_, err := r.db.NamedExec(`UPDATE tasks SET user_id = :user_id, title = :title, description = :description, status = :status WHERE task_id = :task_id`, task)

	if err != nil {
		return fmt.Errorf("%vタスクの更新時にエラーが発生しました", task)
	}

	return nil
}

func (r TaskRepository) DeleteTask(taskId string) error {
	// deletedAt に現在時刻を設定
	deletedAt := time.Now()

	_, err := r.db.Exec(`UPDATE tasks SET deleted_at = $1 WHERE task_id = $2`, deletedAt, taskId)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("%vタスクの削除時にエラーが発生しました", taskId)
	}

	return nil
}
