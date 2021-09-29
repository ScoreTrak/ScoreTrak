package report_client

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/report_repo"
	"github.com/gofrs/uuid"
	"log"
	"sync"
)

type Client struct {
	repo   report_repo.Repo
	pubsub queue.MasterStreamPubSub
	cnf    queueing.MasterConfig

	signal      map[uuid.UUID]chan struct{}
	signalMutex *sync.RWMutex
}

func NewReportClient(cnf queueing.MasterConfig, repo report_repo.Repo, pubsub queue.MasterStreamPubSub) *Client {
	return &Client{repo: repo, cnf: cnf, signalMutex: &sync.RWMutex{}, pubsub: pubsub, signal: make(map[uuid.UUID]chan struct{})}
}

func (a *Client) Subscribe() (uuid.UUID, <-chan struct{}) {
	a.signalMutex.Lock()
	defer a.signalMutex.Unlock()
	ch := make(chan struct{}, 1)
	uid, err := uuid.NewV4()
	if err != nil {
		log.Panicln("Unable to generate rabdom UUID")
	}
	a.signal[uid] = ch
	return uid, ch
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
	a.pubsub.NotifyTopic(a.cnf.ChannelPrefix + "_report")
}

func (a *Client) ReportClient() {
	recvChannel := a.pubsub.ReceiveUpdateFromTopic(a.cnf.ChannelPrefix + "_report")
	for {
		<-recvChannel
		a.Publish()
	}
}
