package usecase

import (
	"context"

	"firebase-authentication/domain/model/key"
)

type KeyInteractor struct {
	keyOutputPort KeyOutputPort
}

func NewKeyInteractor(keyOutputPort KeyOutputPort) KeyInputPort {
	return &KeyInteractor{
		keyOutputPort: keyOutputPort,
	}
}

func (i *KeyInteractor) GetPublic(ctx context.Context, input keyGetPublicInput) {

	output := KeyGetPublicOutput{
		PublicKey: key.PublicKey,
	}

	i.keyOutputPort.GetPublic(ctx, output)
}
