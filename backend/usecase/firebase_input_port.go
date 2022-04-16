package usecase

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/key"
)

type FirebaseInputPort interface {
	Register(context.Context, firebaseRegisterInput)
	Login(context.Context, firebaseLoginInput)
	ChangePassword(context.Context, firebaseChangePasswordInput)
	CheckLogin(context.Context, firebaseCheckLoginInput)
	Logout(context.Context, firebaseLogoutInput)
	Withdraw(context.Context, firebaseWithdrawInput)
	Authorize(context.Context, firebaseAuthorizeInput)
}

type firebaseRegisterInput struct {
	Email    firebase.Email
	Password firebase.Password
}

func NewFirebaseRegisterInput(email string, password []byte) (firebaseRegisterInput, error) {
	e, err := firebase.NewEmail(email)

	if err != nil {
		return firebaseRegisterInput{}, err
	}

	pv, err := key.Decrypt(password)

	if err != nil {
		return firebaseRegisterInput{}, err
	}

	p, err := firebase.NewPassword(string(pv))

	if err != nil {
		return firebaseRegisterInput{}, err
	}

	return firebaseRegisterInput{
		Email:    e,
		Password: p,
	}, nil
}

type firebaseLoginInput struct {
	Email    firebase.Email
	Password firebase.Password
}

func NewFirebaseLoginInput(email string, password []byte) (firebaseLoginInput, error) {
	e, err := firebase.NewEmail(email)

	if err != nil {
		return firebaseLoginInput{}, err
	}

	pv, err := key.Decrypt(password)

	if err != nil {
		return firebaseLoginInput{}, err
	}

	p, err := firebase.NewPassword(string(pv))

	if err != nil {
		return firebaseLoginInput{}, err
	}

	return firebaseLoginInput{
		Email:    e,
		Password: p,
	}, nil
}

type firebaseChangePasswordInput struct {
	AccessToken firebase.AccessToken
	OldPassword firebase.Password
	NewPassword firebase.Password
}

func NewFirebaseChangePasswordInput(token string, oldPassword, newPassword []byte) (firebaseChangePasswordInput, error) {
	at, err := firebase.NewAccessToken(token)

	if err != nil {
		return firebaseChangePasswordInput{}, err
	}

	opv, err := key.Decrypt(oldPassword)

	if err != nil {
		return firebaseChangePasswordInput{}, err
	}

	op, err := firebase.NewPassword(string(opv))

	if err != nil {
		return firebaseChangePasswordInput{}, err
	}

	npv, err := key.Decrypt(newPassword)

	if err != nil {
		return firebaseChangePasswordInput{}, err
	}

	np, err := firebase.NewPassword(string(npv))

	if err != nil {
		return firebaseChangePasswordInput{}, err
	}

	return firebaseChangePasswordInput{
		AccessToken: at,
		OldPassword: op,
		NewPassword: np,
	}, nil
}

type firebaseCheckLoginInput struct {
	AccessToken  firebase.AccessToken
	RefreshToken firebase.RefreshToken
}

func NewFirebaseCheckLoginInput(access, refresh string) (firebaseCheckLoginInput, error) {
	at, err := firebase.NewAccessToken(access)

	if err != nil {
		return firebaseCheckLoginInput{}, err
	}

	ft, err := firebase.NewRefreshToken(refresh)

	if err != nil {
		return firebaseCheckLoginInput{}, err
	}

	return firebaseCheckLoginInput{
		AccessToken:  at,
		RefreshToken: ft,
	}, nil
}

type firebaseLogoutInput struct {
	AccessToken  firebase.AccessToken
	RefreshToken firebase.RefreshToken
}

func NewFirebaseLogoutInput(access, refresh string) (firebaseLogoutInput, error) {
	at, err := firebase.NewAccessToken(access)

	if err != nil {
		return firebaseLogoutInput{}, err
	}

	ft, err := firebase.NewRefreshToken(refresh)

	if err != nil {
		return firebaseLogoutInput{}, err
	}

	return firebaseLogoutInput{
		AccessToken:  at,
		RefreshToken: ft,
	}, nil
}

type firebaseWithdrawInput struct {
	AccessToken firebase.AccessToken
	Password    firebase.Password
}

func NewFirebaseWithdrawInput(token string, password []byte) (firebaseWithdrawInput, error) {
	t, err := firebase.NewAccessToken(token)

	if err != nil {
		return firebaseWithdrawInput{}, err
	}

	pv, err := key.Decrypt(password)

	if err != nil {
		return firebaseWithdrawInput{}, err
	}

	p, err := firebase.NewPassword(string(pv))

	if err != nil {
		return firebaseWithdrawInput{}, err
	}
	return firebaseWithdrawInput{
		AccessToken: t,
		Password:    p,
	}, nil
}

type firebaseAuthorizeInput struct {
	AccessToken firebase.AccessToken
}

func NewFirebaseAuthorizeInput(token string) (firebaseAuthorizeInput, error) {
	t, err := firebase.NewAccessToken(token)

	if err != nil {
		return firebaseAuthorizeInput{}, err
	}

	return firebaseAuthorizeInput{
		AccessToken: t,
	}, nil
}
