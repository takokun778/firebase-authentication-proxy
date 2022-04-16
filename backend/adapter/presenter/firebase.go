package presenter

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	cx "github.com/takokun778/firebase-authentication-proxy/driver/context"
	"github.com/takokun778/firebase-authentication-proxy/usecase"
)

type FirebasePresenter struct {
	*errorRender
}

func NewFirebasePresenter() usecase.FirebaseOutputPort {
	return &FirebasePresenter{
		errorRender: ErrorRender,
	}
}

func (p *FirebasePresenter) Register(ctx context.Context, output usecase.FirebaseRegisterOutput) {
	r, _ := cx.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebasePresenter) Login(ctx context.Context, output usecase.FirebaseLoginOutput) {
	r, _ := cx.GetResWriter(ctx)

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

	res, _ := json.Marshal(output)

	r.WriteHeader(http.StatusOK)

	_, _ = r.Write(res)
}

func (p *FirebasePresenter) ChangePassword(ctx context.Context, output usecase.FirebaseChangePasswordOutput) {
	r, _ := cx.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebasePresenter) CheckLogin(ctx context.Context, output usecase.FirebaseCheckLoginOutput) {
	r, _ := cx.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebasePresenter) Logout(ctx context.Context, output usecase.FirebaseLogoutOutput) {
	r, _ := cx.GetResWriter(ctx)

	access := &http.Cookie{
		Name:     "access-token",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(r, access)

	refresh := &http.Cookie{
		Name:     "refresh-token",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(r, refresh)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebasePresenter) Withdraw(ctx context.Context, output usecase.FirebaseWithdrawOutput) {
	r, _ := cx.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebasePresenter) Authorize(ctx context.Context, output usecase.FirebaseAuthorizeOutput) {
	r, _ := cx.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p *FirebasePresenter) ErrorRender(ctx context.Context, err error) {
	p.errorRender.ErrRender(ctx, err)
}
