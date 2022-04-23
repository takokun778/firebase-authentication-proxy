package gateway

import (
	"context"

	"github.com/google/uuid"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/user"
)

type UserGateway struct{}

func NewUserGateway() user.Repository {
	return &UserGateway{}
}

func (g *UserGateway) Save(ctx context.Context) (user.Primitive, error) {
	return user.Primitive{}, nil
}

func (g *UserGateway) Delete(ctx context.Context, id user.ID) error {
	return nil
}

type UserInMemory struct{}

func NewUserInMemory() user.Repository {
	return &UserInMemory{}
}

func (*UserInMemory) Save(ctx context.Context) (user.Primitive, error) {
	id := uuid.New()
	uid, _ := user.NewID(id.String())
	u := user.NewPrimitive(uid)

	return u, nil
}

func (*UserInMemory) Delete(ctx context.Context, id user.ID) error {
	return nil
}
