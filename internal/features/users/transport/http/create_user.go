package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserRequest struct {
	FirstName 	string 	`json:"first_name"`
	LastName  	string 	`json:"last_name"`
	PhoneNumber *string `json:"phone_number"`
}

type CreateUserResponse struct {
	ID 			string 	`json:"id"`
	Version 	int 	`json:"version"`
	FirstName 	string 	`json:"first_name"`
	LastName 	string 	`json:"last_name"`
	PhoneNumber *string `json:"phone_number"`
	CreatedAt 	string 	`json:"created_at"`
	UpdatedAt 	string 	`json:"updated_at"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println(err)
	}
}