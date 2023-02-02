package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	*redis.Client
}

func New(addr string, passwd string, db int) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	return &Cache{client}
}

func (c *Cache) Ping(ctx context.Context) error {
	return c.Client.Ping(ctx).Err()
}
