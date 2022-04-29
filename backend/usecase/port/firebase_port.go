package port

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"
)

type FirebaseRegisterInputPort interface {
	Execute(context.Context, FirebaseRegisterInput)
}

type FirebaseRegisterInput struct {
	Email    firebase.Email
	Password firebase.Password
}

func NewFirebaseRegisterInput(email string, password string) (FirebaseRegisterInput, error) {
	e, err := firebase.NewEmail(email)
	if err != nil {
		return FirebaseRegisterInput{}, err
	}

	p, err := firebase.NewPassword(password)
	if err != nil {
		return FirebaseRegisterInput{}, err
	}

	return FirebaseRegisterInput{
		Email:    e,
		Password: p,
	}, nil
}

type FirebaseRegisterOutputPort interface {
	Render(context.Context, FirebaseRegisterOutput)
	ErrorRender(context.Context, error)
}

type FirebaseRegisterOutput struct{}

type FirebaseLoginInputPort interface {
	Execute(context.Context, FirebaseLoginInput)
}

type FirebaseLoginInput struct {
	Email    firebase.Email
	Password firebase.Password
}

func NewFirebaseLoginInput(email, password string) (FirebaseLoginInput, error) {
	e, err := firebase.NewEmail(email)
	if err != nil {
		return FirebaseLoginInput{}, err
	}

	p, err := firebase.NewPassword(password)
	if err != nil {
		return FirebaseLoginInput{}, err
	}

	return FirebaseLoginInput{
		Email:    e,
		Password: p,
	}, nil
}

type FirebaseLoginOutputPort interface {
	Render(context.Context, FirebaseLoginOutput)
	ErrorRender(context.Context, error)
}

type FirebaseLoginOutput struct {
	AccessToken  string
	RefreshToken string
	Expires      int
}

type FirebaseChangePasswordInputPort interface {
	Execute(context.Context, FirebaseChangePasswordInput)
}

type FirebaseChangePasswordInput struct {
	AccessToken firebase.AccessToken
	OldPassword firebase.Password
	NewPassword firebase.Password
}

func NewFirebaseChangePasswordInput(
	token string,
	oldPassword string,
	newPassword string,
) (FirebaseChangePasswordInput, error) {
	at, err := firebase.NewAccessToken(token)
	if err != nil {
		return FirebaseChangePasswordInput{}, err
	}

	op, err := firebase.NewPassword(oldPassword)
	if err != nil {
		return FirebaseChangePasswordInput{}, err
	}

	np, err := firebase.NewPassword(newPassword)
	if err != nil {
		return FirebaseChangePasswordInput{}, err
	}

	return FirebaseChangePasswordInput{
		AccessToken: at,
		OldPassword: op,
		NewPassword: np,
	}, nil
}

type FirebaseChangePasswordOutputPort interface {
	Render(context.Context, FirebaseChangePasswordOutput)
	ErrorRender(context.Context, error)
}

type FirebaseChangePasswordOutput struct{}

type FirebaseCheckLoginInputPort interface {
	Execute(context.Context, FirebaseCheckLoginInput)
}

type FirebaseCheckLoginInput struct {
	AccessToken  firebase.AccessToken
	RefreshToken firebase.RefreshToken
}

func NewFirebaseCheckLoginInput(access, refresh string) (FirebaseCheckLoginInput, error) {
	at, err := firebase.NewAccessToken(access)
	if err != nil {
		return FirebaseCheckLoginInput{}, err
	}

	ft, err := firebase.NewRefreshToken(refresh)
	if err != nil {
		return FirebaseCheckLoginInput{}, err
	}

	return FirebaseCheckLoginInput{
		AccessToken:  at,
		RefreshToken: ft,
	}, nil
}

type FirebaseCheckLoginOutputPort interface {
	Render(context.Context, FirebaseCheckLoginOutput)
	ErrorRender(context.Context, error)
}

type FirebaseCheckLoginOutput struct{}

type FirebaseLogoutInputPort interface {
	Execute(context.Context, FirebaseLogoutInput)
}

type FirebaseLogoutInput struct {
	AccessToken  firebase.AccessToken
	RefreshToken firebase.RefreshToken
}

func NewFirebaseLogoutInput(access, refresh string) (FirebaseLogoutInput, error) {
	at, err := firebase.NewAccessToken(access)
	if err != nil {
		return FirebaseLogoutInput{}, err
	}

	ft, err := firebase.NewRefreshToken(refresh)
	if err != nil {
		return FirebaseLogoutInput{}, err
	}

	return FirebaseLogoutInput{
		AccessToken:  at,
		RefreshToken: ft,
	}, nil
}

type FirebaseLogoutOutputPort interface {
	Render(context.Context, FirebaseLogoutOutput)
	ErrorRender(context.Context, error)
}

type FirebaseLogoutOutput struct{}

type FirebaseWithdrawInputPort interface {
	Execute(context.Context, FirebaseWithdrawInput)
}

type FirebaseWithdrawInput struct {
	AccessToken firebase.AccessToken
	Password    firebase.Password
}

func NewFirebaseWithdrawInput(token, password string) (FirebaseWithdrawInput, error) {
	t, err := firebase.NewAccessToken(token)
	if err != nil {
		return FirebaseWithdrawInput{}, err
	}

	p, err := firebase.NewPassword(password)
	if err != nil {
		return FirebaseWithdrawInput{}, err
	}

	return FirebaseWithdrawInput{
		AccessToken: t,
		Password:    p,
	}, nil
}

type FirebaseWithdrawOutputPort interface {
	Render(context.Context, FirebaseWithdrawOutput)
	ErrorRender(context.Context, error)
}

type FirebaseWithdrawOutput struct{}

type FirebaseAuthorizeInputPort interface {
	Execute(context.Context, FirebaseAuthorizeInput)
}

type FirebaseAuthorizeInput struct {
	AccessToken firebase.AccessToken
}

func NewFirebaseAuthorizeInput(token string) (FirebaseAuthorizeInput, error) {
	t, err := firebase.NewAccessToken(token)
	if err != nil {
		return FirebaseAuthorizeInput{}, err
	}

	return FirebaseAuthorizeInput{
		AccessToken: t,
	}, nil
}

type FirebaseAuthorizeOutputPort interface {
	Render(context.Context, FirebaseAuthorizeOutput)
	ErrorRender(context.Context, error)
}

type FirebaseAuthorizeOutput struct{}

type Tokens struct {
	AccessToken  firebase.AccessToken
	RefreshToken firebase.RefreshToken
}

type FirebaseRepository interface {
	Save(context.Context, user.ID, firebase.Email, firebase.Password) error
	Login(context.Context, firebase.Email, firebase.Password) (Tokens, error)
	ChangePassword(context.Context, firebase.UID, firebase.Password) error
	Verify(context.Context, firebase.AccessToken) error
	Delete(context.Context, firebase.UID) error
}
