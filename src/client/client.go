package client

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
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

	for {
		prod, err := client.Get(ctx, "product").Result()
		if err != nil {
			log.Println(err)
		} else {
			log.Println(prod)
		}

		time.Sleep(1 * time.Second)
	}

}
