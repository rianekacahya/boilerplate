package bootstrap

import (
	"context"
	"github.com/rianekacahya/boilerplate/domain/bootstrap"
	"github.com/rianekacahya/boilerplate/pkg/goconf"
	"github.com/rianekacahya/boilerplate/pkg/gologger"
	"github.com/spf13/cobra"
	"log"

	// repository
	aur "github.com/rianekacahya/boilerplate/internal/auth/repository"
	our "github.com/rianekacahya/boilerplate/internal/oauth/repository"

	// usecase
	"github.com/rianekacahya/boilerplate/internal/auth"
	"github.com/rianekacahya/boilerplate/internal/oauth"
)

var (
	dependency bootstrap.Dependency
	repository bootstrap.Repository
	usecase    bootstrap.Usecase
)

const (
	cfgPath = "files/config"
)

func Dependencies() {
	// init dependency
	config := goconf.New(cfgPath)
	dependency = bootstrap.Dependency{
		Cfg:    config,
		Logger: gologger.New(),
		Dbr:    newpostgresread(config),
		Dbw:    newpostgreswrite(config),
		Rdb:    newredis(context.Background(), config),
		Jwt:    newjwt(config, newrsa(config)),
	}

	// init repository
	repository = bootstrap.Repository{
		Auth:  aur.NewAuthRepository(dependency),
		Oauth: our.NewOauthRepository(dependency),
	}

	// init usecase
	usecase = bootstrap.Usecase{
		Auth:  auth.NewAuthUsecase(repository, dependency),
		Oauth: oauth.NewOauthUsecase(repository, dependency),
	}
}

func Execute() {
	var command = new(cobra.Command)

	command.AddCommand(
		rest,
		migrationup,
		migrationdown,
	)

	if err := command.Execute(); err != nil {
		log.Fatalf("got an error while initialize, error: %s", err)
	}
}
