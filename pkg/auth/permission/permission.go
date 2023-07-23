package permission

type Permission string

const (
	PERMISSION_API_TOKEN_CREATE    Permission = "scoretrak.api_token.create"
	PERMISSION_API_TOKEN_READ      Permission = "scoretrak.api_token.read"
	PERMISSION_API_TOKEN_UPDATE    Permission = "scoretrak.api_token.update"
	PERMISSION_API_TOKEN_DELETE    Permission = "scoretrak.api_token.delete"
	PERMISSION_COMPETITION_CREATE  Permission = "scoretrak.competition.create"
	PERMISSION_COMPETITION_READ    Permission = "scoretrak.competition.read"
	PERMISSION_COMPETITION_UPDATE  Permission = "scoretrak.competition.update"
	PERMISSION_COMPETITION_DELETE  Permission = "scoretrak.competition.delete"
	PERMISSION_TEAM_CREATE         Permission = "scoretrak.team.create"
	PERMISSION_TEAM_READ           Permission = "scoretrak.team.read"
	PERMISSION_TEAM_UPDATE         Permission = "scoretrak.team.update"
	PERMISSION_TEAM_DELETE         Permission = "scoretrak.team.delete"
	PERMISSION_HOST_CREATE         Permission = "scoretrak.host.create"
	PERMISSION_HOST_READ           Permission = "scoretrak.host.read"
	PERMISSION_HOST_UPDATE         Permission = "scoretrak.host.update"
	PERMISSION_HOST_DELETE         Permission = "scoretrak.host.delete"
	PERMISSION_HOST_SERVICE_CREATE Permission = "scoretrak.host_service.create"
	PERMISSION_HOST_SERVICE_READ   Permission = "scoretrak.host_service.read"
	PERMISSION_HOST_SERVICE_UPDATE Permission = "scoretrak.host_service.update"
	PERMISSION_HOST_SERVICE_DELETE Permission = "scoretrak.host_service.delete"
	PERMISSION_PROPERTY_CREATE     Permission = "scoretrak.property.create"
	PERMISSION_PROPERTY_READ       Permission = "scoretrak.property.read"
	PERMISSION_PROPERTY_UPDATE     Permission = "scoretrak.property.update"
	PERMISSION_PROPERTY_DELETE     Permission = "scoretrak.property.delete"
	PERMISSION_SERVICE_CREATE      Permission = "scoretrak.service.create"
	PERMISSION_SERVICE_READ        Permission = "scoretrak.service.read"
	PERMISSION_SERVICE_UPDATE      Permission = "scoretrak.service.update"
	PERMISSION_SERVICE_DELETE      Permission = "scoretrak.service.delete"
	PERMISSION_CHECK_CREATE        Permission = "scoretrak.check.create"
	PERMISSION_CHECK_READ          Permission = "scoretrak.check.read"
	PERMISSION_CHECK_UPDATE        Permission = "scoretrak.check.update"
	PERMISSION_CHECK_DELETE        Permission = "scoretrak.check.delete"
	PERMISSION_ROUND_CREATE        Permission = "scoretrak.round.create"
	PERMISSION_ROUND_READ          Permission = "scoretrak.round.read"
	PERMISSION_ROUND_UPDATE        Permission = "scoretrak.round.update"
	PERMISSION_ROUND_DELETE        Permission = "scoretrak.round.delete"
)

func (Permission) Values() (kinds []string) {
	for _, s := range []Permission{
		PERMISSION_COMPETITION_CREATE,
		PERMISSION_COMPETITION_READ,
		PERMISSION_COMPETITION_UPDATE,
		PERMISSION_COMPETITION_DELETE,
		PERMISSION_TEAM_CREATE,
		PERMISSION_TEAM_READ,
		PERMISSION_TEAM_UPDATE,
		PERMISSION_TEAM_DELETE,
		PERMISSION_HOST_CREATE,
		PERMISSION_HOST_READ,
		PERMISSION_HOST_UPDATE,
		PERMISSION_HOST_DELETE,
		PERMISSION_HOST_SERVICE_CREATE,
		PERMISSION_HOST_SERVICE_READ,
		PERMISSION_HOST_SERVICE_UPDATE,
		PERMISSION_HOST_SERVICE_DELETE,
		PERMISSION_PROPERTY_CREATE,
		PERMISSION_PROPERTY_READ,
		PERMISSION_PROPERTY_UPDATE,
		PERMISSION_PROPERTY_DELETE,
		PERMISSION_SERVICE_CREATE,
		PERMISSION_SERVICE_READ,
		PERMISSION_SERVICE_UPDATE,
		PERMISSION_SERVICE_DELETE,
		PERMISSION_CHECK_CREATE,
		PERMISSION_CHECK_READ,
		PERMISSION_CHECK_UPDATE,
		PERMISSION_CHECK_DELETE,
		PERMISSION_ROUND_CREATE,
		PERMISSION_ROUND_READ,
		PERMISSION_ROUND_UPDATE,
		PERMISSION_ROUND_DELETE,
	} {
		kinds = append(kinds, string(s))
	}
	return
}
