package server

import (
	"fmt"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/driver/injector"
)

type Server struct{}

func Serve(port string) error {
	mux := http.NewServeMux()

	router := NewRouter(injector.NewInjector())

	mux.Handle("/", Middleware(router.Handle()))

	return http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
