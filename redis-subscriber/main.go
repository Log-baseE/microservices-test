package main

import (
	"fmt"
	"log"
	"microservice-case2/common"
	"runtime"

	"github.com/go-redis/redis"
	"github.com/nats-io/go-nats"
)

var redisClient *redis.Client

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	nc, err := nats.Connect(common.NatsURL)
	if err != nil {
		log.Fatalf("Failed to start app. Error=%s", err.Error())
	}

	fmt.Println(`Connect to`, common.NatsURL)
	fmt.Println(`------------------------------`)

	nc.Subscribe(common.TopicCreatePerson, func(msg *nats.Msg) {
		fmt.Println(`Received a create person message, topic=`, common.TopicCreatePerson, ` name=`, string(msg.Data))

		_, err := redisClient.Set(`name`, string(msg.Data), 0).Result()
		if err != nil {
			panic(err)
		}
	})

	err = nc.Flush()
	if err != nil {
		log.Fatalf(`Failed to flush. Error=%s`, err.Error())
	}

	runtime.Goexit()
}
