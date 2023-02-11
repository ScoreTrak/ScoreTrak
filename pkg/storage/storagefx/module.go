package storagefx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition/competitionservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup/hostgroupservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertyservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/serviceservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup/servicegroupservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/teamservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userservice"
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
	// Add Service Pattern Implementations
	fx.Provide(
		checkservice.NewCheckServ,
		hostgroupservice.NewHostGroupServ,
		hostservice.NewHostServ,
		propertyservice.NewPropertyServ,
		configservice.NewConfigServ,
		configservice.NewStaticConfigServ,
		competitionservice.NewCompetitionServ,
		policyservice.NewPolicyServ,
		reportservice.NewReportServ,
		userservice.NewUserServ,
		teamservice.NewTeamServ,
		servicegroupservice.NewServiceGroupServ,
		roundservice.NewRoundServ,
		serviceservice.NewServiceServ,
	),
	fx.Invoke(
		util.CheckDBTimeSync,
	),
)
