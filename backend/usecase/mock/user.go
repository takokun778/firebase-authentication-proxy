package mock

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"
)

var ErrMockUserRepository = errors.New("")

type UserRepository struct {
	WantSaveError   bool
	WantDeleteError bool
}

func (r *UserRepository) Save(ctx context.Context) (user.Primitive, error) {
	if r.WantSaveError {
		return user.Primitive{}, ErrMockUserRepository
	}

	id := uuid.New()

	uid, _ := user.NewID(id.String())

	return user.NewPrimitive(uid), nil
}

func (r *UserRepository) Delete(ctx context.Context, id user.ID) error {
	if r.WantDeleteError {
		return ErrMockUserRepository
	}

	return nil
}
