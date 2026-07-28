package core_http_middleware

import "slices"

import "net/http"

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(
	h http.Handler,
	m ...Middleware,
) http.Handler {
	if len(m) == 0 {
		return h
	}

	for _, middleware := range slices.Backward(m) {
		h = middleware(h)
	}

	return h
}
