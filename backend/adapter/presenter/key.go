package presenter

import (
	"context"
	"net/http"

	cx "firebase-authentication/driver/context"
	"firebase-authentication/usecase"
)

type KeyPresenter struct {
	*errorRender
}

func NewKeyPresenter() usecase.KeyOutputPort {
	return &KeyPresenter{
		errorRender: ErrorRender,
	}
}

func (p *KeyPresenter) GetPublic(ctx context.Context, output usecase.KeyGetPublicOutput) {
	r, _ := cx.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)

	_, _ = r.Write(output.PublicKey)
}

func (p *KeyPresenter) ErrorRender(ctx context.Context, err error) {
	p.errorRender.ErrRender(ctx, err)
}
