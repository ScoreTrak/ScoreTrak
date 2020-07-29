package property

type Serv interface {
	Delete(id uint32) error
	GetAll() ([]*Property, error)
	GetByID(id uint32) (*Property, error)
	Store(u *Property) error
	Update(u *Property) error
}

type propertyServ struct {
	repo Repo
}

func NewPropertyServ(repo Repo) Serv {
	return &propertyServ{
		repo: repo,
	}
}

func (svc *propertyServ) Delete(id uint32) error { return svc.repo.Delete(id) }

func (svc *propertyServ) GetAll() ([]*Property, error) { return svc.repo.GetAll() }

func (svc *propertyServ) GetByID(id uint32) (*Property, error) { return svc.repo.GetByID(id) }

func (svc *propertyServ) Store(u *Property) error { return svc.repo.Store(u) }

func (svc *propertyServ) Update(u *Property) error { return svc.repo.Update(u) }
