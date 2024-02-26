package seed

import (
	"context"
	"fmt"
	"github.com/scoretrak/scoretrak/internal/entities"
	"github.com/scoretrak/scoretrak/pkg/scorer"
	"time"
)

func DevSeed(ctx context.Context, entitiesClient *entities.Client) {
	ids := []string{
		"00000000000000000000000000",
		"11111111111111111111111111",
		"22222222222222222222222222",
		"33333333333333333333333333",
		"44444444444444444444444444",
		"55555555555555555555555555",
		"66666666666666666666666666",
	}

	_, _ = entitiesClient.ApiToken.Create().SetToken("stt_00000000000000000000000000").Save(ctx)

	_, _ = entitiesClient.Competition.Create().SetName("lockdown").SetDisplayName("Lockdown").SetStartedAt(time.Now()).SetFinishedAt(time.Now().Add(40 * time.Minute)).Save(ctx)

	//roles := entitiesClient.Role.CreateBulk(
	//	entitiesClient.Role.Create().SetID(ids[0]).SetName("competitor"),
	//	entitiesClient.Role.Create().SetID(ids[1]).SetName("spectator"),
	//	entitiesClient.Role.Create().SetID(ids[2]).SetName("staff"),
	//).SaveX(ctx)
	//
	//entitiesClient.RolePermission.CreateBulk(
	//	entitiesClient.RolePermission.Create().SetRole(roles[0]).SetPermission(permission.PERMISSION_COMPETITION_READ),
	//	entitiesClient.RolePermission.Create().SetRole(roles[0]).SetPermission(permission.PERMISSION_TEAM_READ),
	//	entitiesClient.RolePermission.Create().SetRole(roles[0]).SetPermission(permission.PERMISSION_HOST_READ),
	//	entitiesClient.RolePermission.Create().SetRole(roles[0]).SetPermission(permission.PERMISSION_CHECK_READ),
	//	entitiesClient.RolePermission.Create().SetRole(roles[0]).SetPermission(permission.PERMISSION_HOST_SERVICE_READ),
	//	entitiesClient.RolePermission.Create().SetRole(roles[0]).SetPermission(permission.PERMISSION_CUMULATIVE_REPORT_READ),
	//).SaveX(ctx)

	teams, _ := entitiesClient.Team.CreateBulk(
		entitiesClient.Team.Create().SetID(ids[0]).SetName("mainteam1").SetDisplayName("Main Team 1").SetNumber(0),
		entitiesClient.Team.Create().SetID(ids[1]).SetName("mainteam2").SetDisplayName("Main Team 2").SetNumber(1),
		entitiesClient.Team.Create().SetID(ids[2]).SetName("mainteam3").SetDisplayName("Main Team 3").SetNumber(2),
		entitiesClient.Team.Create().SetID(ids[3]).SetName("mainteam4").SetDisplayName("Main Team 4").SetNumber(3),
		entitiesClient.Team.Create().SetID(ids[4]).SetName("mainteam5").SetDisplayName("Main Team 5").SetNumber(4),
		entitiesClient.Team.Create().SetID(ids[5]).SetName("mainteam6").SetDisplayName("Main Team 6").SetNumber(5),
		entitiesClient.Team.Create().SetID(ids[6]).SetName("mainteam7").SetDisplayName("Main Team 7").SetNumber(6),
	).Save(ctx)

	services, _ := entitiesClient.Service.CreateBulk(
		//entitiesClient.Service.Create().SetName("http").SetType(scorer.SERVICE_HTTP).SetDisplayName("HTTP"),
		//entitiesClient.Service.Create().SetName("ssh").SetType(scorer.SERVICE_SSH).SetDisplayName("SSH"),
		entitiesClient.Service.Create().SetID(ids[0]).SetName("dns").SetType(scorer.SERVICE_DNS).SetRoundDelay(1).SetDisplayName("DNS"),
		// entitiesClient.Service.Create().SetID(ids[0]).SetName("dns").SetType(scorer.SERVICE_DNS).SetDisplayName("DNS"),
		entitiesClient.Service.Create().SetID(ids[1]).SetName("ping").SetType(scorer.SERVICE_PING).SetDisplayName("PING"),
		//entitiesClient.Service.Create().SetName("ftp").SetType(scorer.SERVICE_FTP).SetDisplayName("FTP"),
	).Save(ctx)

	for _, serv := range services {
		for _, team := range teams {
			host := entitiesClient.Host.Create().SetAddress("1.1.1.1").SetTeam(team).SaveX(ctx)
			hostservice := entitiesClient.HostService.Create().SetName(fmt.Sprintf("%s_%s_%s", team.Name, host.Address, serv.Name)).SetDisplayName(fmt.Sprintf("%s %s %s", team.DisplayName, host.Address, serv.DisplayName)).SetTeam(team).SetService(serv).SetHost(host).SaveX(ctx)
			if serv.Type == scorer.SERVICE_DNS {
				_, _ = entitiesClient.Property.Create().SetHostservice(hostservice).SetKey("Lookup").SetValue("google.com").Save(ctx)
				//_ = entitiesClient.Property.Create().SetHostservice(hostservice).SetKey("ExpectedOutput").SetValue("8.8.8.8").SaveX(ctx)
			}
			if serv.Type == scorer.SERVICE_PING {
				_, _ = entitiesClient.Property.Create().SetHostservice(hostservice).SetKey("host").SetValue("1.1.1.1").Save(ctx)
				_, _ = entitiesClient.Property.Create().SetHostservice(hostservice).SetKey("protocol").SetValue("ipv4").Save(ctx)
				_, _ = entitiesClient.Property.Create().SetHostservice(hostservice).SetKey("attempts").SetValue("5").Save(ctx)
			}
		}
	}

}
