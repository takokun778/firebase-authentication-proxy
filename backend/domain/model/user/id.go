package user

import (
	"github.com/google/uuid"
)

type Id uuid.UUID

func NewId(value string) (Id, error) {
	i, err := uuid.Parse(value)
	if err != nil {
		return Id(uuid.UUID{}), err
	}
	return Id(i), nil
}

func (i Id) Value() uuid.UUID {
	return uuid.UUID(i)
}

func (i Id) String() string {
	return i.Value().String()
}
