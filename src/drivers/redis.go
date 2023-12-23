package drivers

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisDriver struct {
	client *redis.Client
}

func NewRedisDriver() *RedisDriver {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisDriver{client: client}
}

func (rd RedisDriver) Ping(ctx context.Context) (string, error) {
	result, err := rd.client.Ping(ctx).Result()

	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}
