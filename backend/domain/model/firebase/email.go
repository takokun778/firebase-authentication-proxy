package firebase

import (
	"fmt"
	"regexp"

	"firebase-authentication/domain/model/errors"
)

type Email string

const email = `^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`

func NewEmail(value string) (Email, error) {
	v := Email(value)
	if err := v.validate(); err != nil {
		return Email(""), err
	}
	return Email(value), nil
}

func (p Email) Value() string {
	return string(p)
}

func (p Email) validate() error {
	if !regexp.MustCompile(email).MatchString(p.Value()) {
		return errors.NewErrBadRequest(fmt.Sprintf("validate error email: %s", p.Value()), nil)
	}
	return nil
}
