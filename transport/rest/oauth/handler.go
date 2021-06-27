package oauth

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rianekacahya/boilerplate/domain/bootstrap"
	"github.com/rianekacahya/boilerplate/domain/entity"
	"github.com/rianekacahya/boilerplate/pkg/goerror"
	"github.com/rianekacahya/boilerplate/pkg/response"
	"net/http"
)

type rest struct {
	usecase bootstrap.Usecase
}

func NewHandler(http chi.Router, usecase bootstrap.Usecase) {
	transport := rest{usecase}

	http.Route("/oauth2", func(r chi.Router) {
		r.Post("/token", transport.token)
	})
}

func (t *rest) token(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		ctx   = r.Context()
		req   = new(entity.RequestToken)
		token = new(entity.Token)
	)

	if err := render.Bind(r, req); err != nil {
		response.Nay(w, r, goerror.Wrap(err, goerror.ErrCodeFormatting, "please check your request body"))
		return
	}

	// generate access token
	if req.GrantType == entity.AccessToken {
		token, err = t.usecase.Oauth.Token(ctx, req)
		if err != nil {
			response.Nay(w, r, err)
			return
		}
	}

	// call refresh token
	if req.GrantType == entity.RefreshToken {
		token, err = t.usecase.Oauth.RefreshToken(ctx, req)
		if err != nil {
			response.Nay(w, r, err)
			return
		}
	}

	response.Yay(w, r, token)
}
