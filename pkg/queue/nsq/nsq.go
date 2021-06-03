package nsq

import (
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/nsqio/go-nsq"
	"time"
)

func nsqProducerConfig(conf *nsq.Config, config queueing.Config) {
	tlsConfig(conf, config)
}

func nsqConsumerConfig(conf *nsq.Config, config queueing.Config) {
	conf.LookupdPollInterval = time.Second * 1
	conf.MaxInFlight = config.NSQ.MaxInFlight
	tlsConfig(conf, config)
}

func tlsConfig(conf *nsq.Config, config queueing.Config) {
	if config.NSQ.AuthSecret != "" {
		conf.AuthSecret = config.NSQ.AuthSecret
	}
	if config.NSQ.ClientRootCA != "" && config.NSQ.ClientSSLKey != "" && config.NSQ.ClientSSLCert != "" {
		err := conf.Set("tls_v1", true)
		if err != nil {
			panic(err)
		}
		err = conf.Set("tls_insecure_skip_verify", false)
		if err != nil {
			panic(err)
		}
		err = conf.Set("tls_root_ca_file", config.NSQ.ClientRootCA)
		if err != nil {
			panic(err)
		}
		err = conf.Set("tls_cert", config.NSQ.ClientSSLCert)
		if err != nil {
			panic(err)
		}
		err = conf.Set("tls_key", config.NSQ.ClientSSLKey)
		if err != nil {
			panic(err)
		}
	}
}

func connectConsumer(consumer *nsq.Consumer, config queueing.Config) (err error) {
	if len(config.NSQ.NSQLookupd) != 0 {
		err = consumer.ConnectToNSQLookupds(config.NSQ.NSQLookupd)
	} else {
		err = consumer.ConnectToNSQDs(config.NSQ.ConsumerNSQDPool)
	}
	return nil
}

func validateNSQConfig(config queueing.Config) error {
	if config.NSQ.ProducerNSQD == "" {
		return errors.New("must provide nsqd producer address")
	}
	//Emulates exclusive-or
	if (len(config.NSQ.NSQLookupd) != 0) != (len(config.NSQ.ConsumerNSQDPool) != 0) {
		return errors.New("must either provide a list of nsqlookupd nodes, or list of nsqd nodes for the consumer, but not both")
	}
	return nil
}
