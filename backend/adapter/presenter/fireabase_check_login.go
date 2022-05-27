package presenter

import (
	"context"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseCheckLoginPresenter struct{}

func NewFirebaseCheckLoginPresenter() port.FirebaseCheckLoginOutputPort {
	return &FirebaseCheckLoginPresenter{}
}

func (p *FirebaseCheckLoginPresenter) Render(ctx context.Context, output port.FirebaseCheckLoginOutput) {
	r, _ := adapter.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebaseCheckLoginPresenter) ErrorRender(ctx context.Context, err error) {
	ErrorRender(ctx, err)
}
