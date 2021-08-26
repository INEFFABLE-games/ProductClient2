package client

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func getRedis() *redis.Client {
	var (
		RedisAddres    = "redis-14436.c54.ap-northeast-1-2.ec2.cloud.redislabs.com:14436"
		RedisPassword  = "Drektarov3698!"
		ReddisUserName = "Admin"
	)

	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddres,
		Password: RedisPassword,
		Username: ReddisUserName,
		DB:       0,
	})

	res, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	return client
}

func Start(ctx context.Context) {

	client := getRedis()
	opt := "$"
	lastKey := 0

	for {
		//prod, err := client.Get(ctx, "product").Result()
		prod, err := client.XRead(ctx, &redis.XReadArgs{
			Streams: []string{"products", opt},
			Count:   0,
			Block:   2 * time.Second,
		}).Result()

		if err != nil {
			log.Println(err)
		} else {
			for _, val := range prod {
				for k, v := range val.Messages {
					if k >= lastKey {
						log.Printf("[%d]: %v \n", k, v)
						lastKey = k + 1
					}
				}

				fmt.Println("\n")
			}
		}

		opt = "1526999644174-" + strconv.Itoa(lastKey)
		//println(opt)
		time.Sleep(1 * time.Second)
	}

}
