package client

import "github.com/go-redis/redis"

func NewRedisClient(dbName int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		//Addr:     "redis-14450.c1.asia-northeast1-1.gce.cloud.redislabs.com:14450",
		//Password: "37uaACndCvuQ1heADnHkishnAhMmosWq", // no password set
		Addr:     "localhost:6379",
		Password: "",
		DB:       dbName, // use default DB
	})

	_, err := client.Ping().Result()
	return client, err
}
