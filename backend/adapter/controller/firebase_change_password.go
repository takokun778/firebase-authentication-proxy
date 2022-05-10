package controller

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/key"
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

type FirebaseChangePasswordPutBody struct {
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

	var body FirebaseChangePasswordPutBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	op, err := base64.StdEncoding.DecodeString(body.OldPassword)
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	opv, err := key.Decrypt(op)
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	np, err := base64.StdEncoding.DecodeString(body.NewPassword)
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
