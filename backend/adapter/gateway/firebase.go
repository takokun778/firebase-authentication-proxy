package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"
	df "github.com/takokun778/firebase-authentication-proxy/driver/firebase"
	"github.com/takokun778/firebase-authentication-proxy/driver/log"

	"firebase.google.com/go/auth"
)

type FirebaseGateway struct {
	admin *auth.Client
	api   *df.ApiRestClient
}

func NewFirebaseGateway(admin *auth.Client, api *df.ApiRestClient) firebase.Repository {
	return &FirebaseGateway{
		admin: admin,
		api:   api,
	}
}

func (g *FirebaseGateway) Save(ctx context.Context, userId user.Id, email firebase.Email, password firebase.Password) error {
	params := (&auth.UserToCreate{}).
		UID(userId.String()).
		Email(email.Value()).
		EmailVerified(false).
		Password(password.Value()).
		Disabled(false)

	st := time.Now()

	_, err := g.admin.CreateUser(ctx, params)

	log.Elapsed(ctx, st, "create fireabase user")

	if err != nil {
		log.WithCtx(ctx).Warn(err.Error())
		return errors.NewErrBadRequest("bad request", nil)
	}

	return nil
}

type SignInRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type SignInResponse struct {
	ExpiresIn    string `json:"expiresIn"`
	LocalId      string `json:"localId"`
	IdToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}

func (g *FirebaseGateway) GenerateTokens(ctx context.Context, email firebase.Email, password firebase.Password) (firebase.Tokens, error) {
	url := fmt.Sprintf("%s/v1/accounts:signInWithPassword?key=%s", g.api.Endpoint, g.api.ApiKey)

	req := SignInRequest{
		Email:             email.Value(),
		Password:          password.Value(),
		ReturnSecureToken: true,
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(req)

	if err != nil {
		return firebase.Tokens{}, errors.NewErrBadRequest("bad request", nil)
	}

	st := time.Now()

	res, err := g.api.Post(url, "application/json", &buf)

	log.Elapsed(ctx, st, "post fireabase login")

	if err != nil {
		log.WithCtx(ctx).Sugar().Debugf("%+v", err)
		return firebase.Tokens{}, errors.NewErrBadRequest("bad request", nil)
	}

	code := res.StatusCode

	if code != 200 {
		log.WithCtx(ctx).Sugar().Debugf("%+v", res.Status)
		return firebase.Tokens{}, errors.NewErrBadRequest("bad request", nil)
	}

	var response SignInResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return firebase.Tokens{}, errors.NewErrBadRequest("bad request", nil)
	}

	log.WithCtx(ctx).Sugar().Debugf("%+v", response)

	uid, err := firebase.NewUid(response.LocalId)

	if err != nil {
		return firebase.Tokens{}, errors.NewErrBadRequest("bad request", nil)
	}

	access, err := firebase.NewAccessToken(response.IdToken)

	if err != nil {
		return firebase.Tokens{}, errors.NewErrBadRequest("bad request", nil)
	}

	refresh, err := firebase.NewRefreshToken(response.RefreshToken)

	if err != nil {
		return firebase.Tokens{}, errors.NewErrBadRequest("bad request", nil)
	}

	expires, err := strconv.Atoi(response.ExpiresIn)

	if err != nil {
		return firebase.Tokens{}, errors.NewErrBadRequest("bad request", nil)
	}

	return firebase.Tokens{
		Uid:          uid,
		AccessToken:  access,
		RefreshToken: refresh,
		Expires:      expires,
	}, nil
}

func (g *FirebaseGateway) ChangePassword(ctx context.Context, uid firebase.Uid, password firebase.Password) error {
	params := (&auth.UserToUpdate{}).
		Password(password.Value())

	if _, err := g.admin.UpdateUser(ctx, uid.String(), params); err != nil {
		return err
	}

	return nil
}

func (g *FirebaseGateway) Verify(ctx context.Context, token firebase.AccessToken) error {
	st := time.Now()

	_, err := g.admin.VerifyIDToken(ctx, token.Value())

	log.Elapsed(ctx, st, "verify fireabase id token")

	if err != nil {
		return err
	}

	return nil
}

func (g *FirebaseGateway) Delete(ctx context.Context, uid firebase.Uid) error {
	if err := g.admin.DeleteUser(ctx, uid.String()); err != nil {
		log.WithCtx(ctx).Warn(err.Error())
		return err
	}

	return nil
}
