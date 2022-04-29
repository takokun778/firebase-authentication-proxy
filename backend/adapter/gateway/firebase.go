package gateway

import (
	"context"
	"time"

	"github.com/takokun778/firebase-authentication-proxy/adapter"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"
	"github.com/takokun778/firebase-authentication-proxy/driver/log"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type FirebaseGateway struct {
	client FirebaseClient
}

func NewFirebaseGateway(client FirebaseClient) port.FirebaseRepository {
	return &FirebaseGateway{
		client: client,
	}
}

func (g *FirebaseGateway) Save(
	ctx context.Context,
	userID user.ID,
	email firebase.Email,
	password firebase.Password,
) error {
	st := time.Now()

	err := g.client.CreateUser(ctx, userID.Value().String(), email.Value(), password.Value())

	log.Elapsed(ctx, st, "create fireabase user")

	if err != nil {
		log.WithCtx(ctx).Warn(err.Error())

		return adapter.NewBadRequestError("bad request")
	}

	return nil
}

func (g *FirebaseGateway) Login(
	ctx context.Context,
	email firebase.Email,
	password firebase.Password,
) (port.Tokens, error) {
	st := time.Now()

	tokens, err := g.client.Login(ctx, email.Value(), password.Value())

	log.Elapsed(ctx, st, "post fireabase login")

	if err != nil {
		log.WithCtx(ctx).Sugar().Debugf("%+v", err)

		return port.Tokens{}, adapter.NewBadRequestError("bad request")
	}

	log.WithCtx(ctx).Debug(tokens.Access)

	access, err := firebase.NewAccessToken(tokens.Access)
	if err != nil {
		return port.Tokens{}, adapter.NewBadRequestError("bad request")
	}

	refresh, err := firebase.NewRefreshToken(tokens.Refresh)
	if err != nil {
		return port.Tokens{}, adapter.NewBadRequestError("bad request")
	}

	log.WithCtx(ctx).Debug(access.Value())

	return port.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (g *FirebaseGateway) ChangePassword(ctx context.Context, uid firebase.UID, password firebase.Password) error {
	if err := g.client.ChangePassword(ctx, uid.Value().String(), password.Value()); err != nil {
		return err
	}

	return nil
}

func (g *FirebaseGateway) Verify(ctx context.Context, token firebase.AccessToken) error {
	st := time.Now()

	err := g.client.VerifyIDToken(ctx, token.Value())

	log.Elapsed(ctx, st, "verify fireabase id token")

	if err != nil {
		return err
	}

	return nil
}

func (g *FirebaseGateway) Delete(ctx context.Context, uid firebase.UID) error {
	if err := g.client.DeleteUser(ctx, uid.Value().String()); err != nil {
		log.WithCtx(ctx).Warn(err.Error())

		return err
	}

	return nil
}
