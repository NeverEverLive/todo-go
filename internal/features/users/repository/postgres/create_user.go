package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO todo_app.users (
			first_name,
			last_name,
			phone_number
		) VALUES (
			$1, $2, $3
		) RETURNING id, version, first_name, last_name, phone_number;
	`

	row := r.pool.QueryRow(ctx, query, user.FirstName, user.LastName, user.PhoneNumber)
	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FirstName,
		&userModel.LastName,
		&userModel.PhoneNumber,
	)
	if err != nil {
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
