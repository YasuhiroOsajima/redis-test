package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"

	"redis-test/internal/client"
)

var cl *redis.Client

// Return value is like [ db0:keys=1,expires=0,avg_ttl=0 ]
func getDbList() ([]string, error) {
	res := cl.Info()
	val, err := res.Result()
	if err != nil {
		return nil, err
	}

	slice := strings.Split(val, "Keyspace")
	dbListString := slice[len(slice)-1]
	dbListSlice := strings.Split(dbListString, "\r\n")
	return dbListSlice, nil
}

// Get all queue
// Return is like [{stream [{1599030184079-0 map[test1-0:value1-0]}]}]
func getQueueList(dbName string) ([]redis.XStream, error) {
	res, err := cl.XReadStreams("stream", dbName).Result()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Return is like {stream [{1599030184079-0 map[test1-0:value1-0]}]}
func getNextQueue(dbName string) (res redis.XStream, err error) {
	resList, err := cl.XRead(&redis.XReadArgs{
		Streams: []string{"stream", dbName},
		Count:   1,
		Block:   100 * time.Millisecond,
	}).Result()

	res = resList[0]
	if err != nil {
		return res, err
	}

	return res, nil
}

func deleteQeueue(targetQueueID string) {
	n, err := cl.XDel("stream", targetQueueID).Result()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(n)
}

func main() {
	dbName := 0
	dbNameStr := strconv.Itoa(dbName)

	var err error
	cl, err = client.NewRedisClient(dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer cl.Close()

	//getDbList()

	//reslist, _ := getQueueList(dbNameStr)
	//fmt.Println(reslist)

	res, _ := getNextQueue(dbNameStr)
	fmt.Println(res)

	oldestID := res.Messages[0].ID
	deleteQeueue(oldestID)

	res, _ = getNextQueue(dbNameStr)
	fmt.Println(res)
}
