package api

import (
	e "api/internal/error"
)

const (
	CodeInvalidUrlParams      e.Code = 200
	CodeInvalidUrlQueryParams e.Code = 201
	CodeInvalidRequestBody    e.Code = 202
)

var (
	ErrInvalidUrlParams      = e.New(CodeInvalidUrlParams, "invalid url params")
	ErrInvalidUrlQueryParams = e.New(CodeInvalidUrlQueryParams, "invalid url query params")
	ErrInvalidRequestBody    = e.New(CodeInvalidUrlQueryParams, "invalid request body")
)
