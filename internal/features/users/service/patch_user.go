package users_service

import (
	"context"
	"fmt"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
)


func (s *UsersService) PatchUser(
	ctx context.Context,
	userId string,
	userPatch domain.UserPatch,
) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, userId)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	if err := user.ApplyPatch(userPatch); err != nil {
		return domain.User{}, fmt.Errorf("apply patch to user: %w", err)
	}

	patchedUser, err := s.usersRepository.PatchUser(ctx, userId, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("patch user: %w", err)
	}

	return patchedUser, nil
}
