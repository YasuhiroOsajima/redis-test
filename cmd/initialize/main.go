package main

import (
	"fmt"
	"os"

	"redis-test/internal/client"
)

func main() {
	dbName := 0

	cl, err := client.NewRedisClient(dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer cl.Close()

	cl.FlushAll()
}
