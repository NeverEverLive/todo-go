package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
	core_logger "github.com/NeverEverLive/todo-go/internal/core/logger"
	core_http_request "github.com/NeverEverLive/todo-go/internal/core/transport/http/request"
	core_http_response "github.com/NeverEverLive/todo-go/internal/core/transport/http/response"
	core_http_types "github.com/NeverEverLive/todo-go/internal/core/transport/http/types"
	core_http_utils "github.com/NeverEverLive/todo-go/internal/core/transport/http/utils"
)


type PatchUserRequest struct {
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
	FirstName   core_http_types.Nullable[string] `json:"first_name"`
	LastName    core_http_types.Nullable[string] `json:"last_name"`
}


func (r *PatchUserRequest) Validate() error {
	if r.FirstName.Set {
		if r.FirstName.Value == nil {
			return fmt.Errorf(
				"'FirstName' can't be NULL: %w", 
				core_errors.ErrInvalidArgument,
			)
		}
		firstNameLen := len([]rune(*r.FirstName.Value))

		if firstNameLen < 3 || firstNameLen > 100 {
			return fmt.Errorf(
				"Invalid `FirstName` length: %d. %w",
				firstNameLen,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.LastName.Set {
		if r.LastName.Value == nil {
			return fmt.Errorf(
				"'LastName' can't be NULL: %w", 
				core_errors.ErrInvalidArgument,
			)
		}
		lastNameLen := len([]rune(*r.LastName.Value))

		if lastNameLen < 3 || lastNameLen > 100 {
			return fmt.Errorf(
				"Invalid `LastName` length: %d. %w",
				lastNameLen,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))

			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf(
					"Invalid `PhoneNumber` length: %d. %w",
					phoneNumberLen,
					core_errors.ErrInvalidArgument,
				)
			}

			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf(
					"Invalid `PhoneNumber` prefix: %s. %w",
					*r.PhoneNumber.Value,
					core_errors.ErrInvalidArgument,
				)
			}
		}
	}

	return nil
}


type PatchUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUser(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	userId, err := core_http_utils.GetPathParam(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get 'id' path param")
		return
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.usersService.PatchUser(ctx, userId, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch user")
		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}


func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.UserPatch{
		FirstName:   request.FirstName.ToDomain(),
		LastName:    request.LastName.ToDomain(),
		PhoneNumber: request.PhoneNumber.ToDomain(),
	}
}
