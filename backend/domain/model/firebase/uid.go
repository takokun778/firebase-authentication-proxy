package firebase

import (
	"github.com/google/uuid"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"
)

type UID uuid.UUID

func NewUID(value string) (UID, error) {
	i, err := uuid.Parse(value)
	if err != nil {
		return UID(uuid.UUID{}), errors.NewValidateError(err.Error())
	}

	return UID(i), nil
}

func (i UID) Value() uuid.UUID {
	return uuid.UUID(i)
}

func (i UID) String() string {
	return i.Value().String()
}

func (i UID) ToUserID() (user.ID, error) {
	id, err := user.NewID(i.String())
	if err != nil {
		return user.ID{}, errors.NewValidateError(err.Error())
	}

	return id, nil
}
