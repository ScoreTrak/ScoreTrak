package storagefx

import (
	"github.com/scoretrak/scoretrak/pkg/storage"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		storage.NewDB,
	),
	fx.Invoke(
		storage.AutoMigrate,
		//seed.DevSeed,
	),
)
