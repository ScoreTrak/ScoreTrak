package main

import (
	"bytes"
	"encoding/gob"
	"github.com/nsqio/go-nsq"
	"log"
)

type Check_details struct {
	Round_Number    int
	Values_to_check []int
}

func main() {
	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	test_struct := Check_details{5, []int{1, 2}}
	println(test_struct.Round_Number)
	buf := &bytes.Buffer{}

	if err := gob.NewEncoder(buf).Encode(test_struct); err != nil {
		panic(err)
	}

	topicName := "topic"

	// Synchronously publish a single message to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	err = producer.Publish(topicName, buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	// Gracefully stop the producer.

	producer.Stop()

}
