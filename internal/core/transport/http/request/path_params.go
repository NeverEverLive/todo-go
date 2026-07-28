package core_http_request

import (
	"fmt"
	"net/http"

	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
)

func GetPathParam(r *http.Request, key string) (string, error) {
	param := r.PathValue(key)
	if param == "" {
		return "", fmt.Errorf(
			"get '%s' path param: %w",
			key,
			core_errors.ErrInvalidArgument,
		)
	}

	return param, nil
}
