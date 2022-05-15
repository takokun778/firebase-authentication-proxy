package controller

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/key"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

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

type FirebaseWithdrawPostBody struct {
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

	var body FirebaseWithdrawPostBody

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
