package repository

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
)

type oauthRepository struct {
	dbwrite *sql.DB
	dbread  *sql.DB
	redis *redis.Client
}

func NewOauthRepository(dbwrite, dbread *sql.DB, redis *redis.Client) *oauthRepository {
	return &oauthRepository{dbwrite, dbread, redis}
}
