package main

import (
	"fmt"
	"log"
	"microservice-case2/common"
	"runtime"

	"github.com/nats-io/go-nats"
)

func main() {
	nc, err := nats.Connect(common.NatsURL)
	if err != nil {
		log.Fatalf("Failed to start app. Error=%s", err.Error())
	}

	fmt.Println(`Connect to`, common.NatsURL)
	fmt.Println(`------------------------------`)

	nc.Subscribe(common.TopicHelloWorld, func(msg *nats.Msg) {
		fmt.Println(`Receive a message, topic=`, common.TopicHelloWorld, ` message=`, string(msg.Data))
	})

	nc.Subscribe(common.TopicRequestReply, func(msg *nats.Msg) {
		fmt.Println(`Receive a message, topic=`, common.TopicHelloWorld, ` message=`, string(msg.Data))
		nc.Publish(msg.Reply, []byte("I can help c:"))
	})

	err = nc.Flush()
	if err != nil {
		log.Fatalf(`Failed to flush. Error=%s`, err.Error())
	}

	runtime.Goexit()
}
