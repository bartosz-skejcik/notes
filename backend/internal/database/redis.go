package database

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func CreateRedisConnection() (*redis.Client, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})

	return rdb, nil
}