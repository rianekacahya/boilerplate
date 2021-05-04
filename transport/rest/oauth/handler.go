package oauth

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"github.com/rianekacahya/boilerplate/domain/entity"
	"github.com/rianekacahya/boilerplate/domain/usecase"
	"github.com/rianekacahya/boilerplate/pkg/goerror"
	"github.com/rianekacahya/boilerplate/pkg/response"
)

type rest struct {
	oauthUsecase usecase.Oauth
}

func NewHandler(http chi.Router, oauthUsecase usecase.Oauth) {
	transport := rest{oauthUsecase}

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
		token, err = t.oauthUsecase.Token(ctx, req)
		if err != nil {
			response.Nay(w, r, err)
			return
		}
	}

	// call refresh token
	if req.GrantType == entity.RefreshToken {
		token, err = t.oauthUsecase.RefreshToken(ctx, req)
		if err != nil {
			response.Nay(w, r, err)
			return
		}
	}

	response.Yay(w, r, token)
}
