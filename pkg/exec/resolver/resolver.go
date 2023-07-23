package resolver

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec/services"
	"github.com/ScoreTrak/ScoreTrak/pkg/scorer/scorerservice"
)

// ExecutableByName converts the name of the service to a specific executable type defined in pkg/exec/services
func ExecutableByName(s scorerservice.Service) exec.Executable {
	switch s {
	case scorerservice.SERVICE_FTP:
		return services.NewFTP()
	case scorerservice.SERVICE_SSH:
		return services.NewSSH()
	case scorerservice.SERVICE_WINRM:
		return services.NewWinrm()
	case scorerservice.SERVICE_PING:
		return services.NewPing()
	case scorerservice.SERVICE_HTTP:
		return services.NewHTTP()
	case scorerservice.SERVICE_LDAP:
		return services.NewLDAP()
	case scorerservice.SERVICE_DNS:
		return services.NewDNS()
	case scorerservice.SERVICE_SMB:
		return services.NewSMB()
	case scorerservice.SERVICE_IMAP:
		return services.NewIMAP()
	case scorerservice.SERVICE_SQL:
		return services.NewSQL()
	case scorerservice.SERVICE_CALDAV:
		return services.NewCalDav()
	}
	return nil
}
