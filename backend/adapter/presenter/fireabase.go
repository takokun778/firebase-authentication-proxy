package presenter

import (
	"context"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseCheckLoginPresenter struct{}

func NewFirebaseCheckLoginPresenter() port.FirebaseCheckLoginOutputPort {
	return &FirebaseCheckLoginPresenter{}
}

func (p *FirebaseCheckLoginPresenter) Render(ctx context.Context, output port.FirebaseCheckLoginOutput) {
	r, _ := adapter.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebaseCheckLoginPresenter) ErrorRender(ctx context.Context, err error) {
	ErrorRender(ctx, err)
}

type FirebaseLogoutPresenter struct{}

func NewFirebaseLogoutPresenter() port.FirebaseLogoutOutputPort {
	return &FirebaseLogoutPresenter{}
}

func (p *FirebaseLogoutPresenter) Render(ctx context.Context, output port.FirebaseLogoutOutput) {
	r, _ := adapter.GetResWriter(ctx)

	access := &http.Cookie{
		Name:     "access-token",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(r, access)

	refresh := &http.Cookie{
		Name:     "refresh-token",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(r, refresh)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebaseLogoutPresenter) ErrorRender(ctx context.Context, err error) {
	ErrorRender(ctx, err)
}

type FirebaseWithdrawPresenter struct{}

func NewFirebaseWithdrawPresenter() port.FirebaseWithdrawOutputPort {
	return &FirebaseWithdrawPresenter{}
}

func (p *FirebaseWithdrawPresenter) Render(ctx context.Context, output port.FirebaseWithdrawOutput) {
	r, _ := adapter.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebaseWithdrawPresenter) ErrorRender(ctx context.Context, err error) {
	ErrorRender(ctx, err)
}

type FirebaseAuthorizePresenter struct{}

func NewFirebaseAuthorizePresenter() port.FirebaseAuthorizeOutputPort {
	return &FirebaseAuthorizePresenter{}
}

func (p *FirebaseAuthorizePresenter) Render(ctx context.Context, output port.FirebaseAuthorizeOutput) {
	r, _ := adapter.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebaseAuthorizePresenter) ErrorRender(ctx context.Context, err error) {
	ErrorRender(ctx, err)
}
