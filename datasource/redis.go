package datasource

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnectRedis() {
	Redis = redis.NewClient(&redis.Options{})
	err := Redis.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal().Err(err).Msg("Couldn't connect to Redis")
	}
	log.Info().Msg("Redis Connected")
}
