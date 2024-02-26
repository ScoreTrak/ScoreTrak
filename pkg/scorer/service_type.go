package scorer

import (
	"context"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome_writer"
)

type ServiceType string

const (
	SERVICE_FTP    ServiceType = "ftp"
	SERVICE_SSH    ServiceType = "ssh"
	SERVICE_WINRM  ServiceType = "winrm"
	SERVICE_PING   ServiceType = "ping"
	SERVICE_HTTP   ServiceType = "http"
	SERVICE_LDAP   ServiceType = "ldap"
	SERVICE_DNS    ServiceType = "dns"
	SERVICE_SMB    ServiceType = "smb"
	SERVICE_IMAP   ServiceType = "imap"
	SERVICE_SQL    ServiceType = "sql"
	SERVICE_CALDAV ServiceType = "caldav"
)

func (ServiceType) Values() (kinds []string) {
	for _, s := range []ServiceType{
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

type Executor func(ctx context.Context, ow *outcome_writer.OutcomeWriter, properties []byte)
