package statistics_postgres_repository

import (
	"time"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
)

type TaskModel struct {
	ID           string
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CompletedAt  *time.Time
	CreatedAt    time.Time
	AuthorUserID string
}

func tasksDomainsFromModels(tasks []TaskModel) []domain.Task {
	taskDomains := make([]domain.Task, len(tasks))

	for i, task := range tasks {
		taskDomains[i] = taskDomainFromModel(task)
	}

	return taskDomains
}

func taskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserID,
	)
}
