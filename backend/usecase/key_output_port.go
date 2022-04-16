package usecase

import (
	"context"
)

type KeyOutputPort interface {
	GetPublic(context.Context, KeyGetPublicOutput)
	ErrorRender(context.Context, error)
}

type KeyGetPublicOutput struct {
	PublicKey []byte
}
