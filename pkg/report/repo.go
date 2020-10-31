package report

import "context"

type Repo interface {
	Get(ctx context.Context) (*Report, error)
	Update(context.Context, *Report) error
}
