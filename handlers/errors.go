package handlers

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrorResponse struct {
	Err        error
	StatusCode int
	StatusText string
	Message    string
}

var (
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: 405, Message: "Method not allowed"}
	ErrNotFound         = &ErrorResponse{StatusCode: 404, Message: "Resource not found"}
	ErrBadRequest       = &ErrorResponse{StatusCode: 400, Message: "Bad request"}
)

func (e *ErrorResponse) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func ErrorRenderer(err error, statusCode int) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: statusCode,
		StatusText: "Bad request",
		Message:    err.Error(),
	}
}

func ServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 500,
		StatusText: "Internal server error",
		Message:    err.Error(),
	}
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	_ = render.Render(w, r, ErrMethodNotAllowed)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	_ = render.Render(w, r, ErrNotFound)
}
