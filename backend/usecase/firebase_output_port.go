package usecase

import (
	"context"
)

type FirebaseOutputPort interface {
	Register(context.Context, FirebaseRegisterOutput)
	Login(context.Context, FirebaseLoginOutput)
	ChangePassword(context.Context, FirebaseChangePasswordOutput)
	Logout(context.Context, FirebaseLogoutOutput)
	CheckLogin(context.Context, FirebaseCheckLoginOutput)
	Withdraw(context.Context, FirebaseWithdrawOutput)
	Authorize(context.Context, FirebaseAuthorizeOutput)
	ErrorRender(context.Context, error)
}

type FirebaseRegisterOutput struct{}

type FirebaseLoginOutput struct {
	AccessToken  string
	RefreshToken string
	Expires      int
}

type FirebaseChangePasswordOutput struct{}

type FirebaseCheckLoginOutput struct{}

type FirebaseLogoutOutput struct{}

type FirebaseWithdrawOutput struct{}

type FirebaseAuthorizeOutput struct{}
