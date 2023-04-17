// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/check"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/predicate"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/round"
)

// RoundUpdate is the builder for updating Round entities.
type RoundUpdate struct {
	config
	hooks    []Hook
	mutation *RoundMutation
}

// Where appends a list predicates to the RoundUpdate builder.
func (ru *RoundUpdate) Where(ps ...predicate.Round) *RoundUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetRoundNumber sets the "round_number" field.
func (ru *RoundUpdate) SetRoundNumber(i int) *RoundUpdate {
	ru.mutation.ResetRoundNumber()
	ru.mutation.SetRoundNumber(i)
	return ru
}

// AddRoundNumber adds i to the "round_number" field.
func (ru *RoundUpdate) AddRoundNumber(i int) *RoundUpdate {
	ru.mutation.AddRoundNumber(i)
	return ru
}

// SetNote sets the "note" field.
func (ru *RoundUpdate) SetNote(s string) *RoundUpdate {
	ru.mutation.SetNote(s)
	return ru
}

// SetErr sets the "err" field.
func (ru *RoundUpdate) SetErr(s string) *RoundUpdate {
	ru.mutation.SetErr(s)
	return ru
}

// SetStartedAt sets the "started_at" field.
func (ru *RoundUpdate) SetStartedAt(t time.Time) *RoundUpdate {
	ru.mutation.SetStartedAt(t)
	return ru
}

// SetFinishedAt sets the "finished_at" field.
func (ru *RoundUpdate) SetFinishedAt(t time.Time) *RoundUpdate {
	ru.mutation.SetFinishedAt(t)
	return ru
}

// AddCheckIDs adds the "checks" edge to the Check entity by IDs.
func (ru *RoundUpdate) AddCheckIDs(ids ...string) *RoundUpdate {
	ru.mutation.AddCheckIDs(ids...)
	return ru
}

// AddChecks adds the "checks" edges to the Check entity.
func (ru *RoundUpdate) AddChecks(c ...*Check) *RoundUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.AddCheckIDs(ids...)
}

// Mutation returns the RoundMutation object of the builder.
func (ru *RoundUpdate) Mutation() *RoundMutation {
	return ru.mutation
}

// ClearChecks clears all "checks" edges to the Check entity.
func (ru *RoundUpdate) ClearChecks() *RoundUpdate {
	ru.mutation.ClearChecks()
	return ru
}

// RemoveCheckIDs removes the "checks" edge to Check entities by IDs.
func (ru *RoundUpdate) RemoveCheckIDs(ids ...string) *RoundUpdate {
	ru.mutation.RemoveCheckIDs(ids...)
	return ru
}

// RemoveChecks removes "checks" edges to Check entities.
func (ru *RoundUpdate) RemoveChecks(c ...*Check) *RoundUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.RemoveCheckIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoundUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, RoundMutation](ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoundUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoundUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoundUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RoundUpdate) check() error {
	if _, ok := ru.mutation.CompetitionID(); ru.mutation.CompetitionCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Round.competition"`)
	}
	return nil
}

