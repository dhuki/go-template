package cache

import (
	"github.com/dhuki/go-template/internal/infra/configloader"
	"github.com/go-redis/redis"
)

func NewRedisClient(conf *configloader.RedisConfig) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Password: conf.Password,
		DB:       conf.DB,
	})

	if _, err := redisClient.Ping().Result(); err != nil {
		return nil, err
	}
	return redisClient, nil
}
