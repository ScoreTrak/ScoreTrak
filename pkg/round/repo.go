package round

type Repo interface {
	GetLastRound() (*Round, error)
}
