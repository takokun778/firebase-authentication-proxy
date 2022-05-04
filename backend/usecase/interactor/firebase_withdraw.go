package interactor

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/driver/log"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseWithdrawInteractor struct {
	firebaseRepository port.FirebaseRepository
	output             port.FirebaseWithdrawOutputPort
}

func NewFirebaseWithdrawInteractor(
	firebaseRepository port.FirebaseRepository,
	output port.FirebaseWithdrawOutputPort,
) port.FirebaseWithdrawInputPort {
	return &FirebaseWithdrawInteractor{
		firebaseRepository: firebaseRepository,
		output:             output,
	}
}

func (i *FirebaseWithdrawInteractor) Execute(ctx context.Context, input port.FirebaseWithdrawInput) {
	if err := i.firebaseRepository.Verify(ctx, input.AccessToken); err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	email, err := input.AccessToken.GetEmail()
	if err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	uid, err := input.AccessToken.GetUID()
	if err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	log.WithCtx(ctx).Sugar().Debugf("password: %s", input.Password.Value())

	_, err = i.firebaseRepository.Login(ctx, email, input.Password)

	if err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	if err := i.firebaseRepository.Delete(ctx, uid); err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	i.output.Render(ctx, port.FirebaseWithdrawOutput{})
}
