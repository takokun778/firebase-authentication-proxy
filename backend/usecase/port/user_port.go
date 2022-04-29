package port

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"
)

type UserRepository interface {
	Save(context.Context) (user.Primitive, error)
	Delete(context.Context, user.ID) error
}
