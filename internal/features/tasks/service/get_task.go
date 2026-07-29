package tasks_service

import (
	"context"
	"fmt"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
	"github.com/google/uuid"
)

func (s *TasksService) GetTask(
	ctx context.Context,
	id string,
) (domain.Task, error) {
	if err := uuid.Validate(id); err != nil {
		return domain.Task{}, fmt.Errorf(
			"`user_id` must be a valid uuid: %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	task, err := s.tasksRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf(
			"get task from repository: %w",
			err,
			)
	}

	return task, nil
}
