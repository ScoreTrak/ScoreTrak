package competition

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
)

//Competition is a struct that holds an aggregate of all models. This is used to upload/export competition as a file.
type Competition struct {
	Config        *config.DynamicConfig
	Report        *report.Report
	HostGroups    []*hostgroup.HostGroup
	Hosts         []*host.Host
	Teams         []*team.Team
	Services      []*service.Service
	ServiceGroups []*servicegroup.ServiceGroup
	Rounds        []*round.Round
	Properties    []*property.Property
	Checks        []*check.Check
	Users         []*user.User
	Policy        *policy.Policy
}
