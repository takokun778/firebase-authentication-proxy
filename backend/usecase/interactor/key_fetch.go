package interactor

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/key"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type KeyFetchInteractor struct {
	output port.KeyFetchOutputPort
}

func NewKeyFetchInteractor(output port.KeyFetchOutputPort) port.KeyFetchInputPort {
	return &KeyFetchInteractor{
		output: output,
	}
}

func (i *KeyFetchInteractor) Execute(ctx context.Context) {
	output := port.KeyFetchOutput{
		PublicKey: key.GetPublic(),
	}

	i.output.Render(ctx, output)
}
