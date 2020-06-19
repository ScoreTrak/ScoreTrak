package nsq

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/exec"
	"ScoreTrak/pkg/exec/resolver"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/queue/queueing"
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	"net/http"
	"sync"
	"time"
)

type NSQ struct {
	l logger.LogInfoFormat
}

func (n NSQ) Send(sds []*queueing.ScoringData) ([]*queueing.QCheck, error, error) {
	c := config.GetStaticConfig()
	addresses := n.GenerateNSQLookupdAddresses(c.Queue.NSQ.NSQLookupd.Hosts, c.Queue.NSQ.NSQLookupd.Port)
	returningTopicName := queueing.TopicFromServiceRound(&sds[0].Service, sds[0].RoundID)
	bErr, tErr := n.TopicAbsent(returningTopicName, addresses)
	if tErr != nil {
		return nil, bErr, tErr
	}
	confp := nsq.NewConfig()
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%s", c.Queue.NSQ.NSQD.Host, c.Queue.NSQ.NSQD.Port), confp)
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
		return nil, bErr, tErr
	}
	defer consumer.Stop()
	ret := make([]*queueing.QCheck, len(sds))
	consumer.ChangeMaxInFlight(len(sds))
	wg := &sync.WaitGroup{}
	wg.Add(len(sds))
	consumer.SetLoggerLevel(nsq.LogLevelError)
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
		defer wg.Done()
		buf := bytes.NewBuffer(m.Body)
		var qc queueing.QCheck
		if err := gob.NewDecoder(buf).Decode(&qc); err != nil {
			n.l.Error(err)
			return err
		}
		for i, sd := range sds {
			if sd.Service.ID == qc.Service.ID {
				ret[i] = &qc
			}
		}
		return nil
	}), len(sds))

	err = consumer.ConnectToNSQLookupds(addresses)
	if err != nil {
		return nil, bErr, err
	}
	if queueing.WaitTimeout(wg, sds[0].Deadline) {
		return nil, bErr, errors.New("round took too long to score. this might be due to many reasons like a worker going down, or the number of rounds being too big for one master")
	}
	return ret, bErr, nil
}

func (n NSQ) Receive() {
	c := config.GetConfig()
	conf := nsq.NewConfig()
	conf.LookupdPollInterval = time.Second * 2
	conf.MaxInFlight = c.Queue.NSQ.MaxInFlight
	consumer, err := nsq.NewConsumer(c.Queue.NSQ.Topic, "channel", conf)
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
				n.l.Error(err)
				qc := queueing.QCheck{Service: sd.Service, Passed: false, Log: "Encountered an unexpected error during the check. This is most likely a bug", Err: err.Error(), RoundID: sd.RoundID}
				n.Acknowledge(qc)
				return
			}
		}()
		if err := gob.NewDecoder(buf).Decode(&sd); err != nil {
			panic(err)
		}
		if time.Now().After(sd.Deadline) {
			n.Acknowledge(queueing.QCheck{Service: sd.Service, Passed: false, Log: "", Err: "The check arrived late to the worker", RoundID: sd.RoundID})
			return nil
		}
		executable := resolver.ExecutableByName(sd.Service.Name)
		exec.UpdateExecutableProperties(executable, sd.Properties)
		ctx := context.Background()
		execDeadline := sd.Deadline.Add(-2 * time.Second)
		ctx, cancel := context.WithDeadline(ctx, execDeadline)
		defer cancel()
		e := exec.NewExec(ctx, sd.Host, executable, n.l)
		wg := sync.WaitGroup{}
		wg.Add(1)
		var (
			passed bool
			log    string
			err    error
		)
		go func() {
			passed, log, err = e.Execute()
			wg.Done()
		}()
		if queueing.WaitTimeout(&wg, execDeadline.Add(time.Second)) {
			panic(errors.New("check timed out, which should not have happened. this is most likely a bug. Please check logs for more info"))
		}
		var errstr string
		if err != nil {
			errstr = err.Error()
		}
		qc := queueing.QCheck{Service: sd.Service, Passed: passed, Log: log, Err: errstr, RoundID: sd.RoundID}
		if time.Now().After(sd.Deadline) {
			n.l.Error("Service scored late. Please fix the implementation of the following service: ", sd, qc)
		}
		n.Acknowledge(qc)
		return nil

	}), c.Queue.NSQ.ConcurrentHandlers)
	addresses := n.GenerateNSQLookupdAddresses(c.Queue.NSQ.NSQLookupd.Hosts, c.Queue.NSQ.NSQLookupd.Port)
	err = consumer.ConnectToNSQLookupds(addresses)
	if err != nil {
		panic(err)
	}
	select {}

}

func (n NSQ) GenerateNSQLookupdAddresses(hostNames []string, port string) []string {
	var addresses []string
	for _, h := range hostNames {
		addresses = append(addresses, fmt.Sprintf("%s:%s", h, port))
	}
	return addresses
}

func (n NSQ) Acknowledge(q queueing.QCheck) {
	c := config.GetStaticConfig()
	confp := nsq.NewConfig()
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%s", c.Queue.NSQ.NSQD.Host, c.Queue.NSQ.NSQD.Port), confp)
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
		fmt.Println(fmt.Sprintf("http://%s/topic/delete?topic=%s", a, topic))
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
func NewNSQQueue(l logger.LogInfoFormat) (*NSQ, error) {
	return &NSQ{l}, nil
}
