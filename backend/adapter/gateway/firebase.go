package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"firebase.google.com/go/auth"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"
	df "github.com/takokun778/firebase-authentication-proxy/driver/firebase"
	"github.com/takokun778/firebase-authentication-proxy/driver/log"
)

type FirebaseGateway struct {
	client *df.Client
}

func NewFirebaseGateway(client *df.Client) firebase.Repository {
	return &FirebaseGateway{
		client: client,
	}
}

func (g *FirebaseGateway) Save(ctx context.Context, userID user.ID, email firebase.Email, password firebase.Password) error {
	params := (&auth.UserToCreate{}).
		UID(userID.String()).
		Email(email.Value()).
		EmailVerified(false).
		Password(password.Value()).
		Disabled(false)

	st := time.Now()

	_, err := g.client.Admin.CreateUser(ctx, params)

	log.Elapsed(ctx, st, "create fireabase user")

	if err != nil {
		log.WithCtx(ctx).Warn(err.Error())

		return errors.NewBadRequestError("bad request")
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
	LocalID      string `json:"localID"`
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}

func (g *FirebaseGateway) Login(ctx context.Context, email firebase.Email, password firebase.Password) (firebase.Tokens, error) {
	url := fmt.Sprintf("%s/v1/accounts:signInWithPassword?key=%s", g.client.API.Endpoint, g.client.API.APIKey)

	req := SignInRequest{
		Email:             email.Value(),
		Password:          password.Value(),
		ReturnSecureToken: true,
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(req)
	if err != nil {
		return firebase.Tokens{}, errors.NewBadRequestError("bad request")
	}

	st := time.Now()

	res, err := g.client.API.Post(url, "application/json", &buf)

	log.Elapsed(ctx, st, "post fireabase login")

	if err != nil {
		log.WithCtx(ctx).Sugar().Debugf("%+v", err)

		return firebase.Tokens{}, errors.NewBadRequestError("bad request")
	}

	if res.StatusCode != http.StatusOK {
		log.WithCtx(ctx).Sugar().Debugf("%+v", res.Status)

		return firebase.Tokens{}, errors.NewBadRequestError("bad request")
	}

	var response SignInResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return firebase.Tokens{}, errors.NewBadRequestError("bad request")
	}

	log.WithCtx(ctx).Sugar().Debugf("%+v", response)

	uid, err := firebase.NewUID(response.LocalID)
	if err != nil {
		return firebase.Tokens{}, errors.NewBadRequestError("bad request")
	}

	access, err := firebase.NewAccessToken(response.IDToken)
	if err != nil {
		return firebase.Tokens{}, errors.NewBadRequestError("bad request")
	}

	refresh, err := firebase.NewRefreshToken(response.RefreshToken)
	if err != nil {
		return firebase.Tokens{}, errors.NewBadRequestError("bad request")
	}

	expires, err := strconv.Atoi(response.ExpiresIn)
	if err != nil {
		return firebase.Tokens{}, errors.NewBadRequestError("bad request")
	}

	return firebase.Tokens{
		UID:          uid,
		AccessToken:  access,
		RefreshToken: refresh,
		Expires:      expires,
	}, nil
}

func (g *FirebaseGateway) ChangePassword(ctx context.Context, uid firebase.UID, password firebase.Password) error {
	params := (&auth.UserToUpdate{}).
		Password(password.Value())

	if _, err := g.client.Admin.UpdateUser(ctx, uid.String(), params); err != nil {
		return err
	}

	return nil
}

func (g *FirebaseGateway) Verify(ctx context.Context, token firebase.AccessToken) error {
	st := time.Now()

	_, err := g.client.Admin.VerifyIDToken(ctx, token.Value())

	log.Elapsed(ctx, st, "verify fireabase id token")

	if err != nil {
		return err
	}

	return nil
}

func (g *FirebaseGateway) Delete(ctx context.Context, uid firebase.UID) error {
	if err := g.client.Admin.DeleteUser(ctx, uid.String()); err != nil {
		log.WithCtx(ctx).Warn(err.Error())

		return err
	}

	return nil
}
