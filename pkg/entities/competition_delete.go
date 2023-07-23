// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/competition"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/predicate"
)

// CompetitionDelete is the builder for deleting a Competition entity.
type CompetitionDelete struct {
	config
	hooks    []Hook
	mutation *CompetitionMutation
}

// Where appends a list predicates to the CompetitionDelete builder.
func (cd *CompetitionDelete) Where(ps ...predicate.Competition) *CompetitionDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CompetitionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, CompetitionMutation](ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CompetitionDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CompetitionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(competition.Table, sqlgraph.NewFieldSpec(competition.FieldID, field.TypeString))
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// CompetitionDeleteOne is the builder for deleting a single Competition entity.
type CompetitionDeleteOne struct {
	cd *CompetitionDelete
}

// Where appends a list predicates to the CompetitionDelete builder.
func (cdo *CompetitionDeleteOne) Where(ps ...predicate.Competition) *CompetitionDeleteOne {
	cdo.cd.mutation.Where(ps...)
	return cdo
}

// Exec executes the deletion query.
func (cdo *CompetitionDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{competition.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CompetitionDeleteOne) ExecX(ctx context.Context) {
	if err := cdo.Exec(ctx); err != nil {
		panic(err)
	}
}
