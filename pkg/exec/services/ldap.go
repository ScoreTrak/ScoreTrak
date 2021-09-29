package services

import (
	"crypto/tls"
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/go-ldap/ldap/v3"
	"net"
	"strings"
	"time"
)

type LDAP struct {
	Username            string
	Password            string
	Domain              string
	Port                string
	BaseDN              string
	TransportProtocol   string
	ApplicationProtocol string
	Filter              string
	Attributes          string
}

func NewLDAP() *LDAP {
	f := LDAP{TransportProtocol: "tcp", Filter: "(&(objectClass=organizationalPerson))", Attributes: "dn,cn"}
	return &f
}

func (l *LDAP) Validate() error {
	if l.Password != "" && l.Username != "" && l.Domain != "" {
		return nil
	}
	return errors.New("LDAP check_service needs password, username, and Domain to operate")
}

func (l *LDAP) Execute(e exec.Exec) (passed bool, log string, err error) {
	var ld *ldap.Conn
	if IsSecure(l.ApplicationProtocol) {
		if l.Port == "" {
			l.Port = "636"
		}
		c, err := tls.DialWithDialer(&net.Dialer{Deadline: e.Deadline()}, l.TransportProtocol, e.Host+":"+l.Port, &tls.Config{InsecureSkipVerify: true}) //https://github.com/golang/go/issues/39489
		if err != nil {
			return false, "Unable to dial remote ldap server", err
		}
		ld = ldap.NewConn(c, true)
		ld.Start()
	} else {
		if l.Port == "" {
			l.Port = "389"
		}
		c, err := net.DialTimeout(l.TransportProtocol, e.Host+":"+l.Port, time.Until(e.Deadline()))
		if err != nil {
			return false, "Unable to dial remote ldap server", err
		}
		ld = ldap.NewConn(c, false)
		ld.Start()
	}
	defer ld.Close()

	err = ld.Bind(l.Username+"@"+l.Domain, l.Password)
	if err != nil {
		return false, "Unable to bind", err
	}
	log = Success
	if l.BaseDN != "" {
		attributes := strings.Split(l.Attributes, ",")
		for i := range attributes {
			attributes[i] = strings.TrimSpace(attributes[i])
		}
		searchRequest := ldap.NewSearchRequest(
			l.BaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
			l.Filter, attributes, nil,
		)
		sr, err := ld.Search(searchRequest)
		if err != nil {
			return false, "Unable to search based on the parameters provided", err
		}
		for _, entry := range sr.Entries {
			log += "\n" + entry.DN
		}
	}
	return true, log, nil
}
