package users_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
)

func (r *UsersRepository) DeleteUser(
	ctx context.Context,
	userId string,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		DELETE FROM todo_app.users
		WHERE id = $1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, userId)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user not found: %w", core_errors.ErrNotFound)
	}

	return nil
}
