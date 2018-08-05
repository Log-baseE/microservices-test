package main

import (
	"log"
	"microservice-case2/client"
	"microservice-case2/common"
	"net/http"

	"github.com/go-redis/redis"
)

var natsClient client.NatsInterface
var redisClient *redis.Client

func main() {
	natsClient = client.NewNatsClient(common.NatsURL)

	http.HandleFunc(`/publish`, new(handler).publish)
	http.HandleFunc(`/publish-and-receive`, new(handler).publishAndReceive)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type handler struct{}

func (h *handler) publish(w http.ResponseWriter, r *http.Request) {
	name := `Manusia Jahanam`
	err := natsClient.Publish(common.TopicCreatePerson, name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error=" + err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Message has been submitted successfully: " + name))
	}
}

func (h *handler) publishAndReceive(w http.ResponseWriter, r *http.Request) {
	msg, err := natsClient.Request(common.TopicRequestReply, `Make me a user named`)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error=" + err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(msg.Data))
	}
}
