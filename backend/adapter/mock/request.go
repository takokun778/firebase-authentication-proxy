package mock

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/key"
)

func Encrypt(value string) string {
	enc, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, key.GetPublic(), []byte(value), nil)
	return base64.StdEncoding.EncodeToString(enc)
}

func CreateFirebaseRegisterPostBody(body controller.FirebaseRegisterPostBody) io.ReadCloser {
	pr, pw := io.Pipe()

	go func() {
		_ = json.NewEncoder(pw).Encode(&body)
		pw.Close()
	}()

	return pr
}

func CreateFirebaseLoginPostBody(body controller.FirebaseLoginPostBody) io.ReadCloser {
	pr, pw := io.Pipe()

	go func() {
		_ = json.NewEncoder(pw).Encode(&body)
		pw.Close()
	}()

	return pr
}
