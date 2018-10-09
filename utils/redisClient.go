package utils

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client redis.Client
var isClientInit bool = false

func newClient() {
	host := fmt.Sprintf("%s:%d", *GetConfig().RedisHost, *GetConfig().RedisPort)
	client = *redis.NewClient(&redis.Options{
		Addr:     host,
		Password: *GetConfig().RedisPassword,
		DB:       *GetConfig().RedisDb,
	})
}

func GetClient() redis.Client {
	if !isClientInit {
		newClient()
		isClientInit = true
	}
	return client
}
