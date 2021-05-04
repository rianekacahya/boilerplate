package auth

import "github.com/rianekacahya/boilerplate/domain/repository"

type authUsecase struct {
	authRepository repository.Auth
	dependency     dependency
}

type dependency struct{}

func NewAuthUsecase(authRepository repository.Auth) *authUsecase {
	return &authUsecase{
		authRepository: authRepository,
	}
}
