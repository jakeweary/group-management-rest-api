package database

import (
	e "api/internal/error"
)

const (
	CodeParentGroupDoesntExist e.Code = 100
	CodeGroupDoesntExist       e.Code = 101
	CodeUserDoesntExist        e.Code = 102
)

var (
	ErrParentGroupDoesntExist = e.New(CodeParentGroupDoesntExist, "parent group doesn't exist")
	ErrGroupDoesntExist       = e.New(CodeGroupDoesntExist, "group doesn't exist")
	ErrUserDoesntExist        = e.New(CodeUserDoesntExist, "user doesn't exist")
)
