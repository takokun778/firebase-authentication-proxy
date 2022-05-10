package controller

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/key"
	"github.com/takokun778/firebase-authentication-proxy/driver/log"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseChangePasswordController struct {
	input  port.FirebaseChangePasswordInputPort
	output port.FirebaseChangePasswordOutputPort
}

func NewFirebaseChangePasswordController(
	input port.FirebaseChangePasswordInputPort,
	output port.FirebaseChangePasswordOutputPort,
) *FirebaseChangePasswordController {
	return &FirebaseChangePasswordController{
		input:  input,
		output: output,
	}
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func (c *FirebaseChangePasswordController) Put(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(adapter.SetResWriter(r.Context(), w))

	if r.Method != http.MethodPut {
		c.output.ErrorRender(r.Context(), adapter.NewMethodNotAllowedError())

		return
	}

	atc, _ := r.Cookie("access-token")

	if atc == nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(""))

		return
	}

	var req ChangePasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	op, err := base64.StdEncoding.DecodeString(req.OldPassword)
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	opv, err := key.Decrypt(op)
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	np, err := base64.StdEncoding.DecodeString(req.NewPassword)
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	npv, err := key.Decrypt(np)
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	input, err := port.NewFirebaseChangePasswordInput(atc.Value, string(opv), string(npv))
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	c.input.Execute(r.Context(), input)
}

type FirebaseCheckLoginController struct {
	input  port.FirebaseCheckLoginInputPort
	output port.FirebaseCheckLoginOutputPort
}

func NewFirebaseCheckLoginController(
	input port.FirebaseCheckLoginInputPort,
	output port.FirebaseCheckLoginOutputPort,
) *FirebaseCheckLoginController {
	return &FirebaseCheckLoginController{
		input:  input,
		output: output,
	}
}

func (c *FirebaseCheckLoginController) Post(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(adapter.SetResWriter(r.Context(), w))

	if r.Method != http.MethodPost {
		c.output.ErrorRender(r.Context(), adapter.NewMethodNotAllowedError())

		return
	}

	atc, _ := r.Cookie("access-token")

	if atc == nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(""))

		return
	}

	ftc, _ := r.Cookie("refresh-token")

	if ftc == nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(""))

		return
	}

	input, err := port.NewFirebaseCheckLoginInput(atc.Value, ftc.Value)
	if err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(""))

		return
	}

	c.input.Execute(r.Context(), input)
}

type FirebaseLogoutController struct {
	input  port.FirebaseLogoutInputPort
	output port.FirebaseLogoutOutputPort
}

func NewFirebaseLogoutController(
	input port.FirebaseLogoutInputPort,
	output port.FirebaseLogoutOutputPort,
) *FirebaseLogoutController {
	return &FirebaseLogoutController{
		input:  input,
		output: output,
	}
}

func (c *FirebaseLogoutController) Post(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(adapter.SetResWriter(r.Context(), w))

	if r.Method != http.MethodPost {
		c.output.ErrorRender(r.Context(), adapter.NewMethodNotAllowedError())

		return
	}

	atc, _ := r.Cookie("access-token")

	if atc == nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(""))

		return
	}

	ftc, _ := r.Cookie("refresh-token")

	if ftc == nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(""))

		return
	}

	input, err := port.NewFirebaseLogoutInput(atc.Value, ftc.Value)
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	c.input.Execute(r.Context(), input)
}

type FirebaseWithdrawController struct {
	input  port.FirebaseWithdrawInputPort
	output port.FirebaseWithdrawOutputPort
}

func NewFirebaseWithdrawController(
	input port.FirebaseWithdrawInputPort,
	output port.FirebaseWithdrawOutputPort,
) *FirebaseWithdrawController {
	return &FirebaseWithdrawController{
		input:  input,
		output: output,
	}
}

type WithdrawRequest struct {
	Password string `json:"password"`
}

func (c *FirebaseWithdrawController) Post(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(adapter.SetResWriter(r.Context(), w))

	if r.Method != http.MethodPost {
		c.output.ErrorRender(r.Context(), adapter.NewMethodNotAllowedError())

		return
	}

	atc, _ := r.Cookie("access-token")

	if atc == nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(""))

		return
	}

	var req WithdrawRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	password, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	pv, err := key.Decrypt(password)
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	input, err := port.NewFirebaseWithdrawInput(atc.Value, string(pv))
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	c.input.Execute(r.Context(), input)
}

type FirebaseAuthorizeController struct {
	input  port.FirebaseAuthorizeInputPort
	output port.FirebaseAuthorizeOutputPort
}

func NewFirebaseAuthorizeController(
	input port.FirebaseAuthorizeInputPort,
	output port.FirebaseAuthorizeOutputPort,
) *FirebaseAuthorizeController {
	return &FirebaseAuthorizeController{
		input:  input,
		output: output,
	}
}

func (c *FirebaseAuthorizeController) Post(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(adapter.SetResWriter(r.Context(), w))

	idToken := r.Header.Get("Authorization")

	if idToken == "" {
		c.output.ErrorRender(r.Context(), errors.NewUnauthorizedError(""))

		return
	}

	input, err := port.NewFirebaseAuthorizeInput(idToken)
	if err != nil {
		c.output.ErrorRender(r.Context(), errors.NewUnauthorizedError(""))

		return
	}

	c.input.Execute(r.Context(), input)
}
