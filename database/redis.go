package database

import (
	"context"
	"fmt"
	"rbac-service/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisWithConfig(config *config.Config) Database[*redis.Client] {
	return &Redis{
		Host:     config.CACHE_HOST,
		Port:     config.CACHE_PORT,
		User:     config.CACHE_USER,
		Password: config.CACHE_PASSWD,
	}
}

type Redis struct {
	Host     string
	Port     uint
	User     string
	Password string
	DB       int
}

func (r *Redis) Connect() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Username: r.User,
		Password: r.Password,
		DB:       r.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
