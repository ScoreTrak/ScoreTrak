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
	default:
		return nil
	}

}
