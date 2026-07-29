package tasks_service

import (
	"context"
	"fmt"

	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
	"github.com/google/uuid"
)

func (s *TasksService) DeleteTask(
	ctx context.Context,
	id string,
) error {
	if err := uuid.Validate(id); err != nil {
		return fmt.Errorf(
			"`user_id` must be a valid uuid: %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	if err := s.tasksRepository.DeleteTask(ctx, id); err != nil {
		return fmt.Errorf(
			"failed to delete from repository: %w",
			err,
		)
	}

	return nil
}
