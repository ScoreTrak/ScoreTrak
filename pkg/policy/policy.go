package policy

type Policy struct {
	ID                                 uint  `json:"-" gorm:"primary_key;"`
	AllowUnauthenticatedUsers          *bool `json:"allow_unauthenticated_users" gorm:"not null;default:true"`
	AllowChangingUsernamesAndPasswords *bool `json:"allow_changing_usernames_and_passwords" gorm:"not null;default:false"`
	//AllowLaunchingServiceTestsManually *bool `json:"allow_launching_service_tests_manually" gorm:"not null;default: false"` //Todo: Implement this for API endpoint, and Scoretrak
	ShowPoints    *bool `json:"allow_to_see_points" gorm:"not null;default:true"`
	ShowAddresses *bool `json:"show_addresses" gorm:"not null;default:true"`
}
