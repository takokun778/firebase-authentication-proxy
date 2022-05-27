package presenter

import (
	"context"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

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
