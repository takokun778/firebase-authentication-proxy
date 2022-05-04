package interactor

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/driver/log"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseChangePasswordInteractor struct {
	firebaseRepository port.FirebaseRepository
	output             port.FirebaseChangePasswordOutputPort
}

func NewFirebaseChangePasswordInteractor(
	firebaseRepository port.FirebaseRepository,
	output port.FirebaseChangePasswordOutputPort,
) port.FirebaseChangePasswordInputPort {
	return &FirebaseChangePasswordInteractor{
		firebaseRepository: firebaseRepository,
		output:             output,
	}
}

func (i *FirebaseChangePasswordInteractor) Execute(ctx context.Context, input port.FirebaseChangePasswordInput) {
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

	log.WithCtx(ctx).Sugar().Debugf("old password: %s", input.OldPassword.Value())

	_, err = i.firebaseRepository.Login(ctx, email, input.OldPassword)

	if err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	if input.OldPassword.Equals(input.NewPassword) {
		i.output.ErrorRender(ctx, adapter.NewBadRequestError(""))

		return
	}

	log.WithCtx(ctx).Sugar().Debugf("new password: %s", input.NewPassword.Value())

	err = i.firebaseRepository.ChangePassword(ctx, uid, input.NewPassword)

	if err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	i.output.Render(ctx, port.FirebaseChangePasswordOutput{})
}
