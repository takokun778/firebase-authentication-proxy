package interactor

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/key"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type KeyFetchPublicInteractor struct {
	output port.KeyFetchPublicOutputPort
}

func NewKeyFetchPublicInteractor(output port.KeyFetchPublicOutputPort) port.KeyFetchPublicInputPort {
	return &KeyFetchPublicInteractor{
		output: output,
	}
}

func (i *KeyFetchPublicInteractor) Execute(ctx context.Context) {
	output := port.KeyFetchPublicOutput{
		PublicKey: key.GetPublic(),
	}

	i.output.Render(ctx, output)
}
