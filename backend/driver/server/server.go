package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
)

const shutdownTime = 10

type HTTPServer struct {
	*http.Server
}

func NewHTTPServer(
	key *controller.KeyFetchPublicController,
	register *controller.FirebaseRegisterController,
	login *controller.FirebaseLoginController,
	changePassword *controller.FirebaseChangePasswordController,
	check *controller.FirebaseCheckLoginController,
	logout *controller.FirebaseLogoutController,
	withdraw *controller.FirebaseWithdrawController,
	authorize *controller.FirebaseAuthorizeController,
) *HTTPServer {
	mux := http.NewServeMux()

	mux.Handle("/api/key", Middleware(http.HandlerFunc(key.Get)))
	mux.Handle("/api/register", Middleware(http.HandlerFunc(register.Post)))
	mux.Handle("/api/login", Middleware(http.HandlerFunc(login.Post)))
	mux.Handle("/api/password", Middleware(http.HandlerFunc(changePassword.Put)))
	mux.Handle("/api/login/check", Middleware(http.HandlerFunc(check.Post)))
	mux.Handle("/api/logout", Middleware(http.HandlerFunc(logout.Post)))
	mux.Handle("/api/withdraw", Middleware(http.HandlerFunc(withdraw.Post)))
	mux.Handle("/api/authorize", Middleware(http.HandlerFunc(authorize.Post)))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	return &HTTPServer{
		Server: s,
	}
}

func (s *HTTPServer) Run() {
	go func() {
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("Server closed with error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime*time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Println("Failed to gracefully shutdown:", err)
	}

	log.Println("HTTPServer shutdown")
}
