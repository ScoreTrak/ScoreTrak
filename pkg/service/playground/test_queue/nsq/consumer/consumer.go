package consumer

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	//"sync"
	//"time"
)

//var wg sync.WaitGroup

type Check_details struct { //MAKE SURE TO USE CAPITAL LETTERS OR VARIABLES WILL NOT BE EXPORTED
	Round_Number    int
	Values_to_check []int
}

// HandleMessage implements the Handler interface.
func HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil
	}

	buf := bytes.NewBuffer(m.Body)
	var check_details_instance Check_details
	if err := gob.NewDecoder(buf).Decode(&check_details_instance); err != nil {
		panic(check_details_instance)
	}
	fmt.Println(check_details_instance)
	//time.Sleep(1*time.Second)

	//defer wg.Done()
	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	//wg.Done()

	return nil
}

func main() {

	//wg.Add(2)
	// Instantiate a consumer that will subscribe to the provided channel.
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic", "channel", config)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Stop()
	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.ChangeMaxInFlight(2)
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(HandleMessage), 2)

	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err = consumer.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Stop()
	select {}
	//wg.Wait()
	// Gracefully stop the consumer.

}
