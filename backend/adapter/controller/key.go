package controller

import (
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type KeyFetchPublicController struct {
	input  port.KeyFetchPublicInputPort
	output port.KeyFetchPublicOutputPort
}

func NewKeyFetchPublicController(
	input port.KeyFetchPublicInputPort,
	output port.KeyFetchPublicOutputPort,
) *KeyFetchPublicController {
	return &KeyFetchPublicController{
		input:  input,
		output: output,
	}
}

func (c *KeyFetchPublicController) Get(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(adapter.SetResWriter(r.Context(), w))

	if r.Method != http.MethodGet {
		c.output.ErrorRender(r.Context(), adapter.NewMethodNotAllowedError())

		return
	}

	c.input.Execute(r.Context())
}
