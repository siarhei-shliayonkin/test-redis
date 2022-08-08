package main

import (
	"fmt"

	"github.com/siarhei-shliayonkin/test-redis/pkg/redis"
)

func main() {
	fmt.Println("Hello world!")
	testKeyValue()
}

func testKeyValue() {
	db := redis.DB()
	err := db.Set("key_1", "value_1", 0).Err()
	if err != nil {
		fmt.Printf("db.Set(): %v\n", err)
		return
	}

	val, err := db.Get("key_1").Result()
	if err != nil {
		fmt.Printf("db.Get(): %v\n", err)
		return
	}
	fmt.Printf("value: %v\n", val)
}
