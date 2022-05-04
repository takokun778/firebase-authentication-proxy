package interactor

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseLogoutInteractor struct {
	firebaseRepository port.FirebaseRepository
	output             port.FirebaseLogoutOutputPort
}

func NewFirebaseLogoutInteractor(
	firebaseRepository port.FirebaseRepository,
	output port.FirebaseLogoutOutputPort,
) port.FirebaseLogoutInputPort {
	return &FirebaseLogoutInteractor{
		firebaseRepository: firebaseRepository,
		output:             output,
	}
}

func (i *FirebaseLogoutInteractor) Execute(ctx context.Context, input port.FirebaseLogoutInput) {
	if err := i.firebaseRepository.Verify(ctx, input.AccessToken); err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	i.output.Render(ctx, port.FirebaseLogoutOutput{})
}
