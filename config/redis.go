package config

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient armazena a conexão com Redis
type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

// NewRedisClient cria e inicializa a conexão com Redis
func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Adicione senha se necessário
		DB:       0,
	})

	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("Erro ao conectar ao Redis: %v", err)
	}

	return &RedisClient{Client: client, Ctx: ctx}
}

// Set armazena um valor no cache Redis
func (r *RedisClient) Set(key string, value string, expiration time.Duration) error {
	return r.Client.Set(r.Ctx, key, value, expiration).Err()
}

// Get busca um valor no cache Redis
func (r *RedisClient) Get(key string) (string, error) {
	return r.Client.Get(r.Ctx, key).Result()
}
