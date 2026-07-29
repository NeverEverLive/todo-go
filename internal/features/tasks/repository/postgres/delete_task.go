package tasks_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
)

func (r *TasksRepository) DeleteTask(
	ctx context.Context,
	id string,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM todo_app.tasks
	WHERE id = $1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec delete task: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("delete task: task not found: %w", core_errors.ErrNotFound)
	}

	return nil
}
