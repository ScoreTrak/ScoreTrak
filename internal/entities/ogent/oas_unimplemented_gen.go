// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

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

// CreateHostGroup implements createHostGroup operation.
//
// Creates a new HostGroup and persists it to storage.
//
// POST /host-groups
func (UnimplementedHandler) CreateHostGroup(ctx context.Context, req *CreateHostGroupReq) (r CreateHostGroupRes, _ error) {
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

// CreateReport implements createReport operation.
//
// Creates a new Report and persists it to storage.
//
// POST /reports
func (UnimplementedHandler) CreateReport(ctx context.Context, req *CreateReportReq) (r CreateReportRes, _ error) {
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

// CreateUser implements createUser operation.
//
// Creates a new User and persists it to storage.
//
// POST /users
func (UnimplementedHandler) CreateUser(ctx context.Context, req *CreateUserReq) (r CreateUserRes, _ error) {
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

// DeleteCompetition implements deleteCompetition operation.
//
// Deletes the Competition with the requested ID.
//
// DELETE /competitions/{id}
func (UnimplementedHandler) DeleteCompetition(ctx context.Context, params DeleteCompetitionParams) (r DeleteCompetitionRes, _ error) {
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

// DeleteHostGroup implements deleteHostGroup operation.
//
// Deletes the HostGroup with the requested ID.
//
// DELETE /host-groups/{id}
func (UnimplementedHandler) DeleteHostGroup(ctx context.Context, params DeleteHostGroupParams) (r DeleteHostGroupRes, _ error) {
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

// DeleteReport implements deleteReport operation.
//
// Deletes the Report with the requested ID.
//
// DELETE /reports/{id}
func (UnimplementedHandler) DeleteReport(ctx context.Context, params DeleteReportParams) (r DeleteReportRes, _ error) {
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

// DeleteUser implements deleteUser operation.
//
// Deletes the User with the requested ID.
//
// DELETE /users/{id}
func (UnimplementedHandler) DeleteUser(ctx context.Context, params DeleteUserParams) (r DeleteUserRes, _ error) {
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

// ListCompetition implements listCompetition operation.
//
// List Competitions.
//
// GET /competitions
func (UnimplementedHandler) ListCompetition(ctx context.Context, params ListCompetitionParams) (r ListCompetitionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListCompetitionTeams implements listCompetitionTeams operation.
//
// List attached Teams.
//
// GET /competitions/{id}/teams
func (UnimplementedHandler) ListCompetitionTeams(ctx context.Context, params ListCompetitionTeamsParams) (r ListCompetitionTeamsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListCompetitionUsers implements listCompetitionUsers operation.
//
// List attached Users.
//
// GET /competitions/{id}/users
func (UnimplementedHandler) ListCompetitionUsers(ctx context.Context, params ListCompetitionUsersParams) (r ListCompetitionUsersRes, _ error) {
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

// ListHostGroup implements listHostGroup operation.
//
// List HostGroups.
//
// GET /host-groups
func (UnimplementedHandler) ListHostGroup(ctx context.Context, params ListHostGroupParams) (r ListHostGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListHostGroupHosts implements listHostGroupHosts operation.
//
// List attached Hosts.
//
// GET /host-groups/{id}/hosts
func (UnimplementedHandler) ListHostGroupHosts(ctx context.Context, params ListHostGroupHostsParams) (r ListHostGroupHostsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListHostServices implements listHostServices operation.
//
// List attached Services.
//
// GET /hosts/{id}/services
func (UnimplementedHandler) ListHostServices(ctx context.Context, params ListHostServicesParams) (r ListHostServicesRes, _ error) {
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

// ListReport implements listReport operation.
//
// List Reports.
//
// GET /reports
func (UnimplementedHandler) ListReport(ctx context.Context, params ListReportParams) (r ListReportRes, _ error) {
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

// ListService implements listService operation.
//
// List Services.
//
// GET /services
func (UnimplementedHandler) ListService(ctx context.Context, params ListServiceParams) (r ListServiceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListServiceChecks implements listServiceChecks operation.
//
// List attached Checks.
//
// GET /services/{id}/checks
func (UnimplementedHandler) ListServiceChecks(ctx context.Context, params ListServiceChecksParams) (r ListServiceChecksRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListServiceProperties implements listServiceProperties operation.
//
// List attached Properties.
//
// GET /services/{id}/properties
func (UnimplementedHandler) ListServiceProperties(ctx context.Context, params ListServicePropertiesParams) (r ListServicePropertiesRes, _ error) {
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

// ListTeamUsers implements listTeamUsers operation.
//
// List attached Users.
//
// GET /teams/{id}/users
func (UnimplementedHandler) ListTeamUsers(ctx context.Context, params ListTeamUsersParams) (r ListTeamUsersRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListUser implements listUser operation.
//
// List Users.
//
// GET /users
func (UnimplementedHandler) ListUser(ctx context.Context, params ListUserParams) (r ListUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListUserCompetitions implements listUserCompetitions operation.
//
// List attached Competitions.
//
// GET /users/{id}/competitions
func (UnimplementedHandler) ListUserCompetitions(ctx context.Context, params ListUserCompetitionsParams) (r ListUserCompetitionsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListUserTeams implements listUserTeams operation.
//
// List attached Teams.
//
// GET /users/{id}/teams
func (UnimplementedHandler) ListUserTeams(ctx context.Context, params ListUserTeamsParams) (r ListUserTeamsRes, _ error) {
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

// ReadCheckCompetition implements readCheckCompetition operation.
//
// Find the attached Competition of the Check with the given ID.
//
// GET /checks/{id}/competition
func (UnimplementedHandler) ReadCheckCompetition(ctx context.Context, params ReadCheckCompetitionParams) (r ReadCheckCompetitionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadCheckRounds implements readCheckRounds operation.
//
// Find the attached Round of the Check with the given ID.
//
// GET /checks/{id}/rounds
func (UnimplementedHandler) ReadCheckRounds(ctx context.Context, params ReadCheckRoundsParams) (r ReadCheckRoundsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadCheckServices implements readCheckServices operation.
//
// Find the attached Service of the Check with the given ID.
//
// GET /checks/{id}/services
func (UnimplementedHandler) ReadCheckServices(ctx context.Context, params ReadCheckServicesParams) (r ReadCheckServicesRes, _ error) {
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

// ReadHostCompetition implements readHostCompetition operation.
//
// Find the attached Competition of the Host with the given ID.
//
// GET /hosts/{id}/competition
func (UnimplementedHandler) ReadHostCompetition(ctx context.Context, params ReadHostCompetitionParams) (r ReadHostCompetitionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostGroup implements readHostGroup operation.
//
// Finds the HostGroup with the requested ID and returns it.
//
// GET /host-groups/{id}
func (UnimplementedHandler) ReadHostGroup(ctx context.Context, params ReadHostGroupParams) (r ReadHostGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostGroupCompetition implements readHostGroupCompetition operation.
//
// Find the attached Competition of the HostGroup with the given ID.
//
// GET /host-groups/{id}/competition
func (UnimplementedHandler) ReadHostGroupCompetition(ctx context.Context, params ReadHostGroupCompetitionParams) (r ReadHostGroupCompetitionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostGroupTeam implements readHostGroupTeam operation.
//
// Find the attached Team of the HostGroup with the given ID.
//
// GET /host-groups/{id}/team
func (UnimplementedHandler) ReadHostGroupTeam(ctx context.Context, params ReadHostGroupTeamParams) (r ReadHostGroupTeamRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadHostHostGroup implements readHostHostGroup operation.
//
// Find the attached HostGroup of the Host with the given ID.
//
// GET /hosts/{id}/host-group
func (UnimplementedHandler) ReadHostHostGroup(ctx context.Context, params ReadHostHostGroupParams) (r ReadHostHostGroupRes, _ error) {
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

// ReadPropertyCompetition implements readPropertyCompetition operation.
//
// Find the attached Competition of the Property with the given ID.
//
// GET /properties/{id}/competition
func (UnimplementedHandler) ReadPropertyCompetition(ctx context.Context, params ReadPropertyCompetitionParams) (r ReadPropertyCompetitionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadPropertyServices implements readPropertyServices operation.
//
// Find the attached Service of the Property with the given ID.
//
// GET /properties/{id}/services
func (UnimplementedHandler) ReadPropertyServices(ctx context.Context, params ReadPropertyServicesParams) (r ReadPropertyServicesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadPropertyTeam implements readPropertyTeam operation.
//
// Find the attached Team of the Property with the given ID.
//
// GET /properties/{id}/team
func (UnimplementedHandler) ReadPropertyTeam(ctx context.Context, params ReadPropertyTeamParams) (r ReadPropertyTeamRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadReport implements readReport operation.
//
// Finds the Report with the requested ID and returns it.
//
// GET /reports/{id}
func (UnimplementedHandler) ReadReport(ctx context.Context, params ReadReportParams) (r ReadReportRes, _ error) {
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

// ReadRoundCompetition implements readRoundCompetition operation.
//
// Find the attached Competition of the Round with the given ID.
//
// GET /rounds/{id}/competition
func (UnimplementedHandler) ReadRoundCompetition(ctx context.Context, params ReadRoundCompetitionParams) (r ReadRoundCompetitionRes, _ error) {
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

// ReadServiceCompetition implements readServiceCompetition operation.
//
// Find the attached Competition of the Service with the given ID.
//
// GET /services/{id}/competition
func (UnimplementedHandler) ReadServiceCompetition(ctx context.Context, params ReadServiceCompetitionParams) (r ReadServiceCompetitionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadServiceHosts implements readServiceHosts operation.
//
// Find the attached Host of the Service with the given ID.
//
// GET /services/{id}/hosts
func (UnimplementedHandler) ReadServiceHosts(ctx context.Context, params ReadServiceHostsParams) (r ReadServiceHostsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadServiceTeam implements readServiceTeam operation.
//
// Find the attached Team of the Service with the given ID.
//
// GET /services/{id}/team
func (UnimplementedHandler) ReadServiceTeam(ctx context.Context, params ReadServiceTeamParams) (r ReadServiceTeamRes, _ error) {
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

// ReadTeamCompetition implements readTeamCompetition operation.
//
// Find the attached Competition of the Team with the given ID.
//
// GET /teams/{id}/competition
func (UnimplementedHandler) ReadTeamCompetition(ctx context.Context, params ReadTeamCompetitionParams) (r ReadTeamCompetitionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadUser implements readUser operation.
//
// Finds the User with the requested ID and returns it.
//
// GET /users/{id}
func (UnimplementedHandler) ReadUser(ctx context.Context, params ReadUserParams) (r ReadUserRes, _ error) {
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

// UpdateHostGroup implements updateHostGroup operation.
//
// Updates a HostGroup and persists changes to storage.
//
// PATCH /host-groups/{id}
func (UnimplementedHandler) UpdateHostGroup(ctx context.Context, req *UpdateHostGroupReq, params UpdateHostGroupParams) (r UpdateHostGroupRes, _ error) {
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

// UpdateReport implements updateReport operation.
//
// Updates a Report and persists changes to storage.
//
// PATCH /reports/{id}
func (UnimplementedHandler) UpdateReport(ctx context.Context, req *UpdateReportReq, params UpdateReportParams) (r UpdateReportRes, _ error) {
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

// UpdateUser implements updateUser operation.
//
// Updates a User and persists changes to storage.
//
// PATCH /users/{id}
func (UnimplementedHandler) UpdateUser(ctx context.Context, req *UpdateUserReq, params UpdateUserParams) (r UpdateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}
