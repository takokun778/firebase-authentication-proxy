package mock

import (
	"encoding/json"
	"io"

	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
)

func CreateFirebaseRegisterPostBody(body controller.FirebaseRegisterPostBody) io.ReadCloser {
	pr, pw := io.Pipe()

	go func() {
		_ = json.NewEncoder(pw).Encode(&body)
		pw.Close()
	}()

	return pr
}
