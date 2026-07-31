package users_transport_http

import (
	"net/http"

	core_logger "github.com/NeverEverLive/todo-go/internal/core/logger"
	core_http_request "github.com/NeverEverLive/todo-go/internal/core/transport/http/request"
	core_http_response "github.com/NeverEverLive/todo-go/internal/core/transport/http/response"
)

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 "Successfully deleted user"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "User not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users/{id} [delete]
func (h *UsersHTTPHandler) DeleteUser(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	userId, err := core_http_request.GetPathParam(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get 'id' query param")
		return
	}

	err = h.usersService.DeleteUser(ctx, userId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to delete user")
		return
	}

	responseHandler.NoContentResponse()
}
