package errors

import (
	"errors"
)

func IsClientError(status int) bool {
	return 400 <= status && status <= 499
}

func IsServerError(status int) bool {
	return 500 <= status && status <= 599
}

type BadRequestError struct {
	Msg string
}

func NewBadRequestError(msg string) *BadRequestError {
	return &BadRequestError{
		Msg: msg,
	}
}

func (e BadRequestError) Error() string {
	return e.Msg
}

var ErrUnauthorized = errors.New("unauthorized")

type UnauthorizedError struct {
	Err error
}

func NewUnauthorizedError(err error) *UnauthorizedError {
	// if err != nil {
	// 認証エラーとなった原因を解析するために理由をロギングしておく
	// }
	return &UnauthorizedError{
		Err: ErrUnauthorized,
	}
}

func (e *UnauthorizedError) Error() string {
	return e.Err.Error()
}
