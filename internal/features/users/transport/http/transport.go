package users_transport_http

import (
	"net/http"

	core_http_server "github.com/NeverEverLive/todo-go/internal/core/transport/http/server"
)


type UsersHTTPHandler struct {
	usersService UserService
}


type UserService interface {
}

func NewUsersHTTPHandler(
	usersService UserService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route {
		{
			Method: http.MethodPost,
			Path: "/users",
			Handler: h.CreateUser,
		},
	}
}
