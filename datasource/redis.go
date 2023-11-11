package datasource

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnectRedis() {
	Redis = redis.NewClient(&redis.Options{})
	err := Redis.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Redis Connected")
}
