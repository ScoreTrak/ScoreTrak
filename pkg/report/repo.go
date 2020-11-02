package report

import "github.com/gofrs/uuid"

type Repo interface {
	Get() (*Report, error)
	Update(*Report) error
	CountPassedPerService() (map[uuid.UUID]uint64, error)
}
