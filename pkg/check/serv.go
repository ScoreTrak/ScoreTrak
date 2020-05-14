package check

type Serv interface {
	GetAllByTeamRoundID(t_id string, r_id uint64) ([]*Check, error)
	GetByTeamRoundServiceID(t_id string, r_id uint64, s_id uint64) (*Check, error)
}

type checkServ struct {
	repo Repo
}

func NewCheckServ(repo Repo) Serv {
	return &checkServ{
		repo: repo,
	}
}

func (svc *checkServ) GetAllByTeamRoundID(t_id string, r_id uint64) ([]*Check, error) {
	return svc.repo.GetAllByTeamRoundID(t_id, r_id)
}
func (svc *checkServ) GetByTeamRoundServiceID(t_id string, r_id uint64, s_id uint64) (*Check, error) {
	return svc.repo.GetByTeamRoundServiceID(t_id, r_id, s_id)
}
