package server

import (
	"net/http"
	"os"
	"time"

	"github.com/takokun778/firebase-authentication-proxy/driver/context"
	"github.com/takokun778/firebase-authentication-proxy/driver/log"
)

var corsAllowOrigin string

func init() {
	corsAllowOrigin = os.Getenv("CORS_ALLOW_ORIGIN")
	if corsAllowOrigin == "" {
		log.Log().Fatal("cors allow origin is empty")
	}
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers",
			"Origin, Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Origin", corsAllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		now := time.Now()
		r = r.WithContext(context.SetReq(r.Context()))
		defer log.Access(r.Context(), r.URL.Path, r.Method, now)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

			return
		}

		next.ServeHTTP(w, r)
	})
}
