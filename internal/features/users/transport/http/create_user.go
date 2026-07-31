package users_transport_http

import (
	"net/http"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
	core_logger "github.com/NeverEverLive/todo-go/internal/core/logger"
	core_http_request "github.com/NeverEverLive/todo-go/internal/core/transport/http/request"
	core_http_response "github.com/NeverEverLive/todo-go/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FirstName   string  `json:"first_name" validate:"required,min=3,max=100" example:"Ivan"`
	LastName    string  `json:"last_name" validate:"required,min=3,max=100" example:"Ivanov"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+" example:"+79998887766"`
}

type CreateUserResponse UserDTOResponse

// CreateUser 	godoc
// @Summary 	Create user
// @Description Create new user in system
// @Tags 		users
// @Accept		json
// @Produce		json
// @Param		request body CreateUserRequest true "CreateUser request body"
// @Success 	201 {object} CreateUserResponse "Successfully created user"
// @Failure		400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure		500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router		/users [post]
func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	userDomain := domainFromDTO(request)

	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := CreateUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FirstName, dto.LastName, dto.PhoneNumber)
}
