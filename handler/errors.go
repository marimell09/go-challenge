package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}

const (
	ErrInvalidAccountDestination = ("Invalid account destination id")
	ErrInvalidBalance            = ("Balance not sufficient for transaction")
)

var (
	ErrUnprocessableEntity = &ErrorResponse{StatusCode: 422, Message: "Unprocessable entity"}
	ErrMethodNotAllowed    = &ErrorResponse{StatusCode: 405, Message: "Method not allowed"}
	ErrNotFound            = &ErrorResponse{StatusCode: 404, Message: "Resource not found"}
	ErrNotAuthorized       = &ErrorResponse{StatusCode: 401, Message: "Not Authorized"}
	ErrBadRequest          = &ErrorResponse{StatusCode: 400, Message: "Bad request"}
)

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func ErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 400,
		StatusText: "Bad request",
		Message:    err.Error(),
	}
}

func ErrorUnprocessable(message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 422,
		StatusText: "Unprocessable entity",
		Message:    message,
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
