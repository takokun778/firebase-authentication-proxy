package firebase

import (
	"fmt"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
)

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
	// 5文字以化はNG
	if len(p.Value()) <= 5 {
		fmt.Println("hoge")
		return errors.NewErrBadRequest("password is too short", nil)
	}
	// 大文字のアルファベットを1つ含む
	// 記号をひとつ含む
	return nil
}
