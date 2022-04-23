package server

import (
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/driver/context"
	"github.com/takokun778/firebase-authentication-proxy/driver/injector"
)

type Router struct {
	injector *injector.Injector
}

func NewRouter(injector *injector.Injector) *Router {
	return &Router{
		injector: injector,
	}
}

func (rt *Router) Handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.SetResWriter(r.Context(), w))

		switch r.URL.Path {
		case "/api/register":
			rt.injector.Firebase.Register(w, r)
		case "/api/login":
			rt.injector.Firebase.Login(w, r)
		case "/api/change/password":
			rt.injector.Firebase.ChangePassword(w, r)
		case "/api/login/check":
			rt.injector.Firebase.CheckLogin(w, r)
		case "/api/logout":
			rt.injector.Firebase.Logout(w, r)
		case "/api/withdraw":
			rt.injector.Firebase.Withdraw(w, r)
		case "/api/key":
			rt.injector.Key.GetPublic(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
