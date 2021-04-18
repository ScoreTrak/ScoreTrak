package policy

//Policy is a set of dynamic values that could be edited in runtime. Policy is typically propagated between scoretrak instances via pkg/policy/policy_client
type Policy struct {
	ID                                        uint  `json:"-" gorm:"primary_key;"`
	AllowUnauthenticatedUsers                 *bool `json:"allow_unauthenticated_users" gorm:"not null;default:true"`
	AllowChangingUsernamesAndPasswords        *bool `json:"allow_changing_usernames_and_passwords" gorm:"not null;default:false"`
	AllowRedTeamLaunchingServiceTestsManually *bool `json:"allow_red_team_launching_service_tests_manually" gorm:"not null;default:false"`
	ShowPoints                                *bool `json:"allow_to_see_points" gorm:"not null;default:true"`
	ShowAddresses                             *bool `json:"show_addresses" gorm:"not null;default:true"`
}
