package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"

	"redis-test/internal/client"
)

var cl *redis.Client

func putQueue(key, value string) {
	id, err := cl.XAdd(&redis.XAddArgs{
		Stream: "stream",
		Values: map[string]interface{}{key: value},
	}).Result()

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println(id)
}

func putQueueWithID(queueID, key, value string) {
	id, err := cl.XAdd(&redis.XAddArgs{
		Stream: "stream",
		ID:     queueID,
		Values: map[string]interface{}{key: value},
	}).Result()

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println(id)
}

func main() {
	dbName := 0

	var err error
	cl, err = client.NewRedisClient(dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer cl.Close()

	key := "test"
	val := "value"
	putQueue(key, val)
	putQueue(key, val)
}
