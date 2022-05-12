package controller

import (
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

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
