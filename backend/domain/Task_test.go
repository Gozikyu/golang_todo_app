package domain_test

import (
	"errors"
	"testing"
	"todo_app/domain"
)

func TestNewTask(t *testing.T) {
	type args struct {
		t domain.NotValidatedTask
	}

	normalCase := []struct {
		name string
		args domain.NotValidatedTask
		want domain.Task
	}{
		{name: "NOT_STARTEDの場合",
			args: domain.NotValidatedTask{TaskId: "1", UserId: "1", Title: "タイトル", Description: "本文", Status: "NOT_STARTED"},
			want: domain.Task{TaskId: "1", UserId: "1", Title: "タイトル", Description: "本文", Status: "NOT_STARTED"},
		},
		{name: "IN_PROGRESSの場合",
			args: domain.NotValidatedTask{TaskId: "1", UserId: "1", Title: "タイトル", Description: "本文", Status: "IN_PROGRESS"},
			want: domain.Task{TaskId: "1", UserId: "1", Title: "タイトル", Description: "本文", Status: "IN_PROGRESS"},
		},
		{name: "DONEの場合",
			args: domain.NotValidatedTask{TaskId: "1", UserId: "1", Title: "タイトル", Description: "本文", Status: "DONE"},
			want: domain.Task{TaskId: "1", UserId: "1", Title: "タイトル", Description: "本文", Status: "DONE"},
		},
	}

	for _, tt := range normalCase {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.NewTask(tt.args)

			if err != nil {
				t.Errorf("error: %v", err)
			}

			if *got != tt.want {
				t.Errorf("got: %v, want: %v", *got, tt.want)
			}
		})
	}

	errorCase := []struct {
		name string
		args domain.NotValidatedTask
		want error
	}{
		{name: "不適切な文字列の場合",
			args: domain.NotValidatedTask{TaskId: "1", UserId: "1", Title: "タイトル", Description: "本文", Status: "INVALID"},
			want: &domain.TaskValidationError{Message: "ステータスとして不適な値"},
		},
	}

	for _, tt := range errorCase {
		t.Run(tt.name, func(t *testing.T) {
			var validationErr *domain.TaskValidationError
			_, err := domain.NewTask(tt.args)

			if !errors.As(err, &validationErr) {
				t.Errorf("error: %v, want: %v", err, tt.want)
			}
			if err.Error() != tt.want.Error() {
				t.Errorf("error: %v, want: %v", err, tt.want)
			}
		})
	}

}
