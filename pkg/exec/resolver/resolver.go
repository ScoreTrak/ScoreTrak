package resolver

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/exec"
	"github.com/L1ghtman2k/ScoreTrak/pkg/exec/services"
	"strings"
)

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
	}

	return nil

}
