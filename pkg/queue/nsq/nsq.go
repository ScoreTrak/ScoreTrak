package nsq

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/exec"
	"ScoreTrak/pkg/exec/resolver"
	"ScoreTrak/pkg/queue/queueing"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/nsqio/go-nsq"
	"strconv"
	"sync"
	"time"
)

type NSQ struct{}

func (n NSQ) Send(sds []*queueing.ScoringData) []*queueing.QCheck {
	//TODO: TERMINATION BASED ON TIMEOUT
	c := config.GetStaticConfig()
	confp := nsq.NewConfig()
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%s", c.Queue.NSQ.NSQD.Host, c.Queue.NSQ.NSQD.Port), confp)
	if err != nil {
		panic(err)
	}

	for _, sd := range sds {
		buf := &bytes.Buffer{}
		if err := gob.NewEncoder(buf).Encode(sd); err != nil {
			panic(err)
		}
		err = producer.Publish(sd.Service.Group, buf.Bytes())
		if err != nil {
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
	wg.Wait()
	consumer.Stop()
	return ret
}

func (n NSQ) Receive() {
	//TODO: TERMINATION BASED ON TIMEOUT
	c := config.GetConfig()
	conf := nsq.NewConfig()
	conf.LookupdPollInterval = time.Second * 1
	conf.MaxInFlight = c.Queue.NSQ.MaxInFlight
	consumer, err := nsq.NewConsumer(c.Queue.NSQ.Topic, "channel", conf)
	if err != nil {
		panic(err)
	}
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
		buf := bytes.NewBuffer(m.Body)
		var sd queueing.ScoringData
		if err := gob.NewDecoder(buf).Decode(&sd); err != nil {
			panic(err)
		}
		executable := resolver.ExecutableByName(sd.Service.Name)
		exec.UpdateExecutableProperties(executable, sd.Properties)
		e := exec.NewExec(sd.Timeout, sd.Host, executable)
		fmt.Println(fmt.Sprintf("Executing a check for service ID %d for round %d", sd.Service.ID, sd.RoundID))
		err := e.Validate()
		if err != nil {
			qc := queueing.QCheck{Service: sd.Service, Passed: false, Log: "", Err: err.Error(), RoundID: sd.RoundID}
			n.Acknowledge(qc)
			return nil
		}
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

func NewNSQQueue() (*NSQ, error) {
	return &NSQ{}, nil
}
