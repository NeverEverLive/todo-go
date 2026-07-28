package users_service

import (
	"context"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	GetUser(
		ctx context.Context,
		userId string,
	) (domain.User, error)
	DeleteUser(
		ctx context.Context,
		userId string,
	) (error)
	PatchUser(
		ctx context.Context,
		userId string,
		user domain.User,
	) (domain.User, error)
}

func NewUsersService(
	usersRepository UsersRepository,
) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}
