package presenter

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
	cx "github.com/takokun778/firebase-authentication-proxy/driver/context"
	"github.com/takokun778/firebase-authentication-proxy/driver/log"
)

var ErrorRender = &errorRender{}

type errorRender struct{}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (errorRender) ErrRender(ctx context.Context, err error) {
	r, _ := cx.GetResWriter(ctx)

	var status int

	switch err.(type) {
	case *errors.BadRequestError:
		status = http.StatusBadRequest
	case *errors.UnauthorizedError:
		status = http.StatusUnauthorized
	default:
		status = http.StatusInternalServerError
	}

	if errors.IsClientError(status) {
		log.WithCtx(ctx).Debug(err.Error())
	}

	if errors.IsServerError(status) {
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
