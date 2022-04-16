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

func (t AccessToken) GetUid() (Uid, error) {
	payload := strings.Split(t.Value(), ".")[1]

	decode, _ := base64.RawURLEncoding.DecodeString(payload)

	var c claims

	if err := json.Unmarshal(decode, &c); err != nil {
		return Uid(uuid.UUID{}), err
	}

	return NewUid(c.Uid)
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
	Uid   string `json:"user_id"`
	Email string `json:"email"`
}
