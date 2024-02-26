package executors

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome_writer"
	"net"
)

type LDAPProperties struct {
	Username            string `json:"username" validate:"required"`
	Password            string `json:"password" validate:"required"`
	Host                string
	Domain              string `json:"domain" validate:"required"`
	Port                string
	BaseDN              string
	TransportProtocol   string
	ApplicationProtocol string
	Filter              string
	Attributes          string
}

var ErrLDAPRequiresUsernamePasswordDomain = errors.New("LDAP check_service needs password, username, and Domain to operate")

func ScoreLdap(ctx context.Context, ow *outcome_writer.OutcomeWriter, properties []byte) {
	ldapproperties := &LDAPProperties{}
	err := json.Unmarshal(properties, &ldapproperties)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to unmarshall properties: %v", err))
		return
	}

	err = validate.Struct(ldapproperties)
	if err != nil {
		// validationErrors := err.(validator.ValidationErrors)
		ow.SetError(fmt.Errorf("validation error: %w", err))
		return
	}

	address := fmt.Sprintf("%s:%s", ldapproperties.Host, ldapproperties.Port)
	var ld *ldap.Conn
	if IsSecure(ldapproperties.ApplicationProtocol) {
		if ldapproperties.Port == "" {
			ldapproperties.Port = "636"
		}
		c, err := tls.Dial(ldapproperties.TransportProtocol, address, &tls.Config{InsecureSkipVerify: true}) //nolint:gosec //https://github.com/golang/go/issues/39489
		if err != nil {
			ow.SetError(fmt.Errorf("unable to dial remote ldap server: %w", err))
			return
		}
		ld = ldap.NewConn(c, true)
		ld.Start()
	} else {
		if ldapproperties.Port == "" {
			ldapproperties.Port = "389"
		}
		c, err := net.Dial(ldapproperties.TransportProtocol, address)
		if err != nil {
			ow.SetError(fmt.Errorf("unable to dial remote ldap server: %w", err))
			return
		}
		ld = ldap.NewConn(c, false)
		ld.Start()
	}
	defer ld.Close()

	err = ld.Bind(ldapproperties.Username+"@"+ldapproperties.Domain, ldapproperties.Password)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to bind: %w", err))
		return
	}

	//logOutput = Success
	//if ldapproperties.BaseDN != "" {
	//	attributes := strings.Split(ldapproperties.Attributes, ",")
	//	for i := range attributes {
	//		attributes[i] = strings.TrimSpace(attributes[i])
	//	}
	//	searchRequest := ldap.NewSearchRequest(
	//		ldapproperties.BaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
	//		ldapproperties.Filter, attributes, nil,
	//	)
	//	sr, err := ld.Search(searchRequest)
	//	if err != nil {
	//		return false, "", fmt.Errorf("unable to search based on the parameters provided: %w", err)
	//	}
	//	for _, entry := range sr.Entries {
	//		logOutput += "\n" + entry.DN
	//	}
	//}
	ow.SetStatus(outcome.OUTCOME_STATUS_PASSED)
}
