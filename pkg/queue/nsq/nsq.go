package nsq

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/gofrs/uuid"
	"github.com/nsqio/go-nsq"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type WorkerQueue struct {
	config queueing.Config
}

func (n WorkerQueue) Send(sds []*queueing.ScoringData) (ret []*queueing.QCheck, bErr error, tErr error) {
	addresses := generateNSQLookupdAddresses(n.config.NSQ.NSQLookupd.Hosts, n.config.NSQ.NSQLookupd.Port)
	returningTopicName := queueing.TopicFromServiceRound(sds[0].RoundID)
	//bErr, tErr := n.TopicAbsent(returningTopicName, addresses)
	//if tErr != nil {
	//	return nil, bErr, tErr
	//}
	confp := nsq.NewConfig()
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%s", n.config.NSQ.NSQD.Host, n.config.NSQ.NSQD.Port), confp)
	if err != nil {
		return nil, nil, err
	}
	defer producer.Stop()

	m := make(map[string][][]byte)
	for _, sd := range sds {
		sd.Service.ReturningTopic = returningTopicName
		buf := &bytes.Buffer{}
		if err := gob.NewEncoder(buf).Encode(sd); err != nil {
			return nil, nil, err
		}
		if _, ok := m[sd.Service.Group]; ok {
			m[sd.Service.Group] = append(m[sd.Service.Group], buf.Bytes())
		} else {
			m[sd.Service.Group] = [][]byte{buf.Bytes()}
		}
	}

	for k, v := range m {
		err = producer.MultiPublish(k, v)
		if err != nil {
			return nil, nil, err
		}
	}

	defer func(returningTopicName string, addresses []string) {
		go n.DeleteTopic(returningTopicName, addresses)
	}(returningTopicName, addresses)
	confc := nsq.NewConfig()
	confc.LookupdPollInterval = time.Second * 1
	consumer, err := nsq.NewConsumer(returningTopicName, "worker", confc)
	if err != nil {
		return nil, bErr, err
	}
	defer consumer.Stop()
	ret = make([]*queueing.QCheck, len(sds))
	consumer.ChangeMaxInFlight(len(sds))
	cq := make(chan queueing.IndexedQueue, 1)
	consumer.SetLoggerLevel(nsq.LogLevelError)
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
		buf := bytes.NewBuffer(m.Body)
		var qc queueing.QCheck
		if err := gob.NewDecoder(buf).Decode(&qc); err != nil {
			return err
		}
		for i, sd := range sds {
			if sd.Service.ID == qc.Service.ID {
				cq <- queueing.IndexedQueue{Q: &qc, I: i}
				break
			}
		}
		return nil
	}), len(sds))
	err = consumer.ConnectToNSQLookupds(addresses)
	if err != nil {
		return nil, bErr, err
	}
	counter := len(sds)
	for {
		select {
		case res := <-cq:
			ret[res.I] = res.Q
			counter--
			if counter == 0 {
				return ret, bErr, nil
			}
		case <-time.After(time.Until(sds[0].Deadline)):
			if !n.config.NSQ.IgnoreAllScoresIfWorkerFails {
				return nil, bErr, &queueing.RoundTookTooLongToExecute{Msg: "Round took too long to score. This might be due to many reasons like a worker going down, or the number of rounds being too big for workers to handle"}
			} else {
				return ret, errors.New("some workers failed to receive the checks. Make sure that is by design"), nil
			}

		}
	}
}

func (n WorkerQueue) Receive() {
	conf := nsq.NewConfig()
	conf.LookupdPollInterval = time.Second * 2
	conf.MaxInFlight = n.config.NSQ.MaxInFlight
	consumer, err := nsq.NewConsumer(n.config.NSQ.Topic, "worker", conf)
	if err != nil {
		log.Fatalf("Failed to initialize NSQ consumer. Error: %v", err)
	}
	consumer.SetLoggerLevel(nsq.LogLevelError)
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
		buf := bytes.NewBuffer(m.Body)
		var sd queueing.ScoringData
		defer func() {
			if x := recover(); x != nil {
				var err error
				switch x := x.(type) {
				case string:
					err = errors.New(x)
				case error:
					err = x
				default:
					err = errors.New("unknown panic")
				}
				qc := queueing.QCheck{Service: sd.Service, Passed: false, Log: "Encountered an unexpected error during the check. This is most likely a bug", Err: err.Error(), RoundID: sd.RoundID}
				n.Acknowledge(qc)
				return
			}
		}()
		if err := gob.NewDecoder(buf).Decode(&sd); err != nil {
			panic(err)
		}
		qc := queueing.CommonExecute(&sd, sd.Deadline.Add(-3*time.Second))
		n.Acknowledge(qc)
		return nil

	}), n.config.NSQ.ConcurrentHandlers)
	addresses := generateNSQLookupdAddresses(n.config.NSQ.NSQLookupd.Hosts, n.config.NSQ.NSQLookupd.Port)
	err = consumer.ConnectToNSQLookupds(addresses)
	if err != nil {
		panic(err)
	}
	select {}
}

