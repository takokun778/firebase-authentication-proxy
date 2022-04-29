package gateway

import (
	"context"
)

type Tokens struct {
	Access  string
	Refresh string
}

type FirebaseClient interface {
	CreateUser(ctx context.Context, userID, email, password string) error
	Login(ctx context.Context, email, password string) (Tokens, error)
	ChangePassword(ctx context.Context, uid, password string) error
	VerifyIDToken(ctx context.Context, accessToken string) error
	DeleteUser(ctx context.Context, uid string) error
}
