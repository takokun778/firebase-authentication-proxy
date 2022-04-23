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

const perm = 0o600

// @see https://firebase.google.com/docs/reference/rest/auth
type APIRestClient struct {
	*http.Client
	Endpoint string
	APIKey   string
}

type Client struct {
	Admin *auth.Client
	API   *APIRestClient
}

type secret struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURL                 string `json:"auth_url"`
	TokenURL                string `json:"token_url"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert"`
	ClientX509CertURL       string `json:"client_x509_cert"`
}

func NewFirebaseClient() *Client {
	s := secret{
		Type:                    os.Getenv("TYPE"),
		ProjectID:               os.Getenv("PROJECT_ID"),
		PrivateKeyID:            os.Getenv("PRIVATE_KEY_ID"),
		PrivateKey:              os.Getenv("PRIVATE_KEY"),
		ClientEmail:             os.Getenv("CLIENT_EMAIL"),
		ClientID:                os.Getenv("CLIENT_ID"),
		AuthURL:                 os.Getenv("AUTH_URL"),
		TokenURL:                os.Getenv("TOKEN_URL"),
		AuthProviderX509CertURL: os.Getenv("AUTH_PROVIDER_X509_CERT_URL"),
		ClientX509CertURL:       os.Getenv("CLIENT_X509_CERT_URL"),
	}

	file, _ := json.MarshalIndent(s, "", "")

	if err := ioutil.WriteFile("secret.json", file, perm); err != nil {
		log.Fatalf("error write file secret.json: %v", err)
	}

	opt := option.WithCredentialsFile("./secret.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	admin, err := app.Auth(context.Background())
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

	api := &APIRestClient{
		Client:   &http.Client{},
		Endpoint: endpoint,
		APIKey:   key,
	}

	return &Client{
		Admin: admin,
		API:   api,
	}
}
