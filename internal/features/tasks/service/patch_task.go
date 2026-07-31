package tasks_service

import (
	"context"
	"fmt"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
	"github.com/google/uuid"
)

func (s *TasksService) TaskPatch(
	ctx context.Context,
	id string,
	patch domain.TaskPatch,
) (domain.Task, error) {
	if err := uuid.Validate(id); err != nil {
		return domain.Task{}, fmt.Errorf(
			"`user_id` must be a valid uuid: %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	task, err := s.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task: %w", err)
	}

	if err := task.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("apply task patch: %w", err)
	}

	patchedTask, err := s.tasksRepository.PatchTask(ctx, id, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("patch task: %w", err)
	}

	return patchedTask, nil
}
