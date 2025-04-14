package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStore(c *redis.Client, ctx context.Context) *RedisStore {
	return &RedisStore{client: c, ctx: ctx}
}

func (r *RedisStore) Save(short, original string) error {
	return r.client.Set(r.ctx, short, original, 0).Err()
}

func (r *RedisStore) Get(short string) (string, error) {
	return r.client.Get(r.ctx, short).Result()
}
