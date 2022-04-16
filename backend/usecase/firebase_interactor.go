package usecase

import (
	"context"

	"firebase-authentication/domain/model/errors"
	"firebase-authentication/domain/model/firebase"
	"firebase-authentication/domain/model/user"
	"firebase-authentication/driver/log"
)

type FirebaseInteractor struct {
	firebaseRepository firebase.Repository
	firebaseOutputPort FirebaseOutputPort
	userRepository     user.Repository
}

func NewFirebaseInteractor(
	fRepository firebase.Repository,
	fOutput FirebaseOutputPort,
	uRepository user.Repository,
) FirebaseInputPort {
	return &FirebaseInteractor{
		firebaseRepository: fRepository,
		firebaseOutputPort: fOutput,
		userRepository:     uRepository,
	}
}

func (i *FirebaseInteractor) Register(ctx context.Context, input firebaseRegisterInput) {
	user, _ := i.userRepository.Save(ctx)

	log.WithCtx(ctx).Sugar().Debugf("received password is %s", input.Password.Value())

	err := i.firebaseRepository.Save(ctx, user.Id(), input.Email, input.Password)

	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)
		return
	}

	log.WithCtx(ctx).Sugar().Infof("uid %s register success", user.Id().Value())

	i.firebaseOutputPort.Register(ctx, FirebaseRegisterOutput{})
}

func (i *FirebaseInteractor) Login(ctx context.Context, input firebaseLoginInput) {
	tokens, err := i.firebaseRepository.GenerateTokens(ctx, input.Email, input.Password)

	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)
		return
	}

	log.WithCtx(ctx).Sugar().Infof("uid %s login success", tokens.Uid.Value())

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

	uid, err := input.AccessToken.GetUid()

	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)
		return
	}

	log.WithCtx(ctx).Sugar().Debugf("old password: %s", input.OldPassword.Value())

	_, err = i.firebaseRepository.GenerateTokens(ctx, email, input.OldPassword)

	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)
		return
	}

	if input.OldPassword.Equals(input.NewPassword) {
		i.firebaseOutputPort.ErrorRender(ctx, errors.NewErrBadRequest("", nil))
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
		i.firebaseOutputPort.ErrorRender(ctx, errors.NewErrUnauthorized(nil))
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

	uid, err := input.AccessToken.GetUid()

	if err != nil {
		i.firebaseOutputPort.ErrorRender(ctx, err)
		return
	}

	log.WithCtx(ctx).Sugar().Debugf("old password: %s", input.Password.Value())

	_, err = i.firebaseRepository.GenerateTokens(ctx, email, input.Password)

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
