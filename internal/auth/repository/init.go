package repository

import (
	"github.com/rianekacahya/boilerplate/domain/bootstrap"
)

type authRepository struct {
	dependency bootstrap.Dependency
}

func NewAuthRepository(dependency bootstrap.Dependency) *authRepository {
	return &authRepository{dependency}
}
