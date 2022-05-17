package presenter

import (
	"context"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseRegisterPresenter struct{}

func NewFirebaseRegisterPresenter() port.FirebaseRegisterOutputPort {
	return &FirebaseRegisterPresenter{}
}

func (p *FirebaseRegisterPresenter) Render(ctx context.Context, output port.FirebaseRegisterOutput) {
	r, _ := adapter.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebaseRegisterPresenter) ErrorRender(ctx context.Context, err error) {
	ErrorRender(ctx, err)
}
