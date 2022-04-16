package server

import (
	"net/http"

	"firebase-authentication/driver/context"
	"firebase-authentication/driver/injector"
)

type Router struct {
	injector *injector.Injector
}

func NewRouter(injector *injector.Injector) *Router {
	return &Router{
		injector: injector,
	}
}

func (router *Router) Handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.SetResWriter(r.Context(), w))

		path := r.URL.Path

		method := r.Method

		if method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if path == "/api/register" && method == "POST" {
			router.injector.Firebase.Register(w, r)
			return
		}

		if path == "/api/login" && method == "POST" {
			router.injector.Firebase.Login(w, r)
			return
		}

		if path == "/api/change/password" && method == "PUT" {
			router.injector.Firebase.ChangePassword(w, r)
			return
		}

		if path == "/api/login/check" && method == "POST" {
			router.injector.Firebase.CheckLogin(w, r)
			return
		}

		if path == "/api/logout" && method == "POST" {
			router.injector.Firebase.Logout(w, r)
			return
		}

		if path == "/api/withdraw" && method == "POST" {
			router.injector.Firebase.Withdraw(w, r)
			return
		}

		if path == "/api/key" && method == "GET" {
			router.injector.Key.GetPublic(w, r)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	})
}
