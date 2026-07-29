package tasks_service

import (
	"context"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
)

type TasksService struct {
	tasksRepository TasksRepository
}

type TasksRepository interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)
	GetTasks(
		ctx context.Context,
		userID *string,
		limit *int,
		offset *int,
	) ([]domain.Task, error)
	GetTask(
		ctx context.Context,
		id string,
	) (domain.Task, error)
	DeleteTask(
		ctx context.Context,
		id string,
	) error
	PatchTask(
		ctx context.Context,
		id string,
		task domain.Task,
	) (domain.Task, error)
}

func NewTasksService(
	tasksRepository TasksRepository,
) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
	}
}
