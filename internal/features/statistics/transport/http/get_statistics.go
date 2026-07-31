package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
	core_logger "github.com/NeverEverLive/todo-go/internal/core/logger"
	core_http_request "github.com/NeverEverLive/todo-go/internal/core/transport/http/request"
	core_http_response "github.com/NeverEverLive/todo-go/internal/core/transport/http/response"
)

type GetStatisticsResponse struct {
	TotalTasks                 int      `json:"total_tasks"`
	CompletedTasks             int      `json:"completed_tasks"`
	CompletedRateTasks         *float64 `json:"completed_rate_tasks"`
	AverageCompletionTimeTasks *string  `json:"average_completion_time_tasks"`
}

type queryParams struct {
	userID   *string
	dateFrom *time.Time
	dateTo   *time.Time
}

func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	inputParams, err := getUserIDFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID/dateFrom/dateTo query params",
		)
	}

	statistics, err := h.statisticsService.GetStatistics(
		ctx,
		inputParams.userID,
		inputParams.dateFrom,
		inputParams.dateTo,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get statistics",
		)
	}

	response := toDTOFromDomain(statistics)

	responseHandler.JSONResponse(response, http.StatusOK)
}

func toDTOFromDomain(statistics domain.Statistics) GetStatisticsResponse {
	var avgTime *string
	if statistics.AverageCompletionTimeTasks != nil {
		duration := statistics.AverageCompletionTimeTasks.String()
		avgTime = &duration
	}

	return GetStatisticsResponse{
		TotalTasks:                 statistics.TotalTasks,
		CompletedTasks:             statistics.CompletedTasks,
		CompletedRateTasks:         statistics.CompletedRateTasks,
		AverageCompletionTimeTasks: avgTime,
	}
}

func getUserIDFromToQueryParams(r *http.Request) (queryParams, error) {
	const (
		userIDQueryParamKey = "user_id"
		FromQueryParamKey   = "date_from"
		ToQueryParamKey     = "date_to"
	)
	userID := core_http_request.GetQueryParams(r, userIDQueryParamKey)

	dateFrom, err := core_http_request.GetDateQueryParam(r, FromQueryParamKey)
	if err != nil {
		return queryParams{}, fmt.Errorf(
			"error parsing `from` query params: %w",
			err,
		)
	}

	dateTo, err := core_http_request.GetDateQueryParam(r, ToQueryParamKey)
	if err != nil {
		return queryParams{}, fmt.Errorf(
			"error parsing `to` query params: %w",
			err,
		)
	}

	return queryParams{
		userID:   userID,
		dateFrom: dateFrom,
		dateTo:   dateTo,
	}, nil
}
