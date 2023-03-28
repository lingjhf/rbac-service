package cache

import "github.com/redis/go-redis/v9"

type Cache interface {
}

type RedisCache struct {
	Client *redis.Client
}
