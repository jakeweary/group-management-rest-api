package database

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
	CodeParentGroupDoesntExist Code = 100
	CodeGroupDoesntExist       Code = 101
	CodeUserDoesntExist        Code = 102
)

var (
	ErrParentGroupDoesntExist = e(CodeParentGroupDoesntExist, "parent group doesn't exist")
	ErrGroupDoesntExist       = e(CodeGroupDoesntExist, "group doesn't exist")
	ErrUserDoesntExist        = e(CodeUserDoesntExist, "user doesn't exist")
)

func e(Code Code, Message string) Error {
	return Error{Code, Message}
}
