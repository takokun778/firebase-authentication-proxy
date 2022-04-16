package server

import (
	"fmt"
	"net/http"

	"firebase-authentication/driver/injector"
)

type Server struct{}

func Serve(port string) error {
	mux := http.NewServeMux()

	router := NewRouter(injector.NewInjector())

	mux.Handle("/", Middleware(router.Handle()))

	return http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
