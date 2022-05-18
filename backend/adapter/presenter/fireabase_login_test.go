package presenter_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/adapter/mock"
	"github.com/takokun778/firebase-authentication-proxy/adapter/presenter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

func TestFirebaseLoginPresenterRender(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx    context.Context
		output port.FirebaseLoginOutput
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "正常に動作することを確認",
			args: args{
				ctx: adapter.SetResWriter(context.Background(), &mock.ResponseWriter{
					Assert: func(statusCode int) {
						if statusCode != http.StatusOK {
							t.Errorf("http status code: %d, expected %d", statusCode, http.StatusOK)
						}
					},
				}),
				output: port.FirebaseLoginOutput{
					AccessToken:  "",
					RefreshToken: "",
					Expires:      3600,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := presenter.NewFirebaseLoginPresenter()
			p.Render(tt.args.ctx, tt.args.output)
		})
	}
}
