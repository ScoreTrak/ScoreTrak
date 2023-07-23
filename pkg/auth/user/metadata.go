package user

type TeamID string
type RoleID string

type RoleAssignment struct {
	Roles []RoleID `json:"roles"`
}

type MetadataPublic struct {
	IsAdmin             bool                      `json:"is_admin"`
	Teams               []TeamID                  `json:"teams"`
	TeamRoleAssignments map[TeamID]RoleAssignment `json:"team_role_assignments"`
}
