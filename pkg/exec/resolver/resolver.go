package resolver

import (
	"strings"

	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec/services"
)

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
