package model

import (
	"firebase-authentication/domain/model/user"
)

type User struct {
	primitive user.Primitive
}

func NewUser(primitive user.Primitive) *User {
	return &User{
		primitive: primitive,
	}
}
