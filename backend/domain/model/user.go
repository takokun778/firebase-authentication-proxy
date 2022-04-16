package model

import (
	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"
)

type User struct {
	primitive user.Primitive
}

func NewUser(primitive user.Primitive) *User {
	return &User{
		primitive: primitive,
	}
}