func (n WorkerQueue) Ping(group *service_group.ServiceGroup) error {
	_, bErr, err := n.Send([]*queueing.ScoringData{
		{
			Service: queueing.QService{ID: uuid.Nil, Name: "PING", Group: group.Name}, Host: "localhost", Deadline: time.Now().Add(time.Second * 4), RoundID: 0, Properties: map[string]string{},
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

func generateNSQLookupdAddresses(hostNames []string, port string) []string {
	var addresses []string
	for _, h := range hostNames {
		addresses = append(addresses, fmt.Sprintf("%s:%s", h, port))
	}
	return addresses
}

func (n WorkerQueue) Acknowledge(q queueing.QCheck) {
	confp := nsq.NewConfig()
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%s", n.config.NSQ.NSQD.Host, n.config.NSQ.NSQD.Port), confp)
	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(&q); err != nil {
		panic(err)
	}
	err = producer.Publish(q.Service.ReturningTopic, buf.Bytes())
	if err != nil {
		panic(err)
	}
	producer.Stop()
}

func (n WorkerQueue) DeleteTopic(topic string, nsqAddresses []string) { //This makes NSQ node unusable for a while
	time.Sleep(time.Second * 5)
	for _, a := range nsqAddresses {
		client := http.Client{
			Timeout: time.Second / 2,
		}
		resp, err := client.Post(fmt.Sprintf("http://%s/topic/delete?topic=%s", a, topic), "", nil)
		if err == nil {
			resp.Body.Close()
			return
		}
	}
}

//type topics struct {
//	Topics []string `json:"topics"`
//}

//func (n WorkerQueue) TopicAbsent(topic string, nsqAddresses []string) (bErr error, tErr error) {
//	var err error
//	for _, a := range nsqAddresses {
//		client := http.Client{
//			Timeout: time.Second / 2,
//		}
//		resp, err2 := client.Get(fmt.Sprintf("http://%s/topics", a))
//		if err2 != nil {
//			err = err2
//			continue
//		}
//		topics := topics{}
//		errd := json.NewDecoder(resp.Body).Decode(&topics)
//		if errd != nil {
//			return err, errd
//		}
//		for _, val := range topics.Topics {
//			if val == topic {
//				return err, fmt.Errorf("NSQ Topic with the same name as %s exists. Round will be terminated. Please firt clean NSQ queues", topic)
//			}
//		}
//		return err, nil
//		resp.Body.Close()
//	}
//	return err, errors.New("no NSQLookupd instances answered the request")
//}
func NewNSQWorkerQueue(config queueing.Config) (*WorkerQueue, error) {
	return &WorkerQueue{config}, nil
}

type PubSub struct {
	config queueing.Config
}

func (p PubSub) NotifyTopic(topic string) {
	confp := nsq.NewConfig()
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%s", p.config.NSQ.NSQD.Host, p.config.NSQ.NSQD.Port), confp)
	if err != nil {
		log.Fatalf("Unable to initialize producer to notify masters using queue. Ensure that the queue is reachable from master. Error Details: %v", err)
	}
	err = producer.Publish(topic, make([]byte, 1))
	if err != nil {
		log.Fatalf("Unable to publish to topic to notify masters. Ensure that the queue is reachable from master. Error Details: %v", err)
	}
	producer.Stop()
}

func (p PubSub) ReceiveUpdateFromTopic(topic string) <-chan struct{} {
	n := make(chan struct{})
	go func() {
		conf := nsq.NewConfig()
		conf.LookupdPollInterval = time.Second * 2
		consumer, err := nsq.NewConsumer(topic, "master_"+strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int()), conf)
		if err != nil {
			log.Fatalf("Unable to initualize consumer for topic: %s. Error Details: %v", topic, err)
		}
		consumer.AddHandler(
			nsq.HandlerFunc(func(m *nsq.Message) error {
				n <- struct{}{}
				return nil
			}))
		addresses := generateNSQLookupdAddresses(p.config.NSQ.NSQLookupd.Hosts, p.config.NSQ.NSQLookupd.Port)
		err = consumer.ConnectToNSQLookupds(addresses)
		if err != nil {
			log.Fatalf("Unable to connect to NSQLookupd instances")
		}
		select {}
	}()
	return n
}

func NewNSQPubSub(config queueing.Config) (*PubSub, error) {
	return &PubSub{config}, nil
}

//Todo: For Master-Worker Exchange Queue look into simplifying, and speeding up the process by utilizing single listening topic, and reusing the topic from round to round
