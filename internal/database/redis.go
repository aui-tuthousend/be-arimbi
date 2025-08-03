package database

import (
	"context"
    "os"
    "log"

	"github.com/redis/go-redis/v9"
)

func InitRedisClient() (*redis.Client, context.Context) {
    dsn := os.Getenv("REDIS_URL")
    if dsn == "" {
        log.Fatal("Redis client is not set")
        return nil, nil
    }
    rdb := redis.NewClient(&redis.Options{
        Addr:     dsn,
        Password: "",
        DB:       0,
    })
    return rdb, context.Background()
}
