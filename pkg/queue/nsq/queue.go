package nsq

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup"
	"github.com/gofrs/uuid"
	"github.com/nsqio/go-nsq"
)

type WorkerQueue struct {
	config queueing.Config
}

var ErrWorkersFailed = errors.New("some workers failed to receive the checks. Make sure that is by design")

// Send sends scoring data to the NSQD nodes, and returns either a list of checks with a warning, or an error
func (n WorkerQueue) Send(scoringDataAggregate []*queueing.ScoringData) ([]*queueing.QCheck, error, error) {
	// change this just the checks channel
	returningTopicName, err := queueing.TopicFromServiceRound(scoringDataAggregate[0].RoundID)
	if err != nil {
		return nil, nil, err
	}
	producerConfig := nsq.NewConfig()
	nsqProducerConfig(producerConfig, n.config)
	producer, err := nsq.NewProducer(n.config.NSQ.ProducerNSQD, producerConfig)
	if err != nil {
		return nil, nil, err
	}
	producer.SetLoggerLevel(nsq.LogLevelWarning)
	defer producer.Stop()

	m := make(map[string][][]byte)
	for _, scoringData := range scoringDataAggregate {
		scoringData.Service.ReturningTopic = returningTopicName
		buf := &bytes.Buffer{}
		if err := gob.NewEncoder(buf).Encode(scoringData); err != nil {
			return nil, nil, err
		}
		m[scoringData.Service.Group] = append(m[scoringData.Service.Group], buf.Bytes())
	}

	for k, v := range m {
		err = producer.MultiPublish(k, v)
		if err != nil {
			return nil, nil, err
		}
	}
	consumerConfig := nsq.NewConfig()
	nsqConsumerConfig(consumerConfig, n.config)
	consumer, err := nsq.NewConsumer(returningTopicName, "worker", consumerConfig)
	if err != nil {
		return nil, nil, err
	}
	defer consumer.Stop()
	ret := make([]*queueing.QCheck, len(scoringDataAggregate))
	consumer.ChangeMaxInFlight(len(scoringDataAggregate))
	qCheckChannel := make(chan queueing.IndexedQueue, 1)
	consumer.SetLoggerLevel(nsq.LogLevelWarning)

	idIndexMap := make(map[uuid.UUID]int)

	for i, sd := range scoringDataAggregate {
		idIndexMap[sd.Service.ID] = i
	}

	consumer.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		buf := bytes.NewBuffer(m.Body)
		var qc queueing.QCheck
		if err := gob.NewDecoder(buf).Decode(&qc); err != nil {
			return err
		}
		if i, ok := idIndexMap[qc.Service.ID]; ok {
			qCheckChannel <- queueing.IndexedQueue{Q: &qc, I: i}
			return nil
		}
		return nil
	}))
	err = connectConsumer(consumer, n.config)
	if err != nil {
		return nil, nil, err
	}
	counter := len(scoringDataAggregate)
	for {
		select {
		case res := <-qCheckChannel:
			ret[res.I] = res.Q
			counter--
			if counter == 0 {
				return ret, nil, nil
			}
		case <-time.After(time.Until(scoringDataAggregate[0].Deadline)):
			if !n.config.NSQ.IgnoreAllScoresIfWorkerFails {
				return nil, nil, &queueing.RoundTookTooLongToExecuteError{Msg: "Round took too long to score. This might be due to many reasons like a worker going down, or the number of rounds being too big for workers to handle"}
			}
			return ret, ErrWorkersFailed, nil
		}
	}
}

var ErrUnknownPanic = errors.New("unknown panic")
var ErrPanic = errors.New("panic")

func (n WorkerQueue) Receive() {
	conf := nsq.NewConfig()
	nsqConsumerConfig(conf, n.config)
	consumer, err := nsq.NewConsumer(n.config.NSQ.Topic, "worker", conf)
	if err != nil {
		log.Panicf("Failed to initialize NSQ consumer. Error: %v", err)
	}
	consumer.SetLoggerLevel(nsq.LogLevelWarning)
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
		buf := bytes.NewBuffer(m.Body)
		var sd queueing.ScoringData
		defer func() {
			if x := recover(); x != nil {
				var err error
				switch x := x.(type) {
				case string:
					err = fmt.Errorf("%w: %s", ErrPanic, x)
				case error:
					err = x
				default:
					err = ErrUnknownPanic
				}
				qc := queueing.QCheck{Service: sd.Service, Passed: false, Log: "Encountered an unexpected error during the check.", Err: err.Error(), RoundID: sd.RoundID}
				n.Acknowledge(qc)
				return
			}
		}()
		if err := gob.NewDecoder(buf).Decode(&sd); err != nil {
			panic(err)
		}

		dsync := -time.Since(sd.MasterTime)
		if float64(time.Second*5) < math.Abs(float64(dsync)) {
			name, _ := os.Hostname()
			n.Acknowledge(queueing.QCheck{Service: sd.Service, Passed: false, Log: "Please provide the error to Black Team / Competition Administrator", Err: fmt.Sprintf("Worker with IP: %s, Hostname: %s is either out of sync, or worker received the message late", getOutboundIP(), name), RoundID: sd.RoundID})
			return nil
		}
		qc := queueing.CommonExecute(&sd, sd.Deadline.Add(-3*time.Second+dsync))
		n.Acknowledge(qc)
		return nil
	}), n.config.NSQ.ConcurrentHandlers)
	err = connectConsumer(consumer, n.config)
	if err != nil {
		panic(err)
	}
	select {}
}

// https://stackoverflow.com/a/37382208/9296389
func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr, _ := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func (n WorkerQueue) Ping(group *servicegroup.ServiceGroup) error {
	_, bErr, err := n.Send([]*queueing.ScoringData{
		{
			Service: queueing.QService{ID: uuid.Nil, Name: "PING", Group: group.Name}, MasterTime: time.Now(), Host: "localhost", Deadline: time.Now().Add(time.Second * 4), RoundID: 0, Properties: map[string]string{},
		},
	})
	if err != nil {
		return err
	}
	if bErr != nil {
		return bErr
	}
	return nil
}

func (n WorkerQueue) Acknowledge(qCheck queueing.QCheck) {
	confp := nsq.NewConfig()
	nsqProducerConfig(confp, n.config)
	producer, err := nsq.NewProducer(n.config.NSQ.ProducerNSQD, confp)
	if err != nil {
		panic(err)
	}
	producer.SetLoggerLevel(nsq.LogLevelWarning)
	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(&qCheck); err != nil {
		panic(err)
	}
	err = producer.Publish(qCheck.Service.ReturningTopic, buf.Bytes())
	if err != nil {
		panic(err)
	}
	producer.Stop()
}

func NewNSQWorkerQueue(config queueing.Config) (*WorkerQueue, error) {
	err := validateNSQConfig(config)
	if err != nil {
		return nil, err
	}
	return &WorkerQueue{config}, nil
}

//Todo: For Master-Worker Exchange Queue look into simplifying, and speeding up the process by utilizing single listening topic, and reusing the topic from round to round
