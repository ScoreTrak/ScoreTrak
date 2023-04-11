package storagefx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		storage.NewDB,
	),
	fx.Invoke(
		storage.AutoMigrate,
	// 	util.CheckDBTimeSync,
	),
)
