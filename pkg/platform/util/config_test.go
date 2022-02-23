package util

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/worker"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"reflect"
	"testing"
)

func TestGenerateWorkerCfg(t *testing.T) {
	type args struct {
		originalCfg config.StaticConfig
		info        worker.Info
	}
	tests := []struct {
		name          string
		args          args
		wantWorkerCfg config.StaticConfig
		wantErr       bool
	}{
		{name: "Working Config",
			args: args{originalCfg: config.StaticConfig{Queue: queueing.Config{Use: "nsq", NSQ: struct {
				ProducerNSQD                 string   `default:"nsqd:4150"`
				IgnoreAllScoresIfWorkerFails bool     `default:"true"`
				Topic                        string   `default:"default"`
				MaxInFlight                  int      `default:"200"`
				AuthSecret                   string   `default:""`
				ClientRootCA                 string   `default:""`
				ClientSSLKey                 string   `default:""`
				ClientSSLCert                string   `default:""`
				ConcurrentHandlers           int      `default:"200"`
				NSQLookupd                   []string `default:"[\"nsqlookupd:4161\"]"`
				ConsumerNSQDPool             []string `default:"[\"\"]"`
			}{}}}, info: worker.Info{Topic: "ping", Label: "internal"}},
			wantWorkerCfg: config.StaticConfig{Queue: queueing.Config{Use: "nsq", NSQ: struct {
				ProducerNSQD                 string   `default:"nsqd:4150"`
				IgnoreAllScoresIfWorkerFails bool     `default:"true"`
				Topic                        string   `default:"default"`
				MaxInFlight                  int      `default:"200"`
				AuthSecret                   string   `default:""`
				ClientRootCA                 string   `default:""`
				ClientSSLKey                 string   `default:""`
				ClientSSLCert                string   `default:""`
				ConcurrentHandlers           int      `default:"200"`
				NSQLookupd                   []string `default:"[\"nsqlookupd:4161\"]"`
				ConsumerNSQDPool             []string `default:"[\"\"]"`
			}{Topic: "ping"}}},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWorkerCfg, err := GenerateWorkerCfg(tt.args.originalCfg, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateWorkerCfg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotWorkerCfg, tt.wantWorkerCfg) {
				t.Errorf("GenerateWorkerCfg() gotWorkerCfg = %v, want %v", gotWorkerCfg, tt.wantWorkerCfg)
			}
		})
	}
}
