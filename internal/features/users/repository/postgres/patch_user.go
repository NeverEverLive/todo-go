package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
	core_postgres_pool "github.com/NeverEverLive/todo-go/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) PatchUser(
	ctx context.Context,
	userId string,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE todo_app.users
		SET
			first_name = $3,
			last_name = $4,
			phone_number = $5,
			version = version + 1
		WHERE id = $1 and version = $2
		RETURNING id, version, first_name, last_name, phone_number;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		userId,
		user.Version,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
	)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FirstName,
		&userModel.LastName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user concurrently accessed: %w",
				core_errors.ErrConflict,
			)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FirstName,
		userModel.LastName,
		userModel.PhoneNumber,
	), nil
}