package bootstrap

import (
	"context"
	"crypto/rsa"
	"database/sql"
	"github.com/go-redis/redis/v8"
	"github.com/rianekacahya/boilerplate/pkg/goconf"
	"github.com/rianekacahya/boilerplate/pkg/gologger"
	"github.com/rianekacahya/boilerplate/pkg/token"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var (
	logger *logrus.Logger
	cfg    *viper.Viper
	dbr    *sql.DB
	dbw    *sql.DB
	rdb    *redis.Client
	rsakey *rsa.PrivateKey
	jwt    *token.Token
)

const (
	cfgPath = "files/config"
)

func Dependencies() {
	logger = gologger.New()
	cfg = goconf.New(cfgPath)
	dbr = newpostgresread(cfg)
	dbw = newpostgreswrite(cfg)
	rdb = newredis(context.Background(), cfg)
	rsakey = newrsa(cfg)
	jwt = newjwt(cfg, rsakey)
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
