package conf

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisClient struct {
}

var rdClient *redis.Client

// 数据持续时间
const (
	DEFAULT_DURATION = 30 * 24 * 60 * 60 * time.Second
)

func InitRedis() (*RedisClient, error) {
	rdClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.Addr"),
		Password: "",
		DB:       0,
	})

	_, err := rdClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisClient{}, nil
}

func (rc *RedisClient) Set(key string, val interface{}, args ...interface{}) error {
	d := DEFAULT_DURATION
	if len(args) > 0 {
		if v, ok := args[0].(time.Duration); ok {
			d = v
		}
	}
	return rdClient.Set(context.Background(), key, val, d).Err()
}
func (rc *RedisClient) Get(key string) (string, error) {
	return rdClient.Get(context.Background(), key).Result()
}
func (rc *RedisClient) Delete(keys ...string) error {
	return rdClient.Del(context.Background(), keys...).Err()
}

func (rc *RedisClient) TTL(ctx context.Context, key string) (time.Duration, error) {
	duration, err := rdClient.TTL(ctx, key).Result()
	return duration, err
}
