package utils

import "flag"

var configuration Configuration

type Configuration struct {
	AppPort       *int
	RedisHost     *string
	RedisPort     *int
	RedisDb       *int64
	RedisPassword *string
}

func Init() {
	configuration = Configuration{
		AppPort:       flag.Int("app_port", 8181, "App Listening port"),
		RedisHost:     flag.String("redis_host", "localhost", "Redis Host"),
		RedisPort:     flag.Int("redis_port", 6379, "Redis Port"),
		RedisDb:       flag.Int64("redis_db", 0, "Redis DB"),
		RedisPassword: flag.String("redis_password", "", "Redis Password"),
	}
}

func GetConfig() Configuration {
	return configuration
}
