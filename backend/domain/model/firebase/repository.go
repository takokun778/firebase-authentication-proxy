package firebase

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"
)

type Tokens struct {
	Uid          Uid
	AccessToken  AccessToken
	RefreshToken RefreshToken
	Expires      int
}

type Repository interface {
	Save(context.Context, user.Id, Email, Password) error
	GenerateTokens(context.Context, Email, Password) (Tokens, error)
	ChangePassword(context.Context, Uid, Password) error
	Verify(context.Context, AccessToken) error
	Delete(context.Context, Uid) error
}
