package error

import "fmt"

type Code int

type Error struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%v (code %v)", e.Message, e.Code)
}

func New(Code Code, Message string) Error {
	return Error{Code, Message}
}
