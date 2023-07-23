// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateApiToken implements createApiToken operation.
//
// Creates a new ApiToken and persists it to storage.
//
// POST /api-tokens
func (UnimplementedHandler) CreateApiToken(ctx context.Context, req *CreateApiTokenReq) (r CreateApiTokenRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateCheck implements createCheck operation.
//
// Creates a new Check and persists it to storage.
//
// POST /checks
func (UnimplementedHandler) CreateCheck(ctx context.Context, req *CreateCheckReq) (r CreateCheckRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateCompetition implements createCompetition operation.
//
// Creates a new Competition and persists it to storage.
//
// POST /competitions
func (UnimplementedHandler) CreateCompetition(ctx context.Context, req *CreateCompetitionReq) (r CreateCompetitionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateHost implements createHost operation.
//
// Creates a new Host and persists it to storage.
//
// POST /hosts
func (UnimplementedHandler) CreateHost(ctx context.Context, req *CreateHostReq) (r CreateHostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateHostService implements createHostService operation.
//
// Creates a new HostService and persists it to storage.
//
// POST /host-services
func (UnimplementedHandler) CreateHostService(ctx context.Context, req *CreateHostServiceReq) (r CreateHostServiceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateHostServiceReport implements createHostServiceReport operation.
//
// Creates a new HostServiceReport and persists it to storage.
//
// POST /host-service-reports
func (UnimplementedHandler) CreateHostServiceReport(ctx context.Context, req *CreateHostServiceReportReq) (r CreateHostServiceReportRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateProperty implements createProperty operation.
//
// Creates a new Property and persists it to storage.
//
// POST /properties
func (UnimplementedHandler) CreateProperty(ctx context.Context, req *CreatePropertyReq) (r CreatePropertyRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateRound implements createRound operation.
//
// Creates a new Round and persists it to storage.
//
// POST /rounds
func (UnimplementedHandler) CreateRound(ctx context.Context, req *CreateRoundReq) (r CreateRoundRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateService implements createService operation.
//
// Creates a new Service and persists it to storage.
//
// POST /services
func (UnimplementedHandler) CreateService(ctx context.Context, req *CreateServiceReq) (r CreateServiceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateTeam implements createTeam operation.
//
// Creates a new Team and persists it to storage.
//
// POST /teams
func (UnimplementedHandler) CreateTeam(ctx context.Context, req *CreateTeamReq) (r CreateTeamRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateTeamReport implements createTeamReport operation.
//
// Creates a new TeamReport and persists it to storage.
//
// POST /team-reports
func (UnimplementedHandler) CreateTeamReport(ctx context.Context, req *CreateTeamReportReq) (r CreateTeamReportRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteApiToken implements deleteApiToken operation.
//
// Deletes the ApiToken with the requested ID.
//
// DELETE /api-tokens/{id}
func (UnimplementedHandler) DeleteApiToken(ctx context.Context, params DeleteApiTokenParams) (r DeleteApiTokenRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteCheck implements deleteCheck operation.
//
// Deletes the Check with the requested ID.
//
// DELETE /checks/{id}
func (UnimplementedHandler) DeleteCheck(ctx context.Context, params DeleteCheckParams) (r DeleteCheckRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteHost implements deleteHost operation.
//
// Deletes the Host with the requested ID.
//
// DELETE /hosts/{id}
func (UnimplementedHandler) DeleteHost(ctx context.Context, params DeleteHostParams) (r DeleteHostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteHostService implements deleteHostService operation.
//
// Deletes the HostService with the requested ID.
//
// DELETE /host-services/{id}
func (UnimplementedHandler) DeleteHostService(ctx context.Context, params DeleteHostServiceParams) (r DeleteHostServiceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteHostServiceReport implements deleteHostServiceReport operation.
//
// Deletes the HostServiceReport with the requested ID.
//
// DELETE /host-service-reports/{id}
func (UnimplementedHandler) DeleteHostServiceReport(ctx context.Context, params DeleteHostServiceReportParams) (r DeleteHostServiceReportRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteProperty implements deleteProperty operation.
//
// Deletes the Property with the requested ID.
//
// DELETE /properties/{id}
func (UnimplementedHandler) DeleteProperty(ctx context.Context, params DeletePropertyParams) (r DeletePropertyRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteRound implements deleteRound operation.
//
// Deletes the Round with the requested ID.
//
// DELETE /rounds/{id}
func (UnimplementedHandler) DeleteRound(ctx context.Context, params DeleteRoundParams) (r DeleteRoundRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteService implements deleteService operation.
//
// Deletes the Service with the requested ID.
//
// DELETE /services/{id}
func (UnimplementedHandler) DeleteService(ctx context.Context, params DeleteServiceParams) (r DeleteServiceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteTeam implements deleteTeam operation.
//
// Deletes the Team with the requested ID.
//
// DELETE /teams/{id}
func (UnimplementedHandler) DeleteTeam(ctx context.Context, params DeleteTeamParams) (r DeleteTeamRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteTeamReport implements deleteTeamReport operation.
//
// Deletes the TeamReport with the requested ID.
//
// DELETE /team-reports/{id}
func (UnimplementedHandler) DeleteTeamReport(ctx context.Context, params DeleteTeamReportParams) (r DeleteTeamReportRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListApiToken implements listApiToken operation.
//
// List ApiTokens.
//
// GET /api-tokens
func (UnimplementedHandler) ListApiToken(ctx context.Context, params ListApiTokenParams) (r ListApiTokenRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListCheck implements listCheck operation.
//
// List Checks.
//
// GET /checks
func (UnimplementedHandler) ListCheck(ctx context.Context, params ListCheckParams) (r ListCheckRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListHost implements listHost operation.
//
// List Hosts.
//
// GET /hosts
func (UnimplementedHandler) ListHost(ctx context.Context, params ListHostParams) (r ListHostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListHostHostservices implements listHostHostservices operation.
//
// List attached Hostservices.
//
// GET /hosts/{id}/hostservices
func (UnimplementedHandler) ListHostHostservices(ctx context.Context, params ListHostHostservicesParams) (r ListHostHostservicesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListHostService implements listHostService operation.
//
// List HostServices.
//
// GET /host-services
func (UnimplementedHandler) ListHostService(ctx context.Context, params ListHostServiceParams) (r ListHostServiceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListHostServiceChecks implements listHostServiceChecks operation.
//
// List attached Checks.
//
// GET /host-services/{id}/checks
func (UnimplementedHandler) ListHostServiceChecks(ctx context.Context, params ListHostServiceChecksParams) (r ListHostServiceChecksRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListHostServiceProperties implements listHostServiceProperties operation.
//
// List attached Properties.
//
// GET /host-services/{id}/properties
func (UnimplementedHandler) ListHostServiceProperties(ctx context.Context, params ListHostServicePropertiesParams) (r ListHostServicePropertiesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListHostServiceReport implements listHostServiceReport operation.
//
// List HostServiceReports.
//
// GET /host-service-reports
func (UnimplementedHandler) ListHostServiceReport(ctx context.Context, params ListHostServiceReportParams) (r ListHostServiceReportRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListProperty implements listProperty operation.
//
// List Properties.
//
// GET /properties
func (UnimplementedHandler) ListProperty(ctx context.Context, params ListPropertyParams) (r ListPropertyRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListRound implements listRound operation.
//
// List Rounds.
//
// GET /rounds
func (UnimplementedHandler) ListRound(ctx context.Context, params ListRoundParams) (r ListRoundRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListRoundChecks implements listRoundChecks operation.
//
// List attached Checks.
//
// GET /rounds/{id}/checks
func (UnimplementedHandler) ListRoundChecks(ctx context.Context, params ListRoundChecksParams) (r ListRoundChecksRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListRoundChecksLatest implements listRoundChecksLatest operation.
//
// List attached checks.
//
// GET /rounds/latest/checks
func (UnimplementedHandler) ListRoundChecksLatest(ctx context.Context) (r ListRoundChecksLatestRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListService implements listService operation.
//
// List Services.
//
// GET /services
func (UnimplementedHandler) ListService(ctx context.Context, params ListServiceParams) (r ListServiceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListServiceHostservicereports implements listServiceHostservicereports operation.
//
// List attached Hostservicereports.
//
// GET /services/{id}/hostservicereports
func (UnimplementedHandler) ListServiceHostservicereports(ctx context.Context, params ListServiceHostservicereportsParams) (r ListServiceHostservicereportsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListServiceHostservices implements listServiceHostservices operation.
//
// List attached Hostservices.
//
// GET /services/{id}/hostservices
func (UnimplementedHandler) ListServiceHostservices(ctx context.Context, params ListServiceHostservicesParams) (r ListServiceHostservicesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListTeam implements listTeam operation.
//
// List Teams.
//
// GET /teams
func (UnimplementedHandler) ListTeam(ctx context.Context, params ListTeamParams) (r ListTeamRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListTeamHosts implements listTeamHosts operation.
//
// List attached Hosts.
//
// GET /teams/{id}/hosts
func (UnimplementedHandler) ListTeamHosts(ctx context.Context, params ListTeamHostsParams) (r ListTeamHostsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListTeamHostservicereports implements listTeamHostservicereports operation.
//
// List attached Hostservicereports.
//
// GET /teams/{id}/hostservicereports
func (UnimplementedHandler) ListTeamHostservicereports(ctx context.Context, params ListTeamHostservicereportsParams) (r ListTeamHostservicereportsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListTeamHostservices implements listTeamHostservices operation.
//
// List attached Hostservices.
//
// GET /teams/{id}/hostservices
func (UnimplementedHandler) ListTeamHostservices(ctx context.Context, params ListTeamHostservicesParams) (r ListTeamHostservicesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListTeamReport implements listTeamReport operation.
//
// List TeamReports.
//
// GET /team-reports
func (UnimplementedHandler) ListTeamReport(ctx context.Context, params ListTeamReportParams) (r ListTeamReportRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListTeamReportHostservicereports implements listTeamReportHostservicereports operation.
//
// List attached Hostservicereports.
//
// GET /team-reports/{id}/hostservicereports
func (UnimplementedHandler) ListTeamReportHostservicereports(ctx context.Context, params ListTeamReportHostservicereportsParams) (r ListTeamReportHostservicereportsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadApiToken implements readApiToken operation.
//
// Finds the ApiToken with the requested ID and returns it.
//
// GET /api-tokens/{id}
func (UnimplementedHandler) ReadApiToken(ctx context.Context, params ReadApiTokenParams) (r ReadApiTokenRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadCheck implements readCheck operation.
//
// Finds the Check with the requested ID and returns it.
//
// GET /checks/{id}
func (UnimplementedHandler) ReadCheck(ctx context.Context, params ReadCheckParams) (r ReadCheckRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadCheckHostservice implements readCheckHostservice operation.
//
// Find the attached HostService of the Check with the given ID.
//
// GET /checks/{id}/hostservice
func (UnimplementedHandler) ReadCheckHostservice(ctx context.Context, params ReadCheckHostserviceParams) (r ReadCheckHostserviceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadCheckRound implements readCheckRound operation.
//
// Find the attached Round of the Check with the given ID.
//
// GET /checks/{id}/round
func (UnimplementedHandler) ReadCheckRound(ctx context.Context, params ReadCheckRoundParams) (r ReadCheckRoundRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadCompetition implements readCompetition operation.
//
// Finds the Competition with the requested ID and returns it.
//
// GET /competitions/{id}
func (UnimplementedHandler) ReadCompetition(ctx context.Context, params ReadCompetitionParams) (r ReadCompetitionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHost implements readHost operation.
//
// Finds the Host with the requested ID and returns it.
//
// GET /hosts/{id}
func (UnimplementedHandler) ReadHost(ctx context.Context, params ReadHostParams) (r ReadHostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostService implements readHostService operation.
//
// Finds the HostService with the requested ID and returns it.
//
// GET /host-services/{id}
func (UnimplementedHandler) ReadHostService(ctx context.Context, params ReadHostServiceParams) (r ReadHostServiceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostServiceHost implements readHostServiceHost operation.
//
// Find the attached Host of the HostService with the given ID.
//
// GET /host-services/{id}/host
func (UnimplementedHandler) ReadHostServiceHost(ctx context.Context, params ReadHostServiceHostParams) (r ReadHostServiceHostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostServiceHostservicereport implements readHostServiceHostservicereport operation.
//
// Find the attached HostServiceReport of the HostService with the given ID.
//
// GET /host-services/{id}/hostservicereport
func (UnimplementedHandler) ReadHostServiceHostservicereport(ctx context.Context, params ReadHostServiceHostservicereportParams) (r ReadHostServiceHostservicereportRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostServiceReport implements readHostServiceReport operation.
//
// Finds the HostServiceReport with the requested ID and returns it.
//
// GET /host-service-reports/{id}
func (UnimplementedHandler) ReadHostServiceReport(ctx context.Context, params ReadHostServiceReportParams) (r ReadHostServiceReportRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostServiceReportHostservice implements readHostServiceReportHostservice operation.
//
// Find the attached HostService of the HostServiceReport with the given ID.
//
// GET /host-service-reports/{id}/hostservice
func (UnimplementedHandler) ReadHostServiceReportHostservice(ctx context.Context, params ReadHostServiceReportHostserviceParams) (r ReadHostServiceReportHostserviceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostServiceReportService implements readHostServiceReportService operation.
//
// Find the attached Service of the HostServiceReport with the given ID.
//
// GET /host-service-reports/{id}/service
func (UnimplementedHandler) ReadHostServiceReportService(ctx context.Context, params ReadHostServiceReportServiceParams) (r ReadHostServiceReportServiceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostServiceReportTeam implements readHostServiceReportTeam operation.
//
// Find the attached Team of the HostServiceReport with the given ID.
//
// GET /host-service-reports/{id}/team
func (UnimplementedHandler) ReadHostServiceReportTeam(ctx context.Context, params ReadHostServiceReportTeamParams) (r ReadHostServiceReportTeamRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostServiceReportTeamreport implements readHostServiceReportTeamreport operation.
//
// Find the attached TeamReport of the HostServiceReport with the given ID.
//
// GET /host-service-reports/{id}/teamreport
func (UnimplementedHandler) ReadHostServiceReportTeamreport(ctx context.Context, params ReadHostServiceReportTeamreportParams) (r ReadHostServiceReportTeamreportRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostServiceService implements readHostServiceService operation.
//
// Find the attached Service of the HostService with the given ID.
//
// GET /host-services/{id}/service
func (UnimplementedHandler) ReadHostServiceService(ctx context.Context, params ReadHostServiceServiceParams) (r ReadHostServiceServiceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostServiceTeam implements readHostServiceTeam operation.
//
// Find the attached Team of the HostService with the given ID.
//
// GET /host-services/{id}/team
func (UnimplementedHandler) ReadHostServiceTeam(ctx context.Context, params ReadHostServiceTeamParams) (r ReadHostServiceTeamRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostTeam implements readHostTeam operation.
//
// Find the attached Team of the Host with the given ID.
//
// GET /hosts/{id}/team
func (UnimplementedHandler) ReadHostTeam(ctx context.Context, params ReadHostTeamParams) (r ReadHostTeamRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadProperty implements readProperty operation.
//
// Finds the Property with the requested ID and returns it.
//
// GET /properties/{id}
func (UnimplementedHandler) ReadProperty(ctx context.Context, params ReadPropertyParams) (r ReadPropertyRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadPropertyHostservice implements readPropertyHostservice operation.
//
// Find the attached HostService of the Property with the given ID.
//
// GET /properties/{id}/hostservice
func (UnimplementedHandler) ReadPropertyHostservice(ctx context.Context, params ReadPropertyHostserviceParams) (r ReadPropertyHostserviceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadRound implements readRound operation.
//
// Finds the Round with the requested ID and returns it.
//
// GET /rounds/{id}
func (UnimplementedHandler) ReadRound(ctx context.Context, params ReadRoundParams) (r ReadRoundRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadRoundLatest implements readRoundLatest operation.
//
// Get latest completed round.
//
// GET /rounds/latest
func (UnimplementedHandler) ReadRoundLatest(ctx context.Context) (r ReadRoundLatestRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadService implements readService operation.
//
// Finds the Service with the requested ID and returns it.
//
// GET /services/{id}
func (UnimplementedHandler) ReadService(ctx context.Context, params ReadServiceParams) (r ReadServiceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadTeam implements readTeam operation.
//
// Finds the Team with the requested ID and returns it.
//
// GET /teams/{id}
func (UnimplementedHandler) ReadTeam(ctx context.Context, params ReadTeamParams) (r ReadTeamRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadTeamReport implements readTeamReport operation.
//
// Finds the TeamReport with the requested ID and returns it.
//
// GET /team-reports/{id}
func (UnimplementedHandler) ReadTeamReport(ctx context.Context, params ReadTeamReportParams) (r ReadTeamReportRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadTeamReportTeam implements readTeamReportTeam operation.
//
// Find the attached Team of the TeamReport with the given ID.
//
// GET /team-reports/{id}/team
func (UnimplementedHandler) ReadTeamReportTeam(ctx context.Context, params ReadTeamReportTeamParams) (r ReadTeamReportTeamRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadTeamTeamreport implements readTeamTeamreport operation.
//
// Find the attached TeamReport of the Team with the given ID.
//
// GET /teams/{id}/teamreport
func (UnimplementedHandler) ReadTeamTeamreport(ctx context.Context, params ReadTeamTeamreportParams) (r ReadTeamTeamreportRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateApiToken implements updateApiToken operation.
//
// Updates a ApiToken and persists changes to storage.
//
// PATCH /api-tokens/{id}
func (UnimplementedHandler) UpdateApiToken(ctx context.Context, req *UpdateApiTokenReq, params UpdateApiTokenParams) (r UpdateApiTokenRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateCheck implements updateCheck operation.
//
// Updates a Check and persists changes to storage.
//
// PATCH /checks/{id}
func (UnimplementedHandler) UpdateCheck(ctx context.Context, req *UpdateCheckReq, params UpdateCheckParams) (r UpdateCheckRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateCompetition implements updateCompetition operation.
//
// Updates a Competition and persists changes to storage.
//
// PATCH /competitions/{id}
func (UnimplementedHandler) UpdateCompetition(ctx context.Context, req *UpdateCompetitionReq, params UpdateCompetitionParams) (r UpdateCompetitionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateHost implements updateHost operation.
//
// Updates a Host and persists changes to storage.
//
// PATCH /hosts/{id}
func (UnimplementedHandler) UpdateHost(ctx context.Context, req *UpdateHostReq, params UpdateHostParams) (r UpdateHostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateHostService implements updateHostService operation.
//
// Updates a HostService and persists changes to storage.
//
// PATCH /host-services/{id}
func (UnimplementedHandler) UpdateHostService(ctx context.Context, req *UpdateHostServiceReq, params UpdateHostServiceParams) (r UpdateHostServiceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateHostServiceReport implements updateHostServiceReport operation.
//
// Updates a HostServiceReport and persists changes to storage.
//
// PATCH /host-service-reports/{id}
func (UnimplementedHandler) UpdateHostServiceReport(ctx context.Context, req *UpdateHostServiceReportReq, params UpdateHostServiceReportParams) (r UpdateHostServiceReportRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateProperty implements updateProperty operation.
//
// Updates a Property and persists changes to storage.
//
// PATCH /properties/{id}
func (UnimplementedHandler) UpdateProperty(ctx context.Context, req *UpdatePropertyReq, params UpdatePropertyParams) (r UpdatePropertyRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateRound implements updateRound operation.
//
// Updates a Round and persists changes to storage.
//
// PATCH /rounds/{id}
func (UnimplementedHandler) UpdateRound(ctx context.Context, req *UpdateRoundReq, params UpdateRoundParams) (r UpdateRoundRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateService implements updateService operation.
//
// Updates a Service and persists changes to storage.
//
// PATCH /services/{id}
func (UnimplementedHandler) UpdateService(ctx context.Context, req *UpdateServiceReq, params UpdateServiceParams) (r UpdateServiceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateTeam implements updateTeam operation.
//
// Updates a Team and persists changes to storage.
//
// PATCH /teams/{id}
func (UnimplementedHandler) UpdateTeam(ctx context.Context, req *UpdateTeamReq, params UpdateTeamParams) (r UpdateTeamRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateTeamReport implements updateTeamReport operation.
//
// Updates a TeamReport and persists changes to storage.
//
// PATCH /team-reports/{id}
func (UnimplementedHandler) UpdateTeamReport(ctx context.Context, req *UpdateTeamReportReq, params UpdateTeamReportParams) (r UpdateTeamReportRes, _ error) {
	return r, ht.ErrNotImplemented
}
