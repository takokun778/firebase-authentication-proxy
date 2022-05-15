package controller

import (
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

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
