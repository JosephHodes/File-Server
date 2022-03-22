package main

import (
	utils "./Server/Utils"
	"github.com/Nerzal/gocloak/v10"
	"github.com/go-redis/redis"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

var (
	RedisClient *redis.Client
	KeyCloakClient *gocloak.GoCloak
)

func InstantiateServices() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	KeyCloakClient = gocloak.NewClient("https://mycool.keycloak.instance")

	return nil
}

func main() {
	StatusCommand:=RedisClient.Set(RedisClient.Context(),"hi","hi",time.Minute)
	if StatusCommand != nil {
		log.Println(StatusCommand)
	}
	err := InstantiateServices()
	if err != nil {
		log.Println(err)
	}
	mode := strings.ToUpper(os.Getenv("mode"))
	switch mode {
	case "TEST":
		log.Println("Func tests succeeded ")
	}
	log.Fatalln("fatal error couldn't get mode it was defined as: " + mode)
}
