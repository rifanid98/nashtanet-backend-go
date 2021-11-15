package response

import "errors"

var (
	ErrInvalidInput        = errors.New("invalid input")
	ErrInternalServerError = errors.New("internal server error")
)

type Error struct {
	StatusCode int      `json:"statusCode"`
	Errors     []string `json:"errors"`
}

func NewError(err error, status int) *Error {
	return &Error{
		StatusCode: status,
		Errors:     []string{err.Error()},
	}
}

func NewErrorMessage(messages []string, status int) *Error {
	return &Error{
		StatusCode: status,
		Errors:     messages,
	}
}