func (ru *RoundUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(round.Table, round.Columns, sqlgraph.NewFieldSpec(round.FieldID, field.TypeString))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.RoundNumber(); ok {
		_spec.SetField(round.FieldRoundNumber, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedRoundNumber(); ok {
		_spec.AddField(round.FieldRoundNumber, field.TypeInt, value)
	}
	if value, ok := ru.mutation.Note(); ok {
		_spec.SetField(round.FieldNote, field.TypeString, value)
	}
	if value, ok := ru.mutation.Err(); ok {
		_spec.SetField(round.FieldErr, field.TypeString, value)
	}
	if value, ok := ru.mutation.StartedAt(); ok {
		_spec.SetField(round.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.FinishedAt(); ok {
		_spec.SetField(round.FieldFinishedAt, field.TypeTime, value)
	}
	if ru.mutation.ChecksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   round.ChecksTable,
			Columns: []string{round.ChecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(check.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedChecksIDs(); len(nodes) > 0 && !ru.mutation.ChecksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   round.ChecksTable,
			Columns: []string{round.ChecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(check.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ChecksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   round.ChecksTable,
			Columns: []string{round.ChecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(check.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{round.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoundUpdateOne is the builder for updating a single Round entity.
type RoundUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoundMutation
}

// SetRoundNumber sets the "round_number" field.
func (ruo *RoundUpdateOne) SetRoundNumber(i int) *RoundUpdateOne {
	ruo.mutation.ResetRoundNumber()
	ruo.mutation.SetRoundNumber(i)
	return ruo
}

// AddRoundNumber adds i to the "round_number" field.
func (ruo *RoundUpdateOne) AddRoundNumber(i int) *RoundUpdateOne {
	ruo.mutation.AddRoundNumber(i)
	return ruo
}

// SetNote sets the "note" field.
func (ruo *RoundUpdateOne) SetNote(s string) *RoundUpdateOne {
	ruo.mutation.SetNote(s)
	return ruo
}

// SetErr sets the "err" field.
func (ruo *RoundUpdateOne) SetErr(s string) *RoundUpdateOne {
	ruo.mutation.SetErr(s)
	return ruo
}

// SetStartedAt sets the "started_at" field.
func (ruo *RoundUpdateOne) SetStartedAt(t time.Time) *RoundUpdateOne {
	ruo.mutation.SetStartedAt(t)
	return ruo
}

// SetFinishedAt sets the "finished_at" field.
func (ruo *RoundUpdateOne) SetFinishedAt(t time.Time) *RoundUpdateOne {
	ruo.mutation.SetFinishedAt(t)
	return ruo
}

// AddCheckIDs adds the "checks" edge to the Check entity by IDs.
func (ruo *RoundUpdateOne) AddCheckIDs(ids ...string) *RoundUpdateOne {
	ruo.mutation.AddCheckIDs(ids...)
	return ruo
}

// AddChecks adds the "checks" edges to the Check entity.
func (ruo *RoundUpdateOne) AddChecks(c ...*Check) *RoundUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.AddCheckIDs(ids...)
}

// Mutation returns the RoundMutation object of the builder.
func (ruo *RoundUpdateOne) Mutation() *RoundMutation {
	return ruo.mutation
}

// ClearChecks clears all "checks" edges to the Check entity.
func (ruo *RoundUpdateOne) ClearChecks() *RoundUpdateOne {
	ruo.mutation.ClearChecks()
	return ruo
}

// RemoveCheckIDs removes the "checks" edge to Check entities by IDs.
func (ruo *RoundUpdateOne) RemoveCheckIDs(ids ...string) *RoundUpdateOne {
	ruo.mutation.RemoveCheckIDs(ids...)
	return ruo
}

// RemoveChecks removes "checks" edges to Check entities.
func (ruo *RoundUpdateOne) RemoveChecks(c ...*Check) *RoundUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.RemoveCheckIDs(ids...)
}

// Where appends a list predicates to the RoundUpdate builder.
func (ruo *RoundUpdateOne) Where(ps ...predicate.Round) *RoundUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoundUpdateOne) Select(field string, fields ...string) *RoundUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Round entity.
func (ruo *RoundUpdateOne) Save(ctx context.Context) (*Round, error) {
	return withHooks[*Round, RoundMutation](ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoundUpdateOne) SaveX(ctx context.Context) *Round {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoundUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoundUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RoundUpdateOne) check() error {
	if _, ok := ruo.mutation.CompetitionID(); ruo.mutation.CompetitionCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Round.competition"`)
	}
	return nil
}

func (ruo *RoundUpdateOne) sqlSave(ctx context.Context) (_node *Round, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(round.Table, round.Columns, sqlgraph.NewFieldSpec(round.FieldID, field.TypeString))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "Round.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, round.FieldID)
		for _, f := range fields {
			if !round.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != round.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.RoundNumber(); ok {
		_spec.SetField(round.FieldRoundNumber, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedRoundNumber(); ok {
		_spec.AddField(round.FieldRoundNumber, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.Note(); ok {
		_spec.SetField(round.FieldNote, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Err(); ok {
		_spec.SetField(round.FieldErr, field.TypeString, value)
	}
	if value, ok := ruo.mutation.StartedAt(); ok {
		_spec.SetField(round.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.FinishedAt(); ok {
		_spec.SetField(round.FieldFinishedAt, field.TypeTime, value)
	}
	if ruo.mutation.ChecksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   round.ChecksTable,
			Columns: []string{round.ChecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(check.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedChecksIDs(); len(nodes) > 0 && !ruo.mutation.ChecksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   round.ChecksTable,
			Columns: []string{round.ChecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(check.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ChecksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   round.ChecksTable,
			Columns: []string{round.ChecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(check.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Round{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{round.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
