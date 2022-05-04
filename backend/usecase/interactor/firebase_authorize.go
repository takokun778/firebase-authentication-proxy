package interactor

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseAuthorizeInteractor struct {
	firebaseRepository port.FirebaseRepository
	output             port.FirebaseAuthorizeOutputPort
}

func NewFirebaseAuthorizeInteractor(
	firebaseRepository port.FirebaseRepository,
	output port.FirebaseAuthorizeOutputPort,
) port.FirebaseAuthorizeInputPort {
	return &FirebaseAuthorizeInteractor{
		firebaseRepository: firebaseRepository,
		output:             output,
	}
}

func (i *FirebaseAuthorizeInteractor) Execute(ctx context.Context, input port.FirebaseAuthorizeInput) {
	if err := i.firebaseRepository.Verify(ctx, input.AccessToken); err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	i.output.Render(ctx, port.FirebaseAuthorizeOutput{})
}
