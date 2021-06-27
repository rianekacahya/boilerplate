package bootstrap

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"github.com/rianekacahya/boilerplate/pkg/token"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rianekacahya/boilerplate/domain/repository"
	"github.com/rianekacahya/boilerplate/domain/usecase"
)

type Dependency struct {
	Logger *logrus.Logger
	Cfg    *viper.Viper
	Dbr    *sql.DB
	Dbw    *sql.DB
	Rdb    *redis.Client
	Jwt    *token.Token
}

type Repository struct {
	Auth  repository.Auth
	Oauth repository.Oauth
}

type Usecase struct {
	Auth  usecase.Auth
	Oauth usecase.Oauth
}
