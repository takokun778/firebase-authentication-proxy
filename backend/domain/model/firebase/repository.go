package firebase

import (
	"context"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"
)

type Tokens struct {
	UID          UID
	AccessToken  AccessToken
	RefreshToken RefreshToken
	Expires      int
}

type Repository interface {
	Save(context.Context, user.ID, Email, Password) error
	Login(context.Context, Email, Password) (Tokens, error)
	ChangePassword(context.Context, UID, Password) error
	Verify(context.Context, AccessToken) error
	Delete(context.Context, UID) error
}
