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
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	"strconv"
	"sync"
	"time"
)

type NSQ struct {
	l logger.LogInfoFormat
}

func (n NSQ) Send(sds []*queueing.ScoringData) []*queueing.QCheck {
	//TODO: TERMINATION BASED ON TIMEOUT
	c := config.GetStaticConfig()
	confp := nsq.NewConfig()
	confp.LookupdPollInterval = time.Second * 1
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%s", c.Queue.NSQ.NSQD.Host, c.Queue.NSQ.NSQD.Port), confp)
	if err != nil {
		n.l.Error(err)
		panic(err)
	}
	var buffers []*bytes.Buffer
	for _, sd := range sds {
		buf := &bytes.Buffer{}
		if err := gob.NewEncoder(buf).Encode(sd); err != nil {
			n.l.Error(err)
			panic(err)
		}
		buffers = append(buffers)
		err = producer.Publish(sd.Service.Group, buf.Bytes())
		if err != nil {
			n.l.Error(err)
			panic(err)
		}
	}
	producer.Stop()
	confc := nsq.NewConfig()
	confc.LookupdPollInterval = time.Second * 3
	consumer, err := nsq.NewConsumer(strconv.FormatUint(sds[0].RoundID, 10)+"_ack", "channel", confc)
	if err != nil {
		panic(err)
	}
	defer consumer.Stop()
	ret := make([]*queueing.QCheck, len(sds))
	consumer.ChangeMaxInFlight(len(sds))
	wg := &sync.WaitGroup{}
	wg.Add(len(sds))

	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
		defer wg.Done()
		buf := bytes.NewBuffer(m.Body)
		var qc queueing.QCheck
		if err := gob.NewDecoder(buf).Decode(&qc); err != nil {
			panic(err)
		}
		for i, sd := range sds {
			if sd.Service.ID == qc.Service.ID {
				ret[i] = &qc
			}
		}
		return nil
	}), len(sds))
	err = consumer.ConnectToNSQLookupd(fmt.Sprintf("%s:%s", c.Queue.NSQ.NSQLookupd.Host, c.Queue.NSQ.NSQLookupd.Port))
	if err != nil {
		panic(err)
	}
	if queueing.WaitTimeout(wg, sds[0].Deadline) {
		return nil
	}
	//Make waitgroup terminable to not lose gorutine

	return ret
}

func (n NSQ) Receive() {
	c := config.GetConfig()
	conf := nsq.NewConfig()
	conf.LookupdPollInterval = time.Second * 3
	conf.MaxInFlight = c.Queue.NSQ.MaxInFlight
	consumer, err := nsq.NewConsumer(c.Queue.NSQ.Topic, "channel", conf)
	if err != nil {
		panic(err)
	}
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
		if time.Now().After(sd.Deadline) {
			n.Acknowledge(queueing.QCheck{Service: sd.Service, Passed: false, Log: "", Err: "The check arrived late to the worker", RoundID: sd.RoundID})
			return nil
		}
		if err := gob.NewDecoder(buf).Decode(&sd); err != nil {
			panic(err)
		}
		executable := resolver.ExecutableByName(sd.Service.Name)
		exec.UpdateExecutableProperties(executable, sd.Properties)
		ctx := context.Background()
		ctx, cancel := context.WithDeadline(ctx, sd.Deadline.Add(-time.Second))
		defer cancel()
		e := exec.NewExec(ctx, sd.Host, executable)
		fmt.Println(fmt.Sprintf("Executing a check for service ID %d for round %d", sd.Service.ID, sd.RoundID))
		passed, log, err := e.Execute()
		var errstr string
		if err != nil {
			errstr = err.Error()
		}
		qc := queueing.QCheck{Service: sd.Service, Passed: passed, Log: log, Err: errstr, RoundID: sd.RoundID}
		n.Acknowledge(qc)
		return nil

	}), c.Queue.NSQ.ConcurrentHandlers)
	err = consumer.ConnectToNSQLookupd(fmt.Sprintf("%s:%s", c.Queue.NSQ.NSQLookupd.Host, c.Queue.NSQ.NSQLookupd.Port))
	if err != nil {
		panic(err)
	}
	select {}

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
	err = producer.Publish(strconv.FormatUint(q.RoundID, 10)+"_ack", buf.Bytes())
	if err != nil {
		panic(err)
	}
	producer.Stop()
}

func NewNSQQueue(l logger.LogInfoFormat) (*NSQ, error) {
	return &NSQ{l}, nil
}
