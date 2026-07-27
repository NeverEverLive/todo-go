package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
	"github.com/jackc/pgx/v5"
)


func (r *UsersRepository) GetUser(
	ctx context.Context,
	userId string,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT
			id,
			version,
			first_name,
			last_name,
			phone_number
		FROM todo_app.users
		WHERE id = $1;
	`

	row := r.pool.QueryRow(ctx, query, userId)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FirstName,
		&userModel.LastName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user not found: %w", core_errors.ErrNotFound)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FirstName,
		userModel.LastName,
		userModel.PhoneNumber,
	)
	return userDomain, nil
}