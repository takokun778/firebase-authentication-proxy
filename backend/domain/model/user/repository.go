package user

import (
	"context"
)

type Repository interface {
	Save(context.Context) (Primitive, error)
	Delete(context.Context, ID) error
}
