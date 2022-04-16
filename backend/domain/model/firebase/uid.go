package firebase

import (
	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"

	"github.com/google/uuid"
)

type Uid uuid.UUID

func NewUid(value string) (Uid, error) {
	i, err := uuid.Parse(value)
	if err != nil {
		return Uid(uuid.UUID{}), err
	}
	return Uid(i), nil
}

func (i Uid) Value() uuid.UUID {
	return uuid.UUID(i)
}

func (i Uid) String() string {
	return i.Value().String()
}

func (i Uid) ToUserId() (user.Id, error) {
	return user.NewId(i.String())
}
