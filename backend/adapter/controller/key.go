package controller

import (
	"net/http"

	"firebase-authentication/driver/log"
	"firebase-authentication/usecase"
)

type KeyController struct {
	keyInteractor usecase.KeyInputPort
}

func NewKeyController(input usecase.KeyInputPort) *KeyController {
	return &KeyController{
		keyInteractor: input,
	}
}

func (c *KeyController) GetPublic(w http.ResponseWriter, r *http.Request) {

	input, err := usecase.NewKeyGetPublicInput()

	if err != nil {
		log.WithCtx(r.Context()).Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c.keyInteractor.GetPublic(r.Context(), input)
}
