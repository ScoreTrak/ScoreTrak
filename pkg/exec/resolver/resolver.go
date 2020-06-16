package resolver

import (
	"ScoreTrak/pkg/exec"
	"ScoreTrak/pkg/exec/services"
	"strings"
)

func ExecutableByName(s string) exec.Executable {

	switch strings.ToLower(s) {
	case "FTP":
		return services.NewFTP()
	case "SSH":
		return services.NewSSH()
	case "WINRM":
		return services.NewWinrm()
	case "PING":
		services.NewPing()
	case "HTTP":
		services.NewHTTP()
	case "LDAP":
		services.NewLDAP()
	case "DNS":
		services.NewDNS()
	case "SMB":
		services.NewSMB()
	case "IMAP":
		services.NewIMAP()
	}

	return nil

}
