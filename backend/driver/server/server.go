package server

import (
	"fmt"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/driver/injector"
)

func Initialize(port string) *http.Server {
	mux := http.NewServeMux()

	router := NewRouter(injector.NewInjector())

	mux.Handle("/", Middleware(router.Handle()))

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
}
