package Utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

// IsRateLimited this function you enter the ip and the amount you want to rate limit
// then the expiration for the time you want your ip to expire and it checks the and blocks ips that surpass
// the rate limit
func IsRateLimited(ip string, redisClient *redis.Client, rateLimit int64, expiration time.Duration) error {
	commandable := redisClient.SetNX(redisClient.Context(), ip, 0, expiration)
	if commandable.Err() != nil {
		return fmt.Errorf("setting content in redis : ratelimit %w", commandable.Err())
	}
	CommandInt := redisClient.Incr(redisClient.Context(), ip)
	log.Println(commandable)
	if CommandInt.Err() != nil {
		return fmt.Errorf("setting content in redis : ratelimit: %w", CommandInt.Err())
	}
	NumberOfRequests, err := CommandInt.Result()
	if err != nil {
		return fmt.Errorf("failed to get number of request: %w", err)
	}
	if rateLimit < NumberOfRequests {
		return fmt.Errorf("to many requests : ratelimit")
	}
	return nil
}
