package bootstrap

import (
	"context"
	"crypto/rsa"
	"database/sql"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"github.com/rianekacahya/boilerplate/pkg/gopostgres"
	"github.com/rianekacahya/boilerplate/pkg/goredis"
	"github.com/rianekacahya/boilerplate/pkg/helper"
	"github.com/rianekacahya/boilerplate/pkg/token"
)

func newpostgreswrite(config *viper.Viper) *sql.DB {
	db, err := gopostgres.New(
		config.GetString("postgres.write.dsn"),
		config.GetInt("postgres.write.max_open"),
		config.GetInt("postgres.write.max_idle"),
		config.GetInt("postgres.write.timeout"),
	)

	if err != nil {
		log.Fatalf("got an error while connecting database write server, error: %s", err)
	}

	return db
}

func newpostgresread(config *viper.Viper) *sql.DB {
	db, err := gopostgres.New(
		config.GetString("postgres.read.dsn"),
		config.GetInt("postgres.read.max_open"),
		config.GetInt("postgres.read.max_idle"),
		config.GetInt("postgres.read.timeout"),
	)

	if err != nil {
		log.Fatalf("got an error while connecting database read server, error: %s", err)
	}

	return db
}

func newredis(ctx context.Context, config *viper.Viper) *redis.Client {
	client, err := goredis.New(
		ctx,
		config.GetString("redis.master.host"),
		config.GetString("redis.master.password"),
	)

	if err != nil {
		log.Fatalf("got an error while connecting to redis server, error: %s", err)
	}

	return client
}

func newjwt(config *viper.Viper, rsakey *rsa.PrivateKey) *token.Token {
	t, err := token.NewToken(
		rsakey,
		config.GetString("jwt.issuer"),
		config.GetDuration("jwt.access_token_expire_in"),
		config.GetDuration("jwt.refresh_token_expire_in"),
		config.GetDuration("jwt.leeway"),
	)

	if err != nil {
		log.Fatalf("got an error while initialize JWT, error: %s", err)
	}

	return t
}

func newrsa(config *viper.Viper) *rsa.PrivateKey {
	r, err := helper.ReadFileRSA(config.GetString("jwt.key"))
	if err != nil {
		log.Fatalf("got an error while reading rsa file, error: %s", err)
	}

	return r
}
