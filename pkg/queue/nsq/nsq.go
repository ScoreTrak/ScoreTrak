package nsq

import (
	"errors"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/nsqio/go-nsq"
)

func nsqProducerConfig(conf *nsq.Config, c config.Config) {
	tlsConfig(conf, c)
}

func nsqConsumerConfig(conf *nsq.Config, c config.Config) {
	conf.LookupdPollInterval = time.Second * 1
	conf.MaxInFlight = c.Queue.NSQ.MaxInFlight
	tlsConfig(conf, c)
}

func tlsConfig(conf *nsq.Config, c config.Config) {
	if c.Queue.NSQ.AuthSecret != "" {
		conf.AuthSecret = c.Queue.NSQ.AuthSecret
	}
	if c.Queue.NSQ.ClientRootCA != "" && c.Queue.NSQ.ClientSSLKey != "" && c.Queue.NSQ.ClientSSLCert != "" {
		err := conf.Set("tls_v1", true)
		if err != nil {
			panic(err)
		}
		err = conf.Set("tls_insecure_skip_verify", false)
		if err != nil {
			panic(err)
		}
		err = conf.Set("tls_root_ca_file", c.Queue.NSQ.ClientRootCA)
		if err != nil {
			panic(err)
		}
		err = conf.Set("tls_cert", c.Queue.NSQ.ClientSSLCert)
		if err != nil {
			panic(err)
		}
		err = conf.Set("tls_key", c.Queue.NSQ.ClientSSLKey)
		if err != nil {
			panic(err)
		}
	}
}

func connectConsumer(consumer *nsq.Consumer, c config.Config) (err error) {
	if len(c.Queue.NSQ.NSQLookupd) != 0 {
		err = consumer.ConnectToNSQLookupds(c.Queue.NSQ.NSQLookupd)
	} else {
		err = consumer.ConnectToNSQDs(c.Queue.NSQ.ConsumerNSQDPool)
	}
	return
}

var ErrProvidedBothNSQLookupdAndNSQD = errors.New("must either provide a list of nsqlookupd nodes, or list of nsqd nodes for the consumer, but not both")
var ErrNotProvidedAnyNSQLookupdOrNSQD = errors.New("you haven't provided any nsqlookupd mor nsqd nodes for the consumer")
var ErrNSQDProducerAddressNotProvided = errors.New("must provide nsqd producer address")

func validateNSQConfig(c config.Config) error {
	if c.Queue.NSQ.ProducerNSQD == "" {
		return ErrNSQDProducerAddressNotProvided
	}
	switch {
	case len(c.Queue.NSQ.NSQLookupd) != 0 && len(c.Queue.NSQ.ConsumerNSQDPool) != 0:
		return ErrProvidedBothNSQLookupdAndNSQD
	case len(c.Queue.NSQ.NSQLookupd) == 0 && len(c.Queue.NSQ.ConsumerNSQDPool) == 0:
		return ErrNotProvidedAnyNSQLookupdOrNSQD
	default:
		return nil
	}
}
