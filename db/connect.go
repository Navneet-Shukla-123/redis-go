package db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	DB *redis.Client
}

func ConnectToRedis() (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	log.Println(pong, "Successfully connected to Redis")
	return &Redis{
		DB: rdb,
	}, nil
}

type RedisRepo interface {
	SetKey(ctx context.Context, key, value string) error
	GetKey(ctx context.Context, key string) (string, error)
	ListPush(ctx context.Context, side bool, key, value string) error
	ListPop(ctx context.Context, side bool, key string) (string, error)
}
