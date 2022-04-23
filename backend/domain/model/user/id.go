package user

import (
	"github.com/google/uuid"
)

type ID uuid.UUID

func NewID(value string) (ID, error) {
	i, err := uuid.Parse(value)
	if err != nil {
		return ID(uuid.UUID{}), err
	}

	return ID(i), nil
}

func (i ID) Value() uuid.UUID {
	return uuid.UUID(i)
}

func (i ID) String() string {
	return i.Value().String()
}
