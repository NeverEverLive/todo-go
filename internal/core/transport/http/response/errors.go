package core_http_response

type ErrorResponse struct {
	Error   string `json:"error" example:"Full error message"`
	Message string `json:"message" example:"Short human readable message"`
}
