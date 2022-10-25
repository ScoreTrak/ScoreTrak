package storagefx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		storage.NewDB,
		orm.NewCheckRepo,
		orm.NewConfigRepo,
		orm.NewHostGroupRepo,
		orm.NewHostRepo,
		orm.NewReportRepo,
		orm.NewRoundRepo,
		orm.NewUserRepo,
		orm.NewPolicyRepo,
		orm.NewTeamRepo,
		orm.NewServiceRepo,
		orm.NewServiceGroupRepo,
		orm.NewPropertyRepo,
		util.NewRepoStore,
	),
	fx.Invoke(
		util.CheckDBTimeSync,
		util.CreateAllTables,
		util.LoadConfig,
		util.LoadReport,
	),
)
