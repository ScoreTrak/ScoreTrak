package reportclient

import (
	"context"
	"log"
	"sync"

	"go.uber.org/fx"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportrepo"
	"github.com/gofrs/uuid"
)

type Client struct {
	repo   reportrepo.Repo
	pubsub queue.MasterStreamPubSub
	cnf    config.Config

	signal      map[uuid.UUID]chan struct{}
	signalMutex *sync.RWMutex
}

func NewReportClient(cnf config.Config, repo reportrepo.Repo, pubsub queue.MasterStreamPubSub) *Client {
	return &Client{repo: repo, cnf: cnf, signalMutex: &sync.RWMutex{},
		pubsub: pubsub, signal: make(map[uuid.UUID]chan struct{}),
	}
}

func (a *Client) Subscribe() (uuid.UUID, <-chan struct{}) {
	a.signalMutex.Lock()
	defer a.signalMutex.Unlock()
	channel := make(chan struct{}, 1)
	uid, err := uuid.NewV4()
	if err != nil {
		log.Panicln("Unable to generate rabdom UUID")
	}
	a.signal[uid] = channel
	return uid, channel
}

func (a *Client) Unsubscribe(uid uuid.UUID) {
	a.signalMutex.Lock()
	defer a.signalMutex.Unlock()
	delete(a.signal, uid)
}

func (a *Client) Publish() {
	a.signalMutex.RLock()
	defer a.signalMutex.RUnlock()
	for _, ch := range a.signal {
		ch <- struct{}{}
	}
}

func (a *Client) Notify() {
	a.pubsub.NotifyTopic(a.cnf.PubSubConfig.ChannelPrefix + "_report")
}

func (a *Client) ReportClient() {
	recvChannel := a.pubsub.ReceiveUpdateFromTopic(a.cnf.PubSubConfig.ChannelPrefix + "_report")
	for {
		<-recvChannel
		a.Publish()
	}
}

func InitReportClient(lc fx.Lifecycle, reportClient *Client) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Starting Report Client")
			go reportClient.ReportClient()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping Report Client")
			return nil
		},
	})
}
