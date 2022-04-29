package firebase

import (
	"github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
)

const minLength = 6

type Password string

func NewPassword(value string) (Password, error) {
	v := Password(value)
	if err := v.validate(); err != nil {
		return Password(""), err
	}

	return Password(value), nil
}

func (p Password) Value() string {
	return string(p)
}

func (p Password) Equals(other Password) bool {
	return p.Value() == other.Value()
}

func (p Password) validate() error {
	if len(p.Value()) < minLength {
		return errors.NewValidateError("password is too short")
	}

	return nil
}
