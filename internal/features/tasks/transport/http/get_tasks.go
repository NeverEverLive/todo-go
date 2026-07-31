package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/NeverEverLive/todo-go/internal/core/logger"
	core_http_request "github.com/NeverEverLive/todo-go/internal/core/transport/http/request"
	core_http_response "github.com/NeverEverLive/todo-go/internal/core/transport/http/response"
)

type GetTasksResponse []TaskDTOResponse

// GetTasks godoc
// @Summary Get tasks
// @Description Get tasks filtered by an optional author user ID with optional pagination
// @Tags tasks
// @Accept json
// @Produce json
// @Param user_id query string false "Author user ID" format(uuid)
// @Param limit query int false "Maximum number of tasks to return" minimum(1)
// @Param offset query int false "Number of tasks to skip" minimum(1)
// @Success 200 {object} GetTasksResponse "Successfully retrieved tasks"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /tasks [get]
func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	limit, offset, err := core_http_request.GetLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset query param")
		return
	}
	userID := getUserIDQueryParam(r)
	tasksDomains, err := h.tasksService.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks")
		return
	}

	response := GetTasksResponse(tasksDTOFromDomains(tasksDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIDQueryParam(r *http.Request) *string {
	const (
		userIDQueryParamKey = "user_id"
	)
	return core_http_request.GetQueryParams(r, userIDQueryParamKey)

}
