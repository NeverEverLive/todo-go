package users_service

import (
	"context"
	"fmt"
)


func (s *UsersService) DeleteUser(
	ctx context.Context,
	userId string,
) error {
	err := s.usersRepository.DeleteUser(ctx, userId)
	if err != nil {
		return fmt.Errorf("delete user in repository: %w", err)
	}

	return nil
}
