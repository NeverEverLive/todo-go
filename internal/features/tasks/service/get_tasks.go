package tasks_service

import (
	"context"
	"fmt"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
	"github.com/google/uuid"
)

func (s *TasksService) GetTasks(
	ctx context.Context,
	userID *string,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	if limit != nil && *limit <= 0 {
		return nil, fmt.Errorf(
			"'limit' must be greater than zero: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset <= 0 {
		return nil, fmt.Errorf(
			"'offset' must be greater than zero: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if userID != nil {
		if err := uuid.Validate(*userID); err != nil {
			return nil, fmt.Errorf(
				"`user_id` must be a valid uuid: %v: %w",
				err,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	tasks, err := s.tasksRepository.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get tasks from repository: %w", err)
	}

	return tasks, nil
}
