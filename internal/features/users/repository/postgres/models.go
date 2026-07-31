package users_postgres_repository

import "github.com/NeverEverLive/todo-go/internal/core/domain"

type UserModel struct {
	ID          string
	Version     int
	FirstName   string
	LastName    string
	PhoneNumber *string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))
	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.FirstName,
			user.LastName,
			user.PhoneNumber,
		)
	}

	return userDomains
}
