package presenter

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

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
