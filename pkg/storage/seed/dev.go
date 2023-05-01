package seed

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/internal/entities"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec/resolver"
	client "github.com/ory/kratos-client-go"
)

func DevSeed(cfg *config.Config, entitiesClient *entities.Client, authClient *client.APIClient) {
	ctx := context.Background()

	//flow, _, err := authClient.FrontendApi.CreateNativeLoginFlow(ctx).Aal("aal1").Execute()
	//if err != nil {
	//	panic(err)
	//}
	//ulfwpm := client.NewUpdateLoginFlowWithPasswordMethod("scoretrak", "password", "scoretrak")
	//login, _, err := authClient.FrontendApi.UpdateLoginFlow(ctx).Flow(flow.Id).UpdateLoginFlowBody(
	//	client.UpdateLoginFlowWithPasswordMethodAsUpdateLoginFlowBody(ulfwpm)).Execute()
	//if err != nil {
	//	panic(err)
	//}
	//
	//i, _, _ := authClient.IdentityApi.GetIdentity(ctx, login.Session.Identity.Id).Execute()
	//ctx = user.NewContext(ctx, i)

	if !cfg.Prod {
		compIds := []string{
			"01GZ7MDMACFMZ176YXK1JFWZ92",
			"01GZ7MDNKWK1CASYTYSZXWMQA4",
		}

		comps := entitiesClient.Competition.CreateBulk(
			entitiesClient.Competition.Create().SetID(compIds[0]).SetName("main").SetDisplayName("Main"),
			entitiesClient.Competition.Create().SetID(compIds[1]).SetName("test").SetDisplayName("Test"),
		).SaveX(ctx)

		teamIds := []string{
			"01GZ7MFPZ5N43STNFE9T4BB138",
			"01GZ7MFQ8SSBVXFNDEE0Z4EQGH",
			"01GZ7MFQMG0B0GADP1CPW5VK49",
			"01GZ7MYPB5K9A5KPQ6AXYJVS5Y",
			"01GZ7MFRAGH4X81WMKVYR9XE5P",
			"01GZ7MFRN2MNTPSW4BZ0DP78DJ",
		}

		teams := entitiesClient.Team.CreateBulk(
			entitiesClient.Team.Create().SetName("mainteam1").SetID(teamIds[0]).SetDisplayName("Main Team 1").SetNumber(0).SetCompetition(comps[0]),
			entitiesClient.Team.Create().SetName("mainteam2").SetID(teamIds[1]).SetDisplayName("Main Team 2").SetNumber(1).SetCompetition(comps[0]),
			entitiesClient.Team.Create().SetName("mainteam3").SetID(teamIds[2]).SetDisplayName("Main Team 3").SetNumber(2).SetCompetition(comps[0]),
			entitiesClient.Team.Create().SetName("testteam1").SetID(teamIds[3]).SetDisplayName("Test Team 1").SetNumber(0).SetCompetition(comps[1]),
			entitiesClient.Team.Create().SetName("testteam2").SetID(teamIds[4]).SetDisplayName("Test Team 2").SetNumber(1).SetCompetition(comps[1]),
			entitiesClient.Team.Create().SetName("testteam3").SetID(teamIds[5]).SetDisplayName("Test Team 3").SetNumber(2).SetCompetition(comps[1]),
		).SaveX(ctx)

		services := entitiesClient.Service.CreateBulk(
			entitiesClient.Service.Create().SetName("http").SetType(resolver.SERVICE_HTTP).SetDisplayName("LAMP Stack").SetCompetition(comps[0]),
			entitiesClient.Service.Create().SetName("ssh").SetType(resolver.SERVICE_SSH).SetDisplayName("DB Stack").SetCompetition(comps[0]),
			entitiesClient.Service.Create().SetName("dns").SetType(resolver.SERVICE_DNS).SetDisplayName("FTP Stack").SetCompetition(comps[0]),
			entitiesClient.Service.Create().SetName("ftp").SetType(resolver.SERVICE_FTP).SetDisplayName("IMAP Stack").SetCompetition(comps[0]),
			entitiesClient.Service.Create().SetName("http").SetType(resolver.SERVICE_HTTP).SetDisplayName("LAMP Stack").SetCompetition(comps[1]),
			entitiesClient.Service.Create().SetName("ssh").SetType(resolver.SERVICE_SSH).SetDisplayName("DB Stack").SetCompetition(comps[1]),
			entitiesClient.Service.Create().SetName("dns").SetType(resolver.SERVICE_DNS).SetDisplayName("FTP Stack").SetCompetition(comps[1]),
			entitiesClient.Service.Create().SetName("ftp").SetType(resolver.SERVICE_FTP).SetDisplayName("IMAP Stack").SetCompetition(comps[1]),
		).SaveX(ctx)

		for _, service := range services {
			for _, team := range teams {
				host := entitiesClient.Host.Create().SetAddress("1.1.1.1").SetTeam(team).SaveX(ctx)
				hostservice := entitiesClient.HostService.Create().SetTeam(team).SetService(service).SetHost(host).SetName(service.Name).SetDisplayName(service.DisplayName).SaveX(ctx)
				_ = entitiesClient.Property.Create().SetTeam(team).SetHostservice(hostservice).SetKey("a").SetValue("v").SaveX(ctx)
			}
		}
	}

}
