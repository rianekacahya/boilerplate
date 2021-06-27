package oauth

import (
	"context"
	"github.com/google/uuid"
	"github.com/rianekacahya/boilerplate/domain/bootstrap"
	"github.com/rianekacahya/boilerplate/domain/entity"
	"github.com/rianekacahya/boilerplate/pkg/argon2"
	"github.com/rianekacahya/boilerplate/pkg/goerror"
	"github.com/rianekacahya/boilerplate/pkg/token"
)

type oauthUsecase struct {
	repository bootstrap.Repository
	dependency bootstrap.Dependency
}

func NewOauthUsecase(repository bootstrap.Repository, dependency bootstrap.Dependency) *oauthUsecase {
	return &oauthUsecase{
		repository: repository,
		dependency: dependency,
	}
}

func (us *oauthUsecase) Token(ctx context.Context, req *entity.RequestToken) (*entity.Token, error) {
	// get detail client
	client, err := us.repository.Oauth.GetClientByClientID(ctx, req.ClientID)
	if err != nil {
		return nil, err
	}

	// check client secret correctness
	compare, _ := argon2.Decode([]byte(client.ClientSecret))
	if ok, _ := compare.Verify([]byte(req.ClientSecret)); !ok {
		return nil, goerror.New(goerror.ErrCodeUnauthorized, "client id and client secret not match")
	}

	// generate token
	jwt, err := us.dependency.Jwt.GenerateToken(client.ClientID, uuid.New(), token.ScopeBasic)
	if err != nil {
		return nil, goerror.Wrap(err, goerror.ErrCodeUnexpected, "failed when generating jwt token")
	}

	return &entity.Token{
		TokenType:    jwt.TokenType,
		AccessToken:  jwt.AccessToken,
		ExpireIn:     jwt.AccessTokenExpireIn,
		RefreshToken: jwt.RefreshToken,
	}, nil
}

func (us *oauthUsecase) RefreshToken(ctx context.Context, req *entity.RequestToken) (*entity.Token, error) {
	var (
		sessionID = uuid.New()
		scope     = token.ScopeBasic
	)

	// decode refresh token
	claim, err := us.dependency.Jwt.Decode(req.RefreshToken)
	if err != nil {
		return nil, goerror.Wrap(err, goerror.ErrCodeUnexpected, "error when decoding token")
	}

	// validate refresh token
	if err := us.dependency.Jwt.Validate(claim); err != nil {
		return nil, goerror.Wrap(err, goerror.ErrCodeExpired, "error when validating token")
	}

	// get detail client
	client, err := us.repository.Oauth.GetClientByClientID(ctx, req.ClientID)
	if err != nil {
		return nil, err
	}

	// check client secret correctness
	compare, _ := argon2.Decode([]byte(client.ClientSecret))
	if ok, _ := compare.Verify([]byte(req.ClientSecret)); !ok {
		return nil, goerror.New(goerror.ErrCodeUnauthorized, "client id and client secret not match")
	}

	// check session exist in redis
	exist, err := us.repository.Oauth.CheckSessionExist(ctx, claim.Subject)
	if err != nil {
		return nil, err
	}

	// if session exist
	if exist {
		// set token scope
		scope = claim.Scope

		// parse value session ID
		sessionID, err = uuid.Parse(claim.Subject)
		if err != nil {
			return nil, goerror.Wrap(err, goerror.ErrCodeFormatting, "failed when parse session id")
		}

		// get session
		session, err := us.repository.Oauth.GetSession(ctx, claim.Subject)
		if err != nil {
			return nil, err
		}

		// compare client ID between user login and request body
		if client.ClientID != session.ClientID {
			return nil, goerror.New(goerror.ErrCodeUnauthorized, "client id not match")
		}
	}

	// generate token
	jwt, err := us.dependency.Jwt.GenerateToken(client.ClientID, sessionID, scope)
	if err != nil {
		return nil, goerror.Wrap(err, goerror.ErrCodeUnexpected, "failed when generating jwt token")
	}

	return &entity.Token{
		TokenType:    jwt.TokenType,
		AccessToken:  jwt.AccessToken,
		ExpireIn:     jwt.AccessTokenExpireIn,
		RefreshToken: jwt.RefreshToken,
	}, nil
}
