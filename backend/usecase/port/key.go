package port

import (
	"context"
)

type KeyFetchInputPort interface {
	Execute(context.Context)
}

type KeyFetchOutputPort interface {
	Render(context.Context, KeyFetchOutput)
	ErrorRender(context.Context, error)
}

type KeyFetchOutput struct {
	PublicKey []byte
}
