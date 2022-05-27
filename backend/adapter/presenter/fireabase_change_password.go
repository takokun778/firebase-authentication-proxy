package presenter

import (
	"context"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseChangePasswordPresenter struct{}

func NewFirebaseChangePasswordPresenter() port.FirebaseChangePasswordOutputPort {
	return &FirebaseChangePasswordPresenter{}
}

func (p *FirebaseChangePasswordPresenter) Render(ctx context.Context, output port.FirebaseChangePasswordOutput) {
	r, _ := adapter.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebaseChangePasswordPresenter) ErrorRender(ctx context.Context, err error) {
	ErrorRender(ctx, err)
}
