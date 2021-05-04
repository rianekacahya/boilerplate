package auth

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"github.com/rianekacahya/boilerplate/domain/usecase"
	"github.com/rianekacahya/boilerplate/pkg/response"
)

type rest struct {
	authUsecase usecase.Auth
}

func NewHandler(http chi.Router, authUsecase usecase.Auth) {
	transport := rest{authUsecase}

	http.Route("/auth", func(r chi.Router) {
		r.Get("/login", transport.login)
	})
}

func (t *rest) login(w http.ResponseWriter, r *http.Request) {
	response.Yay(w, r, "Hello World !")
}