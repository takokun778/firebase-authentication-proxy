package interactor

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseRegisterInteractor struct {
	firebaseRepository port.FirebaseRepository
	userRepository     port.UserRepository
	output             port.FirebaseRegisterOutputPort
}

func NewFirebaseRegisterInteractor(
	firebaseRepository port.FirebaseRepository,
	userRepository port.UserRepository,
	output port.FirebaseRegisterOutputPort,
) port.FirebaseRegisterInputPort {
	return &FirebaseRegisterInteractor{
		firebaseRepository: firebaseRepository,
		userRepository:     userRepository,
		output:             output,
	}
}

func (i *FirebaseRegisterInteractor) Execute(ctx context.Context, input port.FirebaseRegisterInput) {
	user, err := i.userRepository.Save(ctx)
	if err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	if err := i.firebaseRepository.Save(ctx, user.ID(), input.Email, input.Password); err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	i.output.Render(ctx, port.FirebaseRegisterOutput{})
}
