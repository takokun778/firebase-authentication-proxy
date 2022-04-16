package model

import (
	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
)

type Firebase struct {
	primitive firebase.Primitive
}

func NewFirebase(primitive firebase.Primitive) *Firebase {
	return &Firebase{
		primitive: primitive,
	}
}
