package api

import "fmt"

type Code int
type Error struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%v (code %v)", e.Message, e.Code)
}

const (
	CodeInvalidUrlParams      Code = 200
	CodeInvalidUrlQueryParams Code = 201
	CodeInvalidRequestBody    Code = 202
)

var (
	ErrInvalidUrlParams      = e(CodeInvalidUrlParams, "invalid url params")
	ErrInvalidUrlQueryParams = e(CodeInvalidUrlQueryParams, "invalid url query params")
	ErrInvalidRequestBody    = e(CodeInvalidUrlQueryParams, "invalid request body")
)

func e(Code Code, Message string) Error {
	return Error{Code, Message}
}
