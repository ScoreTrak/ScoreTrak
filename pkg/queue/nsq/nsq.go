package nsq

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/gofrs/uuid"
	"github.com/nsqio/go-nsq"
	"net/http"
	"time"
)

type NSQ struct {
	l      logger.LogInfoFormat
	config queueing.Config
}

func (n NSQ) Send(sds []*queueing.ScoringData) ([]*queueing.QCheck, error, error) {
	addresses := n.GenerateNSQLookupdAddresses(n.config.NSQ.NSQLookupd.Hosts, n.config.NSQ.NSQLookupd.Port)
	returningTopicName := queueing.TopicFromServiceRound(&sds[0].Service, sds[0].RoundID)
	bErr, tErr := n.TopicAbsent(returningTopicName, addresses)
	if tErr != nil {
		return nil, bErr, tErr
	}
	confp := nsq.NewConfig()
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%s", n.config.NSQ.NSQD.Host, n.config.NSQ.NSQD.Port), confp)
	if err != nil {
		return nil, nil, err
	}
	defer producer.Stop()
	for _, sd := range sds {
		sd.Service.ReturningTopic = returningTopicName
		buf := &bytes.Buffer{}
		if err := gob.NewEncoder(buf).Encode(sd); err != nil {
			return nil, nil, err
		}
		err = producer.Publish(sd.Service.Group, buf.Bytes())
		if err != nil {
			return nil, nil, err
		}
	}
	defer func(returningTopicName string, addresses []string) {
		go n.DeleteTopic(returningTopicName, addresses)
	}(returningTopicName, addresses)
	confc := nsq.NewConfig()
	confc.LookupdPollInterval = time.Second * 1
	consumer, err := nsq.NewConsumer(returningTopicName, "channel", confc)
	if err != nil {
		return nil, bErr, err
	}
	defer consumer.Stop()
	ret := make([]*queueing.QCheck, len(sds))
	consumer.ChangeMaxInFlight(len(sds))
	cq := make(chan queueing.IndexedQueue, 1)
	consumer.SetLoggerLevel(nsq.LogLevelError)
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
		buf := bytes.NewBuffer(m.Body)
		var qc queueing.QCheck
		if err := gob.NewDecoder(buf).Decode(&qc); err != nil {
			n.l.Error(err)
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

func (n NSQ) Receive() {
	conf := nsq.NewConfig()
	conf.LookupdPollInterval = time.Second * 2
	conf.MaxInFlight = n.config.NSQ.MaxInFlight
	consumer, err := nsq.NewConsumer(n.config.NSQ.Topic, "channel", conf)
	if err != nil {
		panic(err)
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
				n.l.Error(err)
				return
			}
		}()
		if err := gob.NewDecoder(buf).Decode(&sd); err != nil {
			panic(err)
		}
		qc := queueing.CommonExecute(&sd, sd.Deadline.Add(-3*time.Second), n.l)
		n.Acknowledge(qc)
		return nil

	}), n.config.NSQ.ConcurrentHandlers)
	addresses := n.GenerateNSQLookupdAddresses(n.config.NSQ.NSQLookupd.Hosts, n.config.NSQ.NSQLookupd.Port)
	err = consumer.ConnectToNSQLookupds(addresses)
	if err != nil {
		panic(err)
	}
	select {}
}

func (n NSQ) Ping(group *service_group.ServiceGroup) error {
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

func (n NSQ) GenerateNSQLookupdAddresses(hostNames []string, port string) []string {
	var addresses []string
	for _, h := range hostNames {
		addresses = append(addresses, fmt.Sprintf("%s:%s", h, port))
	}
	return addresses
}

func (n NSQ) Acknowledge(q queueing.QCheck) {
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

func (n NSQ) DeleteTopic(topic string, nsqAddresses []string) { //THis make NSQ node unusable for a while
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
		n.l.Error(err)
	}
}

type topics struct {
	Topics []string `json:"topics"`
}

func (n NSQ) TopicAbsent(topic string, nsqAddresses []string) (bErr error, tErr error) {
	var err error
	for _, a := range nsqAddresses {
		client := http.Client{
			Timeout: time.Second / 2,
		}
		resp, err2 := client.Get(fmt.Sprintf("http://%s/topics", a))
		if err2 != nil {
			err = err2
			continue
		}
		topics := topics{}
		errd := json.NewDecoder(resp.Body).Decode(&topics)
		if errd != nil {
			return err, errd
		}
		for _, val := range topics.Topics {
			if val == topic {
				return err, errors.New(fmt.Sprintf("NSQ Topic with the same name as %s exists. Round will be terminated. Please firt clean NSQ queues", topic))
			}
		}
		return err, nil
		resp.Body.Close()
	}
	return err, errors.New("no NSQLookupd instances answered the request")
}
func NewNSQQueue(l logger.LogInfoFormat, config queueing.Config) (*NSQ, error) {
	return &NSQ{l, config}, nil
}
