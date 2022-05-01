package interactor

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
	"github.com/takokun778/firebase-authentication-proxy/driver/log"
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
	user, _ := i.userRepository.Save(ctx)

	err := i.firebaseRepository.Save(ctx, user.ID(), input.Email, input.Password)
	if err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	i.output.Render(ctx, port.FirebaseRegisterOutput{})
}

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

	if err := i.firebaseRepository.Verify(ctx, input.AccessToken); err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	if err := i.firebaseRepository.Delete(ctx, uid); err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	i.output.Render(ctx, port.FirebaseWithdrawOutput{})
}

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
