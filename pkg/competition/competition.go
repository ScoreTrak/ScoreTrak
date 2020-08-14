package competition

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
)

type Competition struct {
	Config        *config.DynamicConfig
	Report        *report.Report
	HostGroups    []*host_group.HostGroup
	Hosts         []*host.Host
	Teams         []*team.Team
	Services      []*service.Service
	ServiceGroups []*service_group.ServiceGroup
	Rounds        []*round.Round
	Properties    []*property.Property
	Checks        []*check.Check
}
