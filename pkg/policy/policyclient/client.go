package policyclient

import (
	"context"
	"go.uber.org/fx"
	"log"
	"sync"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/copier"
)

// Client Policy allows for eventually consistent way to distribute pkg/Policy struct throughout ScoreTrak instances.
// This is needed because certain API routes evaluate Policy on every call, and retrieving Policy from database is very expensive, hence having eventually consistent copy is much more efficient.
type Client struct {
	policy      *policy.Policy
	policyMutex *sync.RWMutex

	repo   policyrepo.Repo
	pubsub queue.MasterStreamPubSub
	cnf    queueing.MasterConfig

	signal      map[uuid.UUID]chan struct{}
	signalMutex *sync.RWMutex
}

func NewPolicyClient(policy *policy.Policy, cnf queueing.MasterConfig, repo policyrepo.Repo, pubsub queue.MasterStreamPubSub) *Client {
	return &Client{policy: policy, policyMutex: &sync.RWMutex{}, repo: repo, cnf: cnf, signalMutex: &sync.RWMutex{}, pubsub: pubsub, signal: make(map[uuid.UUID]chan struct{})}
}

func (a *Client) GetAllowUnauthenticatedUsers() bool {
	a.policyMutex.RLock()
	defer a.policyMutex.RUnlock()
	return *a.policy.AllowUnauthenticatedUsers
}
func (a *Client) GetAllowChangingUsernamesAndPasswords() bool {
	a.policyMutex.RLock()
	defer a.policyMutex.RUnlock()
	return *a.policy.AllowChangingUsernamesAndPasswords
}
func (a *Client) GetAllowRedTeamLaunchingServiceTestsManually() bool {
	a.policyMutex.RLock()
	defer a.policyMutex.RUnlock()
	return *a.policy.AllowRedTeamLaunchingServiceTestsManually
}
func (a *Client) GetShowPoints() bool {
	a.policyMutex.RLock()
	defer a.policyMutex.RUnlock()
	return *a.policy.ShowPoints
}
func (a *Client) GetShowAddresses() bool {
	a.policyMutex.RLock()
	defer a.policyMutex.RUnlock()
	return *a.policy.ShowAddresses
}

func (a *Client) Subscribe() (uuid.UUID, <-chan struct{}) {
	a.signalMutex.Lock()
	defer a.signalMutex.Unlock()
	channel := make(chan struct{}, 1)
	uid, err := uuid.NewV4()
	if err != nil {
		log.Panicln("Unable to generate random UUID")
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
	a.pubsub.NotifyTopic(a.cnf.ChannelPrefix + "_policy")
}

func (a *Client) RefreshLocalPolicy() {
	tempPolicy, err := a.repo.Get(context.Background())
	if err != nil {
		log.Panicf("Unable to retrieve policy. Make sure database is reachable")
	}
	a.policyMutex.Lock()
	defer a.policyMutex.Unlock()
	err = copier.Copy(a.policy, tempPolicy)
	if err != nil {
		log.Panicf("Unable to copy policy into destination policy. This is likely a bug")
	}
}

func (a *Client) PolicyClient() {
	recvChannel := a.pubsub.ReceiveUpdateFromTopic(a.cnf.ChannelPrefix + "_policy")
	forceSync := time.NewTimer(time.Duration(a.cnf.ReportForceRefreshSeconds) * time.Second)
	a.RefreshLocalPolicy()
	for {
		select {
		case <-forceSync.C:
			a.RefreshLocalPolicy()
		case <-recvChannel:
			a.RefreshLocalPolicy()
			a.Publish()
		}
	}
}

func (a *Client) GetPolicy() *policy.Policy {
	a.policyMutex.RLock()
	defer a.policyMutex.RUnlock()
	cp := &policy.Policy{}
	err := copier.Copy(&cp, a.policy)
	if err != nil {
		log.Panicf("Unable to copy policy into destination policy. This is likely a bug")
	}
	return cp
}

func InitPolicyClient(lc fx.Lifecycle, policyClient *Client) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Starting Policy Client")
			go policyClient.PolicyClient()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping Policy Client")
			return nil
		},
	})
}
