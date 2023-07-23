package flagbearer

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// TODO: Create etcd
func NewEtcd(cfg *config.Config) {
	_, err := clientv3.New(clientv3.Config{
		Endpoints:   nil,
		TLS:         nil,
		Username:    "",
		Password:    "",
		DialOptions: nil,
		Context:     nil,
		Logger:      nil,
		LogConfig:   nil,
	})
	if err != nil {
		return
	}
}
