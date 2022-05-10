package controller

import (
	"net/http"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type KeyFetchController struct {
	input  port.KeyFetchInputPort
	output port.KeyFetchOutputPort
}

func NewKeyFetchController(
	input port.KeyFetchInputPort,
	output port.KeyFetchOutputPort,
) *KeyFetchController {
	return &KeyFetchController{
		input:  input,
		output: output,
	}
}

func (c *KeyFetchController) Get(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(adapter.SetResWriter(r.Context(), w))

	if r.Method != http.MethodGet {
		c.output.ErrorRender(r.Context(), adapter.NewMethodNotAllowedError())

		return
	}

	c.input.Execute(r.Context())
}
