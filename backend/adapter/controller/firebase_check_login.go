package controller

import (
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/driver/log"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

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
