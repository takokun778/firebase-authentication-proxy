package usecase

import (
	"context"
)

type KeyInputPort interface {
	GetPublic(context.Context, keyGetPublicInput)
}

type keyGetPublicInput struct{}

func NewKeyGetPublicInput() (keyGetPublicInput, error) {
	return keyGetPublicInput{}, nil
}
