package model

import (
	"firebase-authentication/domain/model/firebase"
)

type Firebase struct {
	primitive firebase.Primitive
}

func NewFirebase(primitive firebase.Primitive) *Firebase {
	return &Firebase{
		primitive: primitive,
	}
}
