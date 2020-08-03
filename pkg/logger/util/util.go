package util

import "github.com/ScoreTrak/ScoreTrak/pkg/logger"

func SetupLogger(c logger.Config) logger.LogInfoFormat {
	l, err := logger.NewLogger(c)
	if err != nil {
		panic(err)
	}
	return l
}
