package interactor

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseLoginInteractor struct {
	firebaseRepository port.FirebaseRepository
	output             port.FirebaseLoginOutputPort
}

func NewFirebaseLoginInteractor(
	firebaseRepository port.FirebaseRepository,
	output port.FirebaseLoginOutputPort,
) port.FirebaseLoginInputPort {
	return &FirebaseLoginInteractor{
		firebaseRepository: firebaseRepository,
		output:             output,
	}
}

func (i *FirebaseLoginInteractor) Execute(ctx context.Context, input port.FirebaseLoginInput) {
	tokens, err := i.firebaseRepository.Login(ctx, input.Email, input.Password)
	if err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	expires, err := tokens.AccessToken.CalcExpires()
	if err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	i.output.Render(ctx, port.FirebaseLoginOutput{
		AccessToken:  tokens.AccessToken.Value(),
		RefreshToken: tokens.RefreshToken.Value(),
		Expires:      expires,
	})
}
