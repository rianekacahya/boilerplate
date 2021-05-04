package goredis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func New(ctx context.Context, address, password string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
