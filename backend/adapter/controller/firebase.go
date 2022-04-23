package controller

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/driver/log"
	"github.com/takokun778/firebase-authentication-proxy/usecase"
)

type FirebaseController struct {
	firebaseInteractor usecase.FirebaseInputPort
}

func NewFirebaseController(input usecase.FirebaseInputPort) *FirebaseController {
	return &FirebaseController{
		firebaseInteractor: input,
	}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *FirebaseController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	password, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	input, err := usecase.NewFirebaseRegisterInput(req.Email, password)
	if err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	c.firebaseInteractor.Register(r.Context(), input)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *FirebaseController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	password, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	input, err := usecase.NewFirebaseLoginInput(req.Email, password)
	if err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	c.firebaseInteractor.Login(r.Context(), input)
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func (c *FirebaseController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	atc, _ := r.Cookie("access-token")

	if atc == nil {
		log.WithCtx(r.Context()).Warn("access token is nil")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var req ChangePasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	op, err := base64.StdEncoding.DecodeString(req.OldPassword)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	np, err := base64.StdEncoding.DecodeString(req.NewPassword)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	input, err := usecase.NewFirebaseChangePasswordInput(atc.Value, op, np)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	c.firebaseInteractor.ChangePassword(r.Context(), input)
}

func (c *FirebaseController) CheckLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	atc, _ := r.Cookie("access-token")

	if atc == nil {
		log.WithCtx(r.Context()).Warn("access token is nil")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	ftc, _ := r.Cookie("refresh-token")

	if ftc == nil {
		log.WithCtx(r.Context()).Warn("refresh token is nil")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	input, err := usecase.NewFirebaseCheckLoginInput(atc.Value, ftc.Value)
	if err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	c.firebaseInteractor.CheckLogin(r.Context(), input)
}

func (c *FirebaseController) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	atc, _ := r.Cookie("access-token")

	if atc == nil {
		log.WithCtx(r.Context()).Warn("access token is nil")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	ftc, _ := r.Cookie("refresh-token")

	if ftc == nil {
		log.WithCtx(r.Context()).Warn("refresh token is nil")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	input, err := usecase.NewFirebaseLogoutInput(atc.Value, ftc.Value)
	if err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	c.firebaseInteractor.Logout(r.Context(), input)
}

type WithdrawRequest struct {
	Password string `json:"password"`
}

func (c *FirebaseController) Withdraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	atc, _ := r.Cookie("access-token")

	if atc == nil {
		log.WithCtx(r.Context()).Warn("access token is nil")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var req WithdrawRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	password, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	input, err := usecase.NewFirebaseWithdrawInput(atc.Value, password)
	if err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	c.firebaseInteractor.Withdraw(r.Context(), input)
}

func (c *FirebaseController) Authorize(w http.ResponseWriter, r *http.Request) {
	idToken := r.Header.Get("Authorization")

	if idToken == "" {
		log.WithCtx(r.Context()).Warn("no authorization header")
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	input, err := usecase.NewFirebaseAuthorizeInput(idToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	c.firebaseInteractor.Authorize(r.Context(), input)
}
