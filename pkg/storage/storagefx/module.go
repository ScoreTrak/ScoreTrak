package storagefx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/seed"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		storage.NewDB,
	),
	fx.Invoke(
		storage.AutoMigrate,
		seed.DevSeed,
	),
)
