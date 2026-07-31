package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
	core_postgres_pool "github.com/NeverEverLive/todo-go/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) PatchTask(
	ctx context.Context,
	id string,
	task domain.Task,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE todo_app.tasks
	SET
	    title = $1,
	    description = $2,
	    completed = $3,
	    completed_at = $4,
		version = version + 1
	WHERE id = $5 and version = $6
	RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id;`

	row := r.pool.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.CompletedAt,
		id,
		task.Version,
	)

	var taskModel TaskModel
	err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorUserID,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task concurrently accessed: %w",
				core_errors.ErrConflict,
			)
		}

		return domain.Task{}, fmt.Errorf("patch task: %w", err)
	}

	taskDomain := taskDomainFromModel(taskModel)

	return taskDomain, nil
}
