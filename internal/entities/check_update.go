// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/check"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/predicate"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/round"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/service"
)

// CheckUpdate is the builder for updating Check entities.
type CheckUpdate struct {
	config
	hooks    []Hook
	mutation *CheckMutation
}

// Where appends a list predicates to the CheckUpdate builder.
func (cu *CheckUpdate) Where(ps ...predicate.Check) *CheckUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetPause sets the "pause" field.
func (cu *CheckUpdate) SetPause(b bool) *CheckUpdate {
	cu.mutation.SetPause(b)
	return cu
}

// SetNillablePause sets the "pause" field if the given value is not nil.
func (cu *CheckUpdate) SetNillablePause(b *bool) *CheckUpdate {
	if b != nil {
		cu.SetPause(*b)
	}
	return cu
}

// ClearPause clears the value of the "pause" field.
func (cu *CheckUpdate) ClearPause() *CheckUpdate {
	cu.mutation.ClearPause()
	return cu
}

// SetHidden sets the "hidden" field.
func (cu *CheckUpdate) SetHidden(b bool) *CheckUpdate {
	cu.mutation.SetHidden(b)
	return cu
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (cu *CheckUpdate) SetNillableHidden(b *bool) *CheckUpdate {
	if b != nil {
		cu.SetHidden(*b)
	}
	return cu
}

// ClearHidden clears the value of the "hidden" field.
func (cu *CheckUpdate) ClearHidden() *CheckUpdate {
	cu.mutation.ClearHidden()
	return cu
}

// SetLog sets the "log" field.
func (cu *CheckUpdate) SetLog(s string) *CheckUpdate {
	cu.mutation.SetLog(s)
	return cu
}

// SetError sets the "error" field.
func (cu *CheckUpdate) SetError(s string) *CheckUpdate {
	cu.mutation.SetError(s)
	return cu
}

// SetPassed sets the "passed" field.
func (cu *CheckUpdate) SetPassed(b bool) *CheckUpdate {
	cu.mutation.SetPassed(b)
	return cu
}

// SetRoundsID sets the "rounds" edge to the Round entity by ID.
func (cu *CheckUpdate) SetRoundsID(id string) *CheckUpdate {
	cu.mutation.SetRoundsID(id)
	return cu
}

// SetRounds sets the "rounds" edge to the Round entity.
func (cu *CheckUpdate) SetRounds(r *Round) *CheckUpdate {
	return cu.SetRoundsID(r.ID)
}

// SetServicesID sets the "services" edge to the Service entity by ID.
func (cu *CheckUpdate) SetServicesID(id string) *CheckUpdate {
	cu.mutation.SetServicesID(id)
	return cu
}

// SetServices sets the "services" edge to the Service entity.
func (cu *CheckUpdate) SetServices(s *Service) *CheckUpdate {
	return cu.SetServicesID(s.ID)
}

// Mutation returns the CheckMutation object of the builder.
func (cu *CheckUpdate) Mutation() *CheckMutation {
	return cu.mutation
}

// ClearRounds clears the "rounds" edge to the Round entity.
func (cu *CheckUpdate) ClearRounds() *CheckUpdate {
	cu.mutation.ClearRounds()
	return cu
}

// ClearServices clears the "services" edge to the Service entity.
func (cu *CheckUpdate) ClearServices() *CheckUpdate {
	cu.mutation.ClearServices()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CheckUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CheckMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CheckUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CheckUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CheckUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CheckUpdate) check() error {
	if _, ok := cu.mutation.CompetitionID(); cu.mutation.CompetitionCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Check.competition"`)
	}
	if _, ok := cu.mutation.RoundsID(); cu.mutation.RoundsCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Check.rounds"`)
	}
	if _, ok := cu.mutation.ServicesID(); cu.mutation.ServicesCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Check.services"`)
	}
	return nil
}

