package main

import (
	utils "./Server/Utils"
	"./Tests/FuncTests"
	"github.com/go-redis/redis"
	"github.com/go-redsync/redsync/redis/redigo"
	redis3 "github.com/gomodule/redigo/redis"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

var pools *redis3.Pool

func main() {
	//initializing redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pools = &redis3.Pool{
		Dial: func() (redis3.Conn, error) { return redis3.Dial("tcp", "localhost:6379") },
	}
	pool := redigo.NewPool(pools)
	log.Println(reflect.TypeOf(pool))

	errr := utils.IsRateLimited("111.222.222.333", redisClient, 2, time.Second*10, pool)
	errr = utils.IsRateLimited("111.222.222.333", redisClient, 2, time.Second*10., pool)
	errr = utils.IsRateLimited("111.222.222.333", redisClient, 2, time.Second*10, pool)
	if errr != nil {
		log.Println(errr)
	}
	hi := redisClient.Get(redisClient.Context(), "111.222.222.333")
	log.Println(hi)
	mode := strings.ToUpper(os.Getenv("mode"))
	switch mode {
	case "TEST":
		err := FuncTests.FuncTest()
		log.Println("Func tests succeeded ")
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Fatalln("fatal error couldn't get mode it was defined as: " + mode)
}
