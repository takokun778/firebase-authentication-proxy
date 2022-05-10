package controller

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/key"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseLoginController struct {
	input  port.FirebaseLoginInputPort
	output port.FirebaseLoginOutputPort
}

func NewFirebaseLoginController(
	input port.FirebaseLoginInputPort,
	output port.FirebaseLoginOutputPort,
) *FirebaseLoginController {
	return &FirebaseLoginController{
		input:  input,
		output: output,
	}
}

type FirebaseLoginPostBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *FirebaseLoginController) Post(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(adapter.SetResWriter(r.Context(), w))

	if r.Method != http.MethodPost {
		c.output.ErrorRender(r.Context(), adapter.NewMethodNotAllowedError())

		return
	}

	var body FirebaseLoginPostBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	password, err := base64.StdEncoding.DecodeString(body.Password)
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	pv, err := key.Decrypt(password)
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(""))

		return
	}

	input, err := port.NewFirebaseLoginInput(body.Email, string(pv))
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	c.input.Execute(r.Context(), input)
}
