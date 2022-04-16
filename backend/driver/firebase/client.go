package firebase

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// @see https://firebase.google.com/docs/reference/rest/auth
type ApiRestClient struct {
	*http.Client
	Endpoint string
	ApiKey   string
}

var (
	AdminClient *auth.Client
	ApiClient   *ApiRestClient
)

type secret struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUrl                 string `json:"auth_url"`
	TokenUrl                string `json:"token_url"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert"`
	ClientX509CertUrl       string `json:"client_x509_cert"`
}

func init() {
	s := secret{
		Type:                    os.Getenv("TYPE"),
		ProjectId:               os.Getenv("PROJECT_ID"),
		PrivateKeyId:            os.Getenv("PRIVATE_KEY_ID"),
		PrivateKey:              os.Getenv("PRIVATE_KEY"),
		ClientEmail:             os.Getenv("CLIENT_EMAIL"),
		ClientId:                os.Getenv("CLIENT_ID"),
		AuthUrl:                 os.Getenv("AUTH_URL"),
		TokenUrl:                os.Getenv("TOKEN_URL"),
		AuthProviderX509CertUrl: os.Getenv("AUTH_PROVIDER_X509_CERT_URL"),
		ClientX509CertUrl:       os.Getenv("CLIENT_X509_CERT_URL"),
	}

	file, _ := json.MarshalIndent(s, "", "")

	if err := ioutil.WriteFile("secret.json", file, 0600); err != nil {
		log.Fatalf("error write file secret.json: %v", err)
	}

	opt := option.WithCredentialsFile("./secret.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	AdminClient, err = app.Auth(context.Background())

	if err != nil {
		log.Fatalf("error create auth client: %v", err)
	}

	if err := os.Remove("./secret.json"); err != nil {
		log.Fatalf("error delete file secret.json: %v", err)
	}

	endpoint := os.Getenv("FIREBASE_API_ENDPOINT")

	if endpoint == "" {
		log.Fatalf("error firebase api endpoint is empty")
	}

	key := os.Getenv("API_KEY")

	if key == "" {
		log.Fatalf("error firebase api key is empty")
	}

	ApiClient = &ApiRestClient{
		Client:   &http.Client{},
		Endpoint: endpoint,
		ApiKey:   key,
	}
}
