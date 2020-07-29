package host

type Serv interface {
	Delete(id uint32) error
	GetAll() ([]*Host, error)
	GetByID(id uint32) (*Host, error)
	Store(u *Host) error
	Update(u *Host) error
}

type hostServ struct {
	repo Repo
}

func NewHostServ(repo Repo) Serv {
	return &hostServ{
		repo: repo,
	}
}

func (svc *hostServ) Delete(id uint32) error { return svc.repo.Delete(id) }

func (svc *hostServ) GetAll() ([]*Host, error) { return svc.repo.GetAll() }

func (svc *hostServ) GetByID(id uint32) (*Host, error) { return svc.repo.GetByID(id) }

func (svc *hostServ) Store(u *Host) error { return svc.repo.Store(u) }

func (svc *hostServ) Update(u *Host) error { return svc.repo.Update(u) }
