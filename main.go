package main

import (
	"context"
	"fmt"
	goredis "github.com/redis/go-redis/v9"
	"math/rand"
)

var ctx context.Context
var redisClient *goredis.Client

func main() {
	fmt.Println("Hello World!")
	ctx = context.Background()
	initializeRedis()

	setRandomKeys()
	printRandomKeys()
}

func initializeRedis() {
	redisClient = goredis.NewClient(&goredis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if redisClient == nil {
		panic("Failed to connect to redis")
	}
}

func setRandomKeys() {

	for i := 0; i < 10; i++ {
		randN := rand.Intn(10)
		redisClient.Set(ctx, fmt.Sprintf("key_%d", randN), fmt.Sprintf("value_%d", randN), 0)
	}
}

func printRandomKeys() {
	for i := 0; i < 10; i++ {
		randN := rand.Intn(10)
		key := fmt.Sprintf("key_%d", randN)
		val := redisClient.Get(ctx, key)
		if val != nil {
			fmt.Println(fmt.Sprintf("found value: %v", val))
		} else {
			fmt.Println(fmt.Sprintf("not found value for key: %v", key))
		}
	}
}
