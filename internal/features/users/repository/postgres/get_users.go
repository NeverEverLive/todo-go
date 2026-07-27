package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
)

func (r *UsersRepository) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
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
		ORDER BY id ASC
		LIMIT $1
		OFFSET $2;
	`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("select users: %w", err)
	}

	defer rows.Close()

	var userModels []UserModel
	for rows.Next() {
		var userModel UserModel
		err := rows.Scan(
			&userModel.ID,
			&userModel.Version,
			&userModel.FirstName,
			&userModel.LastName,
			&userModel.PhoneNumber,
		)
		if err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}

		userModels = append(userModels, userModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	userDomains := UserDomainsFromModels(userModels)

	return userDomains, nil
}
