package repository

import (
	"github.com/rianekacahya/boilerplate/domain/bootstrap"
)

type oauthRepository struct {
	dependency bootstrap.Dependency
}

func NewOauthRepository(dependency bootstrap.Dependency) *oauthRepository {
	return &oauthRepository{dependency}
}
