package mock

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/key"
)

func Encrypt(value string) string {
	enc, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, key.GetPublic(), []byte(value), nil)

	return base64.StdEncoding.EncodeToString(enc)
}

func WithToken(req *http.Request) *http.Request {
	access := &http.Cookie{
		Name:     "access-token",
		Value:    createTestJwt(uuid.New().String(), "test@example.com"),
		Path:     "/",
		Expires:  time.Now().Add(time.Second),
		HttpOnly: true,
		Secure:   true,
	}
	req.AddCookie(access)

	refresh := &http.Cookie{
		Name:     "refresh-token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(time.Second),
		HttpOnly: true,
		Secure:   true,
	}
	req.AddCookie(refresh)

	return req
}

func CreateHeaderWithToken() http.Header {
	header := &http.Header{}
	header.Set("Authorization", createTestJwt(uuid.New().String(), "test@example.com"))

	return *header
}

func createTestJwt(userID, email string) string {
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

func CreateFirebaseRegisterPostBody(body controller.FirebaseRegisterPostBody) io.ReadCloser {
	return encode(body)
}

func CreateFirebaseLoginPostBody(body controller.FirebaseLoginPostBody) io.ReadCloser {
	return encode(body)
}

func CreateFirebaseChangePasswordPutBody(body controller.FirebaseChangePasswordPutBody) io.ReadCloser {
	return encode(body)
}

func CreateFirebaseWithdrawPostBody(body controller.FirebaseWithdrawPostBody) io.ReadCloser {
	return encode(body)
}

func encode[T any](body T) *io.PipeReader {
	pr, pw := io.Pipe()

	go func() {
		_ = json.NewEncoder(pw).Encode(&body)
		pw.Close()
	}()

	return pr
}
