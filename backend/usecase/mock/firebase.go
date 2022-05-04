package mock

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

var ErrMockFirebaseRepository = errors.New("")

type FirebaseRepository struct {
	WantSaveError           bool
	WantLoginError          bool
	WantChangePasswordError bool
	WantVerifyError         bool
	WantDeleteError         bool
}

func (r *FirebaseRepository) Save(
	ctx context.Context,
	userID user.ID,
	email firebase.Email,
	password firebase.Password,
) error {
	if r.WantSaveError {
		return ErrMockFirebaseRepository
	}

	return nil
}

func (r *FirebaseRepository) Login(
	ctx context.Context,
	email firebase.Email,
	password firebase.Password,
) (port.Tokens, error) {
	if r.WantLoginError {
		return port.Tokens{}, ErrMockFirebaseRepository
	}

	return port.Tokens{
		AccessToken:  firebase.AccessToken(CreateTestJwt(uuid.New().String(), "test@example.com")),
		RefreshToken: firebase.RefreshToken(""),
	}, nil
}

func (r *FirebaseRepository) ChangePassword(
	ctx context.Context,
	uid firebase.UID,
	password firebase.Password,
) error {
	if r.WantChangePasswordError {
		return ErrMockFirebaseRepository
	}

	return nil
}

func (r *FirebaseRepository) Verify(
	ctx context.Context,
	token firebase.AccessToken,
) error {
	if r.WantVerifyError {
		return ErrMockFirebaseRepository
	}

	return nil
}

func (r *FirebaseRepository) Delete(
	ctx context.Context,
	uid firebase.UID,
) error {
	if r.WantDeleteError {
		return ErrMockFirebaseRepository
	}

	return nil
}

func CreateTestJwt(userID, email string) string {
	// https://firebase.google.com/docs/auth/admin/verify-id-tokens#verify_id_tokens_using_a_third-party_jwt_library
	token := jwt.New(jwt.SigningMethodRS256)

	claims, _ := token.Claims.(jwt.MapClaims)

	now := time.Now().UTC()

	claims["iss"] = "iss"
	claims["aud"] = "aud"
	claims["auth_time"] = now.Unix()
	claims["user_id"] = userID
	claims["sub"] = uuid.New()
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(time.Hour).Unix()
	claims["email"] = email
	claims["email_verified"] = false
	claims["firebase"] = ""

	jwt, err := token.SigningString()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return jwt
}
