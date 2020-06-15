package report

type Repo interface {
	Get() (*Report, error)
}
