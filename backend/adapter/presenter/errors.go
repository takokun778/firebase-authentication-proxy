package presenter

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	dme "github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
	"github.com/takokun778/firebase-authentication-proxy/driver/log"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ErrorRender(ctx context.Context, err error) {
	r, _ := adapter.GetResWriter(ctx)

	var status int

	switch {
	case AsBadRequest(err):
		status = http.StatusBadRequest
	case AsUnauthorized(err):
		status = http.StatusUnauthorized
	case AsMethodNotAllowed(err):
		status = http.StatusMethodNotAllowed
	default:
		status = http.StatusInternalServerError
	}

	switch {
	case IsClientError(status):
		log.WithCtx(ctx).Warn(err.Error())
	case IsServerError(status):
		log.WithCtx(ctx).Error(err.Error())
	}

	er := ErrorResponse{
		Status:  status,
		Message: err.Error(),
	}

	res, err := json.Marshal(er)
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)

		return
	}

	r.WriteHeader(status)

	_, _ = r.Write(res)
}

func AsBadRequest(err error) bool {
	var (
		b *adapter.BadRequestError
		v *dme.ValidateError
	)

	return errors.As(err, &b) || errors.As(err, &v)
}

func AsUnauthorized(err error) bool {
	var target *dme.UnauthorizedError

	return errors.As(err, &target)
}

func AsMethodNotAllowed(err error) bool {
	var target *adapter.MethodNotAllowedError

	return errors.As(err, &target)
}

func IsClientError(status int) bool {
	return 400 <= status && status <= 499
}

func IsServerError(status int) bool {
	return 500 <= status && status <= 599
}
