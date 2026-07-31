package statistics_postgres_repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
)

func (r *StatisticsRepository) GetTasks(
	ctx context.Context,
	userID *string,
	dateFrom *time.Time,
	dateTo *time.Time,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var queryBuilder strings.Builder

	queryBuilder.WriteString(`
		SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
		FROM todo_app.tasks
	`)

	var args []any
	var conditions []string

	if userID != nil {
		conditions = append(conditions, fmt.Sprintf("author_user_id = $%d", len(args)+1))
		args = append(args, userID)
	}

	if dateFrom != nil {
		conditions = append(conditions, fmt.Sprintf("created_at >= $%d", len(args)+1))
		args = append(args, dateFrom)
	}

	if dateTo != nil {
		conditions = append(conditions, fmt.Sprintf("created_at <=  $%d", len(args)+1))
		args = append(args, dateTo)
	}

	if len(conditions) > 0 {
		queryBuilder.WriteString(" WHERE " + strings.Join(conditions, " AND "))
	}

	queryBuilder.WriteString(" ORDER BY id ASC")

	rows, err := r.pool.Query(ctx, queryBuilder.String(), args...)
	if err != nil {
		return nil, fmt.Errorf(
			"error getting tasks: %w",
			err,
		)
	}
	defer rows.Close()

	var tasksModels []TaskModel
	for rows.Next() {
		var taskModel TaskModel
		err := rows.Scan(
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
			return nil, fmt.Errorf(
				"scan tasks: %w",
				err,
			)
		}

		tasksModels = append(tasksModels, taskModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"error iterating rows: %w",
			err,
		)
	}

	taskDomains := tasksDomainsFromModels(tasksModels)

	return taskDomains, nil
}
