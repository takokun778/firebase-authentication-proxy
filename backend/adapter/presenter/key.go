package presenter

import (
	"context"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type KeyFetchPublicPresenter struct{}

func NewKeyGetPublicPresenter() port.KeyFetchPublicOutputPort {
	return &KeyFetchPublicPresenter{}
}

func (p *KeyFetchPublicPresenter) Render(ctx context.Context, output port.KeyFetchPublicOutput) {
	r, _ := adapter.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)

	_, _ = r.Write(output.PublicKey)
}

func (p *KeyFetchPublicPresenter) ErrorRender(ctx context.Context, err error) {
	ErrorRender(ctx, err)
}
