package users_transport_http

import (
	"net/http"

	core_logger "github.com/NeverEverLive/todo-go/internal/core/logger"
	core_http_response "github.com/NeverEverLive/todo-go/internal/core/transport/http/response"
	core_http_utils "github.com/NeverEverLive/todo-go/internal/core/transport/http/utils"
)


func (h *UsersHTTPHandler) DeleteUser(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	userId, err := core_http_utils.GetPathParam(r, "id")
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
