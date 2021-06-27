package auth

import (
	"github.com/rianekacahya/boilerplate/domain/bootstrap"
)

type authUsecase struct {
	repository bootstrap.Repository
	dependency bootstrap.Dependency
}

func NewAuthUsecase(repository bootstrap.Repository, dependency bootstrap.Dependency) *authUsecase {
	return &authUsecase{
		repository: repository,
		dependency: dependency,
	}
}
