package usecase

import (
	"context"
	"github.com/rianekacahya/boilerplate/domain/entity"
)

type Oauth interface{
	Token(ctx context.Context, req *entity.RequestToken) (*entity.Token, error)
	RefreshToken(context.Context, *entity.RequestToken) (*entity.Token, error)
}
