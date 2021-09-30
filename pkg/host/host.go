package host

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/asaskevich/govalidator"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Host model represents a single machine. This could be an IP address or a resolvable hostname
type Host struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`

	// Address represents the hostname or an IP address of the remote computer
	Address string `json:"address" gorm:"not null;default:null"`

	// Comma Separated List of allowed CIDRs, and hostnames
	AddressListRange *string `json:"address_list_range" gorm:"not null;default:''"`

	// The ID of a host group that the host belongs to.
	HostGroupID *uuid.UUID `json:"host_group_id,omitempty" gorm:"type:uuid"`

	// The ID of a team that this host belongs too.
	TeamID uuid.UUID `json:"team_id,omitempty" gorm:"type:uuid;not null"`

	// Hide is responsible for hiding the host on the scoring table
	Hide *bool `json:"pause,omitempty" gorm:"not null;default:false"`
	// Pause is responsible for pausing the host on scoring table
	Pause *bool `json:"hide,omitempty" gorm:"not null;default:false"`

	// Enables to Edit the hostname. If a single host needs to be eddited for one check_service, and kept only visible for other check_service, you can make 2 services that point to same address, and have different edit_host properties.
	EditHost *bool `json:"edit_host,omitempty" gorm:"not null;default:false"`

	// Services is an aggregate of all child Services that belong to a host
	Services []*service.Service `json:"services,omitempty" gorm:"foreignkey:HostID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

var ErrInvalidHostNameOrIPAddress = errors.New("invalid Hostname, or IP address")

func (p *Host) BeforeSave(tx *gorm.DB) (err error) {
	if p.AddressListRange != nil || p.Address != "" {
		p.Address = strings.ReplaceAll(p.Address, " ", "")
		if (p.AddressListRange == nil || p.Address == "") && p.ID != uuid.Nil {
			hst := &Host{}
			err := tx.Where("id = ?", p.ID).First(hst).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("unable to retrieve the requested entry, in order to validate address: %w", err)
			}
			if p.AddressListRange == nil {
				p.AddressListRange = hst.AddressListRange
			}
			if p.Address == "" {
				p.Address = hst.Address
			}
		}

		if !govalidator.IsHost(p.Address) {
			return fmt.Errorf("%w: %s", ErrInvalidHostNameOrIPAddress, p.Address)
		}
		if p.AddressListRange != nil {
			*p.AddressListRange = strings.ReplaceAll(*p.AddressListRange, " ", "")
			err = validateIfAddressInRange(p.Address, *p.AddressListRange)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Host) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		u, err := uuid.NewV4()
		if err != nil {
			return err
		}
		p.ID = u
	}
	return nil
}

var ErrAddressNotInRange = errors.New("the provided ip address was not in allowed range")
var ErrNotAValidHostname = errors.New("not a valid hostname")

func validateIfAddressInRange(addr string, addresses string) (err error) {
	if addresses == "" {
		return nil
	}
	addressList := strings.Split(addresses, ",")
	for i := range addressList {
		switch _, network, err := net.ParseCIDR(addressList[i]); {
		case err == nil:
			if network.Contains(net.ParseIP(addr)) {
				return nil
			}
		case govalidator.IsHost(addressList[i]):
			if strings.EqualFold(addressList[i], addr) {
				return nil
			}
		default:
			return fmt.Errorf("%w: %s", ErrNotAValidHostname, addressList[i])
		}
	}
	return ErrAddressNotInRange
}
