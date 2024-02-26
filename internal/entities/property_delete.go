// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/scoretrak/scoretrak/internal/entities/predicate"
	"github.com/scoretrak/scoretrak/internal/entities/property"
)

// PropertyDelete is the builder for deleting a Property entity.
type PropertyDelete struct {
	config
	hooks    []Hook
	mutation *PropertyMutation
}

// Where appends a list predicates to the PropertyDelete builder.
func (pd *PropertyDelete) Where(ps ...predicate.Property) *PropertyDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PropertyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, PropertyMutation](ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PropertyDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PropertyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(property.Table, sqlgraph.NewFieldSpec(property.FieldID, field.TypeString))
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// PropertyDeleteOne is the builder for deleting a single Property entity.
type PropertyDeleteOne struct {
	pd *PropertyDelete
}

// Where appends a list predicates to the PropertyDelete builder.
func (pdo *PropertyDeleteOne) Where(ps ...predicate.Property) *PropertyDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *PropertyDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{property.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PropertyDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}
