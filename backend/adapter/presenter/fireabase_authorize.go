package presenter

import (
	"context"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseAuthorizePresenter struct{}

func NewFirebaseAuthorizePresenter() port.FirebaseAuthorizeOutputPort {
	return &FirebaseAuthorizePresenter{}
}

func (p *FirebaseAuthorizePresenter) Render(ctx context.Context, output port.FirebaseAuthorizeOutput) {
	r, _ := adapter.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebaseAuthorizePresenter) ErrorRender(ctx context.Context, err error) {
	ErrorRender(ctx, err)
}
