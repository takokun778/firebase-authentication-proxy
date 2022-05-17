package presenter

import (
	"context"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type KeyFetchPresenter struct{}

func NewKeyFetchPresenter() port.KeyFetchOutputPort {
	return &KeyFetchPresenter{}
}

func (p *KeyFetchPresenter) Render(ctx context.Context, output port.KeyFetchOutput) {
	r, _ := adapter.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)

	_, _ = r.Write(output.PublicKey)
}

func (p *KeyFetchPresenter) ErrorRender(ctx context.Context, err error) {
	ErrorRender(ctx, err)
}
