package firebase

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

type AccessToken string

func NewAccessToken(value string) (AccessToken, error) {
	v := AccessToken(value)
	if err := v.validate(); err != nil {
		return AccessToken(""), err
	}

	return v, nil
}

func (t AccessToken) Value() string {
	return string(t)
}

func (t AccessToken) validate() error {
	return nil
}

func (t AccessToken) GetUID() (UID, error) {
	payload := strings.Split(t.Value(), ".")[1]

	decode, _ := base64.RawURLEncoding.DecodeString(payload)

	var c claims

	if err := json.Unmarshal(decode, &c); err != nil {
		return UID(uuid.UUID{}), err
	}

	return NewUID(c.UID)
}

func (t AccessToken) GetEmail() (Email, error) {
	payload := strings.Split(t.Value(), ".")[1]

	decode, _ := base64.RawURLEncoding.DecodeString(payload)

	var c claims

	if err := json.Unmarshal(decode, &c); err != nil {
		return Email(""), err
	}

	return Email(c.Email), nil
}

type claims struct {
	UID   string `json:"user_id"`
	Email string `json:"email"`
}
