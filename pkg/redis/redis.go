package redis

import (
	"fmt"
	"os"
	"sync"

	"github.com/go-redis/redis"
)

var (
	client   *redis.Client
	openOnce sync.Once
)

func DB() *redis.Client {
	openOnce.Do(func() {
		fmt.Println("Testing Golang Redis")
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		pong, err := client.Ping().Result()
		if err != nil {
			fmt.Println("can not open DB connection")
			os.Exit(1)
		}
		fmt.Println(pong, err)
	})

	return client
}
