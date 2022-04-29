package port

import (
	"context"
)

type KeyFetchPublicInputPort interface {
	Execute(context.Context)
}

type KeyFetchPublicOutputPort interface {
	Render(context.Context, KeyFetchPublicOutput)
	ErrorRender(context.Context, error)
}

type KeyFetchPublicOutput struct {
	PublicKey []byte
}
