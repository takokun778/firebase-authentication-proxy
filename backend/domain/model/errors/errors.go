package errors

import (
	"fmt"
)

func IsClientError(status int) bool {
	return 400 <= status && status <= 499
}

func IsServerError(status int) bool {
	return 500 <= status && status <= 599
}

type ErrBadRequest struct {
	Err error
}

func NewErrBadRequest(msg string, err error) *ErrBadRequest {
	var e error
	if err != nil {
		e = fmt.Errorf(msg, err)
	} else {
		e = fmt.Errorf(msg)
	}
	return &ErrBadRequest{
		Err: e,
	}
}

func (e ErrBadRequest) Error() string {
	return e.Err.Error()
}

type ErrUnauthorized struct {
	Err error
}

func NewErrUnauthorized(err error) *ErrUnauthorized {
	// if err != nil {
	// 認証エラーとなった原因を解析するために理由をロギングしておく
	// }
	return &ErrUnauthorized{
		Err: fmt.Errorf("unauthorized"),
	}
}

func (e *ErrUnauthorized) Error() string {
	return e.Err.Error()
}
