package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
)

func GetQueryParams(r *http.Request, key string) *string {
	param := r.URL.Query().Get(key)

	if param == "" {
		return nil
	}
	return &param
}

func GetIntQueryParams(r *http.Request, key string) (*int, error) {
	param := GetQueryParams(r, key)
	if param == nil { return nil, nil}

	intParam, err := strconv.Atoi(*param)
	if err != nil {
		return nil, fmt.Errorf(
			"param='%s' by key='%s' not a valid integer: %v. %w",
			param,
			key,
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return &intParam, nil
}

func GetLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)
	limit, err := GetIntQueryParams(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"get 'limit' query param: %w",
			err,
		)
	}

	offset, err := GetIntQueryParams(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"get 'offset' query param: %w",
			err,
		)
	}

	return limit, offset, nil
}
