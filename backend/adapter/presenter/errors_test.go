package presenter_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/adapter/mock"
	"github.com/takokun778/firebase-authentication-proxy/adapter/presenter"
	dme "github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
)

func TestErrorRender(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		err error
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "400",
			args: args{
				ctx: adapter.SetResWriter(context.Background(), &mock.ResponseWriter{
					Assert: func(statusCode int) {
						if statusCode != http.StatusBadRequest {
							t.Errorf("http status code: %d, expected %d", statusCode, http.StatusBadRequest)
						}
					},
				}),
				err: adapter.NewBadRequestError(""),
			},
		},
		{
			name: "403",
			args: args{
				ctx: adapter.SetResWriter(context.Background(), &mock.ResponseWriter{
					Assert: func(statusCode int) {
						if statusCode != http.StatusUnauthorized {
							t.Errorf("http status code: %d, expected %d", statusCode, http.StatusUnauthorized)
						}
					},
				}),
				err: dme.NewUnauthorizedError(""),
			},
		},
		{
			name: "405",
			args: args{
				ctx: adapter.SetResWriter(context.Background(), &mock.ResponseWriter{
					Assert: func(statusCode int) {
						if statusCode != http.StatusMethodNotAllowed {
							t.Errorf("http status code: %d, expected %d", statusCode, http.StatusMethodNotAllowed)
						}
					},
				}),
				err: adapter.NewMethodNotAllowedError(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			presenter.ErrorRender(tt.args.ctx, tt.args.err)
		})
	}
}
