package presenter

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseRegisterPresenter struct{}

func NewFirebaseRegisterPresenter() port.FirebaseRegisterOutputPort {
	return &FirebaseRegisterPresenter{}
}

func (p *FirebaseRegisterPresenter) Render(ctx context.Context, output port.FirebaseRegisterOutput) {
	r, _ := adapter.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebaseRegisterPresenter) ErrorRender(ctx context.Context, err error) {
	ErrorRender(ctx, err)
}

type FirebaseLoginPresenter struct{}

func NewFirebaseLoginPresenter() port.FirebaseLoginOutputPort {
	return &FirebaseLoginPresenter{}
}

func (p *FirebaseLoginPresenter) Render(ctx context.Context, output port.FirebaseLoginOutput) {
	r, _ := adapter.GetResWriter(ctx)

	access := &http.Cookie{
		Name:     "access-token",
		Value:    output.AccessToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(output.Expires) * time.Second),
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(r, access)

	refresh := &http.Cookie{
		Name:     "refresh-token",
		Value:    output.RefreshToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(output.Expires) * time.Second),
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(r, refresh)

	res, err := json.Marshal(output)
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)

		return
	}

	r.WriteHeader(http.StatusOK)

	_, _ = r.Write(res)
}

func (p *FirebaseLoginPresenter) ErrorRender(ctx context.Context, err error) {
	ErrorRender(ctx, err)
}

type FirebaseChangePasswordPresenter struct{}

func NewFirebaseChangePasswordPresenter() port.FirebaseChangePasswordOutputPort {
	return &FirebaseChangePasswordPresenter{}
}

func (p *FirebaseChangePasswordPresenter) Render(ctx context.Context, output port.FirebaseChangePasswordOutput) {
	r, _ := adapter.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebaseChangePasswordPresenter) ErrorRender(ctx context.Context, err error) {
}

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
