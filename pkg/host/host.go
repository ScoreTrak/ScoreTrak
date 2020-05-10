package host

// Host model represents a single machine. This could be an IP address or a resolvable hostname
type Host struct {
	Id int64 `json:"id,omitempty"`

	Address string `json:"address"`

	// The ID of a host group that the host belongs to.
	HostGroupId int64 `json:"host_group_id,omitempty"`

	// The ID of a team that this host belongs too. This parameter is optional, however is needed to appear on the scoring engine.
	TeamId int64 `json:"team_id,omitempty"`

	// Enables or disables scoring for a single host
	Enabled bool `json:"enabled,omitempty"`

	// Enables to Edit the hostname. If a single host needs to be eddited for one service, and kept only visible for other service, you can make 2 services that point to same address, and have different edit_host properties.
	EditHost bool `json:"edit_host,omitempty"`
}
