// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"context"
	"log"
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// HelloPubSub consumes a Pub/Sub message.
func FlushPubSub(ctx context.Context, m PubSubMessage) error {
	log.Println(string(m.Data))
	redisHost := os.Getenv("REDISHOST")
	redisPort := os.Getenv("REDISPORT")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	const maxConnections = 3
	redisPool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", redisAddr)
	}, maxConnections)

	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("FLUSHALL")
	if err != nil {
		fmt.Println("error flushing all keys", err)
		return err
	}
	fmt.Println("Flushed all keys")
	return nil
}
