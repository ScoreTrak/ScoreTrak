package resolver

import (
	"strings"

	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec/services"
)

type Service string

const (
	SERVICE_FTP    Service = "ftp"
	SERVICE_SSH    Service = "ssh"
	SERVICE_WINRM  Service = "winrm"
	SERVICE_PING   Service = "ping"
	SERVICE_HTTP   Service = "http"
	SERVICE_LDAP   Service = "ldap"
	SERVICE_DNS    Service = "dns"
	SERVICE_SMB    Service = "smb"
	SERVICE_IMAP   Service = "imap"
	SERVICE_SQL    Service = "sql"
	SERVICE_CALDAV Service = "caldav"
)

func (Service) Values() (kinds []string) {
	for _, s := range []Service{
		SERVICE_FTP,
		SERVICE_SSH,
		SERVICE_WINRM,
		SERVICE_PING,
		SERVICE_HTTP,
		SERVICE_LDAP,
		SERVICE_DNS,
		SERVICE_SMB,
		SERVICE_IMAP,
		SERVICE_SQL,
		SERVICE_CALDAV,
	} {
		kinds = append(kinds, string(s))
	}
	return
}

// ExecutableByName converts the name of the service to a specific executable type defined in pkg/exec/services
func ExecutableByName(s string) exec.Executable {
	switch strings.ToLower(s) {
	case "ftp":
		return services.NewFTP()
	case "ssh":
		return services.NewSSH()
	case "winrm":
		return services.NewWinrm()
	case "ping":
		return services.NewPing()
	case "http":
		return services.NewHTTP()
	case "ldap":
		return services.NewLDAP()
	case "dns":
		return services.NewDNS()
	case "smb":
		return services.NewSMB()
	case "imap":
		return services.NewIMAP()
	case "sql":
		return services.NewSQL()
	case "caldav":
		return services.NewCalDav()
	}
	return nil
}
