package main

import (
	"fmt"
	"log"
	"microservice-case2/common"
	"runtime"

	"github.com/nats-io/go-nats"
	mgo "gopkg.in/mgo.v2"
)

type Person struct {
	Name string
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	// defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")

	nc, err := nats.Connect(common.NatsURL)
	if err != nil {
		log.Fatalf("Failed to start app. Error=%s", err.Error())
	}

	fmt.Println(`Connect to`, common.NatsURL)
	fmt.Println(`------------------------------`)

	nc.Subscribe(common.TopicCreatePerson, func(msg *nats.Msg) {
		fmt.Println(`Received a create person message, topic=`, common.TopicCreatePerson, ` name=`, string(msg.Data))

		err = c.Insert(&Person{string(msg.Data)})
		if err != nil {
			log.Fatal(err)
		}
	})

	err = nc.Flush()
	if err != nil {
		log.Fatalf(`Failed to flush. Error=%s`, err.Error())
	}

	runtime.Goexit()
}
