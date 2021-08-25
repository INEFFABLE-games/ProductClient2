package repository

import (
	"github.com/go-redis/redis/v8"
)

type ClientRedisRepository struct {
	db *redis.Client
}

func NewClientRedisRepository(db *redis.Client) ClientRedisRepository {
	return ClientRedisRepository{db: db}
}
