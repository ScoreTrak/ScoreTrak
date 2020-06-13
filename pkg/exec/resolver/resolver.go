package resolver

import (
	"ScoreTrak/pkg/exec"
	"ScoreTrak/pkg/exec/services"
)

func ExecutableByName(s string) exec.Executable {
	if s == "FTP" {
		return &services.FTP{}
	} else {
		return nil
	}
}
