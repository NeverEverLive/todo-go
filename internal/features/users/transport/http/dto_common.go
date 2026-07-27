package users_transport_http

import "github.com/NeverEverLive/todo-go/internal/core/domain"

type UserDTOResponse struct {
	ID 			string 	`json:"id"`
	Version 	int 	`json:"version"`
	FirstName 	string 	`json:"first_name"`
	LastName 	string 	`json:"last_name"`
	PhoneNumber *string `json:"phone_number"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID: user.ID,
		Version: user.Version,
		FirstName: user.FirstName,
		LastName: user.LastName,
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