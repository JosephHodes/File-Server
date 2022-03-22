package Utils

import (
	"github.com/go-redis/redis"
	"testing"
)

func FuncTest(t *testing.T){
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
 ra\\
}