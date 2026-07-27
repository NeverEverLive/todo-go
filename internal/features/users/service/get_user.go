package users_service

import (
	"context"
	"fmt"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
)


func (s *UsersService) GetUser(
	context context.Context,
	userId string,
) (domain.User, error) {
	user, err := s.usersRepository.GetUser(context, userId)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	return user, nil
}
