package users_transport_http

import "github.com/NeverEverLive/todo-go/internal/core/domain"

type UserDTOResponse struct {
	ID          string  `json:"id" example:"99562482-baa9-4b3b-9115-ca65aae323a3"`
	Version     int     `json:"version" example:"3"`
	FirstName   string  `json:"first_name" example:"Ivan"`
	LastName    string  `json:"last_name" example:"Ivanov"`
	PhoneNumber *string `json:"phone_number" example:"+79998887766"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))
	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}
