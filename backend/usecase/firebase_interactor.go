package usecase

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"
	"github.com/takokun778/firebase-authentication-proxy/driver/log"
)

type FirebaseInteractor struct {
	firebaseRepository firebase.Repository
	firebaseOutputPort FirebaseOutputPort
	userRepository     user.Repository
}

func NewFirebaseInteractor(
	firebaseRepository firebase.Repository,
	firebaseOutput FirebaseOutputPort,
	userRepository user.Repository,
) FirebaseInputPort {
	return &FirebaseInteractor{
		firebaseRepository: firebaseRepository,
		firebaseOutputPort: firebaseOutput,
		userRepository:     userRepository,
	}
}

func (i *FirebaseInteractor) Register(ctx context.Context, input firebaseRegisterInput) {
	user, _ := i.userRepository.Save(ctx)

	log.WithCtx(ctx).Sugar().Debugf("received password is %s", input.Password.Value())

	err := i.firebaseRepository.Save(ctx, user.ID(), input.Email, input.Password)
	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}

	log.WithCtx(ctx).Sugar().Infof("uid %s register success", user.ID().Value())

	i.firebaseOutputPort.Register(ctx, FirebaseRegisterOutput{})
}

func (i *FirebaseInteractor) Login(ctx context.Context, input firebaseLoginInput) {
	tokens, err := i.firebaseRepository.Login(ctx, input.Email, input.Password)
	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}

	log.WithCtx(ctx).Sugar().Infof("uid %s login success", tokens.UID.Value())

	i.firebaseOutputPort.Login(ctx, FirebaseLoginOutput{
		AccessToken:  tokens.AccessToken.Value(),
		RefreshToken: tokens.RefreshToken.Value(),
		Expires:      tokens.Expires,
	})
}

func (i *FirebaseInteractor) ChangePassword(ctx context.Context, input firebaseChangePasswordInput) {
	email, err := input.AccessToken.GetEmail()
	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}

	uid, err := input.AccessToken.GetUID()
	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}

	log.WithCtx(ctx).Sugar().Debugf("old password: %s", input.OldPassword.Value())

	_, err = i.firebaseRepository.Login(ctx, email, input.OldPassword)

	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}

	if input.OldPassword.Equals(input.NewPassword) {
		i.firebaseOutputPort.ErrorRender(ctx, errors.NewBadRequestError(""))

		return
	}

	log.WithCtx(ctx).Sugar().Debugf("new password: %s", input.NewPassword.Value())

	err = i.firebaseRepository.ChangePassword(ctx, uid, input.NewPassword)

	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}
}

func (i *FirebaseInteractor) CheckLogin(ctx context.Context, input firebaseCheckLoginInput) {
	if input.AccessToken.Value() == "" || input.RefreshToken.Value() == "" {
		i.firebaseOutputPort.ErrorRender(ctx, errors.NewUnauthorizedError(nil))

		return
	}

	if err := i.firebaseRepository.Verify(ctx, input.AccessToken); err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}

	i.firebaseOutputPort.CheckLogin(ctx, FirebaseCheckLoginOutput{})
}

func (i *FirebaseInteractor) Logout(ctx context.Context, input firebaseLogoutInput) {
	if err := i.firebaseRepository.Verify(ctx, input.AccessToken); err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}

	i.firebaseOutputPort.Logout(ctx, FirebaseLogoutOutput{})
}

func (i *FirebaseInteractor) Withdraw(ctx context.Context, input firebaseWithdrawInput) {
	email, err := input.AccessToken.GetEmail()
	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}

	uid, err := input.AccessToken.GetUID()
	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}

	log.WithCtx(ctx).Sugar().Debugf("old password: %s", input.Password.Value())

	_, err = i.firebaseRepository.Login(ctx, email, input.Password)

	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}

	if err := i.firebaseRepository.Verify(ctx, input.AccessToken); err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}

	if err := i.firebaseRepository.Delete(ctx, uid); err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}

	i.firebaseOutputPort.Withdraw(ctx, FirebaseWithdrawOutput{})
}

func (i *FirebaseInteractor) Authorize(ctx context.Context, input firebaseAuthorizeInput) {
	if err := i.firebaseRepository.Verify(ctx, input.AccessToken); err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)

		return
	}

	i.firebaseOutputPort.Authorize(ctx, FirebaseAuthorizeOutput{})
}
