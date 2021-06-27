package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/rianekacahya/boilerplate/domain/bootstrap"
	"github.com/rianekacahya/boilerplate/pkg/response"
	"net/http"
)

type rest struct {
	usecase bootstrap.Usecase
}

func NewHandler(http chi.Router, usecase bootstrap.Usecase) {
	transport := rest{usecase}

	http.Route("/auth", func(r chi.Router) {
		r.Get("/login", transport.login)
	})
}

func (t *rest) login(w http.ResponseWriter, r *http.Request) {
	response.Yay(w, r, "Hello World !")
}
