package repository

import (
	"context"
	"github.com/rianekacahya/boilerplate/domain/entity"
)

type Oauth interface {
	GetClientByClientID(context.Context, string) (*entity.Clients, error)
	CheckSessionExist(ctx context.Context, key string) (bool, error)
	GetSession(ctx context.Context, key string) (*entity.Session, error)
}
