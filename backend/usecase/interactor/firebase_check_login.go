package interactor

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseCheckLoginInteractor struct {
	firebaseRepository port.FirebaseRepository
	output             port.FirebaseCheckLoginOutputPort
}

func NewFirebaseCheckLoginInteractor(
	firebaseRepository port.FirebaseRepository,
	output port.FirebaseCheckLoginOutputPort,
) port.FirebaseCheckLoginInputPort {
	return &FirebaseCheckLoginInteractor{
		firebaseRepository: firebaseRepository,
		output:             output,
	}
}

func (i *FirebaseCheckLoginInteractor) Execute(ctx context.Context, input port.FirebaseCheckLoginInput) {
	if input.AccessToken.Value() == "" || input.RefreshToken.Value() == "" {
		i.output.ErrorRender(ctx, errors.NewUnauthorizedError(""))

		return
	}

	if err := i.firebaseRepository.Verify(ctx, input.AccessToken); err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	i.output.Render(ctx, port.FirebaseCheckLoginOutput{})
}