func (cu *CheckUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(check.Table, check.Columns, sqlgraph.NewFieldSpec(check.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Pause(); ok {
		_spec.SetField(check.FieldPause, field.TypeBool, value)
	}
	if cu.mutation.PauseCleared() {
		_spec.ClearField(check.FieldPause, field.TypeBool)
	}
	if value, ok := cu.mutation.Hidden(); ok {
		_spec.SetField(check.FieldHidden, field.TypeBool, value)
	}
	if cu.mutation.HiddenCleared() {
		_spec.ClearField(check.FieldHidden, field.TypeBool)
	}
	if value, ok := cu.mutation.Log(); ok {
		_spec.SetField(check.FieldLog, field.TypeString, value)
	}
	if value, ok := cu.mutation.Error(); ok {
		_spec.SetField(check.FieldError, field.TypeString, value)
	}
	if value, ok := cu.mutation.Passed(); ok {
		_spec.SetField(check.FieldPassed, field.TypeBool, value)
	}
	if cu.mutation.RoundsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   check.RoundsTable,
			Columns: []string{check.RoundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RoundsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   check.RoundsTable,
			Columns: []string{check.RoundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   check.ServicesTable,
			Columns: []string{check.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   check.ServicesTable,
			Columns: []string{check.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{check.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CheckUpdateOne is the builder for updating a single Check entity.
type CheckUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CheckMutation
}

// SetPause sets the "pause" field.
func (cuo *CheckUpdateOne) SetPause(b bool) *CheckUpdateOne {
	cuo.mutation.SetPause(b)
	return cuo
}

// SetNillablePause sets the "pause" field if the given value is not nil.
func (cuo *CheckUpdateOne) SetNillablePause(b *bool) *CheckUpdateOne {
	if b != nil {
		cuo.SetPause(*b)
	}
	return cuo
}

// ClearPause clears the value of the "pause" field.
func (cuo *CheckUpdateOne) ClearPause() *CheckUpdateOne {
	cuo.mutation.ClearPause()
	return cuo
}

// SetHidden sets the "hidden" field.
func (cuo *CheckUpdateOne) SetHidden(b bool) *CheckUpdateOne {
	cuo.mutation.SetHidden(b)
	return cuo
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (cuo *CheckUpdateOne) SetNillableHidden(b *bool) *CheckUpdateOne {
	if b != nil {
		cuo.SetHidden(*b)
	}
	return cuo
}

// ClearHidden clears the value of the "hidden" field.
func (cuo *CheckUpdateOne) ClearHidden() *CheckUpdateOne {
	cuo.mutation.ClearHidden()
	return cuo
}

// SetLog sets the "log" field.
func (cuo *CheckUpdateOne) SetLog(s string) *CheckUpdateOne {
	cuo.mutation.SetLog(s)
	return cuo
}

// SetError sets the "error" field.
func (cuo *CheckUpdateOne) SetError(s string) *CheckUpdateOne {
	cuo.mutation.SetError(s)
	return cuo
}

// SetPassed sets the "passed" field.
func (cuo *CheckUpdateOne) SetPassed(b bool) *CheckUpdateOne {
	cuo.mutation.SetPassed(b)
	return cuo
}

// SetRoundsID sets the "rounds" edge to the Round entity by ID.
func (cuo *CheckUpdateOne) SetRoundsID(id string) *CheckUpdateOne {
	cuo.mutation.SetRoundsID(id)
	return cuo
}

// SetRounds sets the "rounds" edge to the Round entity.
func (cuo *CheckUpdateOne) SetRounds(r *Round) *CheckUpdateOne {
	return cuo.SetRoundsID(r.ID)
}

// SetServicesID sets the "services" edge to the Service entity by ID.
func (cuo *CheckUpdateOne) SetServicesID(id string) *CheckUpdateOne {
	cuo.mutation.SetServicesID(id)
	return cuo
}

// SetServices sets the "services" edge to the Service entity.
func (cuo *CheckUpdateOne) SetServices(s *Service) *CheckUpdateOne {
	return cuo.SetServicesID(s.ID)
}

// Mutation returns the CheckMutation object of the builder.
func (cuo *CheckUpdateOne) Mutation() *CheckMutation {
	return cuo.mutation
}

// ClearRounds clears the "rounds" edge to the Round entity.
func (cuo *CheckUpdateOne) ClearRounds() *CheckUpdateOne {
	cuo.mutation.ClearRounds()
	return cuo
}

// ClearServices clears the "services" edge to the Service entity.
func (cuo *CheckUpdateOne) ClearServices() *CheckUpdateOne {
	cuo.mutation.ClearServices()
	return cuo
}

// Where appends a list predicates to the CheckUpdate builder.
func (cuo *CheckUpdateOne) Where(ps ...predicate.Check) *CheckUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CheckUpdateOne) Select(field string, fields ...string) *CheckUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Check entity.
func (cuo *CheckUpdateOne) Save(ctx context.Context) (*Check, error) {
	return withHooks[*Check, CheckMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CheckUpdateOne) SaveX(ctx context.Context) *Check {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CheckUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CheckUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CheckUpdateOne) check() error {
	if _, ok := cuo.mutation.CompetitionID(); cuo.mutation.CompetitionCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Check.competition"`)
	}
	if _, ok := cuo.mutation.RoundsID(); cuo.mutation.RoundsCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Check.rounds"`)
	}
	if _, ok := cuo.mutation.ServicesID(); cuo.mutation.ServicesCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Check.services"`)
	}
	return nil
}

func (cuo *CheckUpdateOne) sqlSave(ctx context.Context) (_node *Check, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(check.Table, check.Columns, sqlgraph.NewFieldSpec(check.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "Check.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, check.FieldID)
		for _, f := range fields {
			if !check.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != check.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Pause(); ok {
		_spec.SetField(check.FieldPause, field.TypeBool, value)
	}
	if cuo.mutation.PauseCleared() {
		_spec.ClearField(check.FieldPause, field.TypeBool)
	}
	if value, ok := cuo.mutation.Hidden(); ok {
		_spec.SetField(check.FieldHidden, field.TypeBool, value)
	}
	if cuo.mutation.HiddenCleared() {
		_spec.ClearField(check.FieldHidden, field.TypeBool)
	}
	if value, ok := cuo.mutation.Log(); ok {
		_spec.SetField(check.FieldLog, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Error(); ok {
		_spec.SetField(check.FieldError, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Passed(); ok {
		_spec.SetField(check.FieldPassed, field.TypeBool, value)
	}
	if cuo.mutation.RoundsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   check.RoundsTable,
			Columns: []string{check.RoundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RoundsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   check.RoundsTable,
			Columns: []string{check.RoundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   check.ServicesTable,
			Columns: []string{check.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   check.ServicesTable,
			Columns: []string{check.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Check{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{check.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
