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
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/predicate"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/round"
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

// SetUpdateTime sets the "update_time" field.
func (ru *RoundUpdate) SetUpdateTime(t time.Time) *RoundUpdate {
	ru.mutation.SetUpdateTime(t)
	return ru
}

// ClearUpdateTime clears the value of the "update_time" field.
func (ru *RoundUpdate) ClearUpdateTime() *RoundUpdate {
	ru.mutation.ClearUpdateTime()
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

// SetNumOfIntendedChecks sets the "num_of_intended_checks" field.
func (ru *RoundUpdate) SetNumOfIntendedChecks(i int) *RoundUpdate {
	ru.mutation.ResetNumOfIntendedChecks()
	ru.mutation.SetNumOfIntendedChecks(i)
	return ru
}

// SetNillableNumOfIntendedChecks sets the "num_of_intended_checks" field if the given value is not nil.
func (ru *RoundUpdate) SetNillableNumOfIntendedChecks(i *int) *RoundUpdate {
	if i != nil {
		ru.SetNumOfIntendedChecks(*i)
	}
	return ru
}

// AddNumOfIntendedChecks adds i to the "num_of_intended_checks" field.
func (ru *RoundUpdate) AddNumOfIntendedChecks(i int) *RoundUpdate {
	ru.mutation.AddNumOfIntendedChecks(i)
	return ru
}

// ClearNumOfIntendedChecks clears the value of the "num_of_intended_checks" field.
func (ru *RoundUpdate) ClearNumOfIntendedChecks() *RoundUpdate {
	ru.mutation.ClearNumOfIntendedChecks()
	return ru
}

// SetStatus sets the "status" field.
func (ru *RoundUpdate) SetStatus(r round.Status) *RoundUpdate {
	ru.mutation.SetStatus(r)
	return ru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ru *RoundUpdate) SetNillableStatus(r *round.Status) *RoundUpdate {
	if r != nil {
		ru.SetStatus(*r)
	}
	return ru
}

// SetStartedAt sets the "started_at" field.
func (ru *RoundUpdate) SetStartedAt(t time.Time) *RoundUpdate {
	ru.mutation.SetStartedAt(t)
	return ru
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (ru *RoundUpdate) SetNillableStartedAt(t *time.Time) *RoundUpdate {
	if t != nil {
		ru.SetStartedAt(*t)
	}
	return ru
}

// SetFinishedAt sets the "finished_at" field.
func (ru *RoundUpdate) SetFinishedAt(t time.Time) *RoundUpdate {
	ru.mutation.SetFinishedAt(t)
	return ru
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (ru *RoundUpdate) SetNillableFinishedAt(t *time.Time) *RoundUpdate {
	if t != nil {
		ru.SetFinishedAt(*t)
	}
	return ru
}

// ClearFinishedAt clears the value of the "finished_at" field.
func (ru *RoundUpdate) ClearFinishedAt() *RoundUpdate {
	ru.mutation.ClearFinishedAt()
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
	ru.defaults()
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

// defaults sets the default values of the builder before save.
func (ru *RoundUpdate) defaults() {
	if _, ok := ru.mutation.UpdateTime(); !ok && !ru.mutation.UpdateTimeCleared() {
		v := round.UpdateDefaultUpdateTime()
		ru.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RoundUpdate) check() error {
	if v, ok := ru.mutation.RoundNumber(); ok {
		if err := round.RoundNumberValidator(v); err != nil {
			return &ValidationError{Name: "round_number", err: fmt.Errorf(`entities: validator failed for field "Round.round_number": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Status(); ok {
		if err := round.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`entities: validator failed for field "Round.status": %w`, err)}
		}
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
	if ru.mutation.CreateTimeCleared() {
		_spec.ClearField(round.FieldCreateTime, field.TypeTime)
	}
	if value, ok := ru.mutation.UpdateTime(); ok {
		_spec.SetField(round.FieldUpdateTime, field.TypeTime, value)
	}
	if ru.mutation.UpdateTimeCleared() {
		_spec.ClearField(round.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := ru.mutation.RoundNumber(); ok {
		_spec.SetField(round.FieldRoundNumber, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedRoundNumber(); ok {
		_spec.AddField(round.FieldRoundNumber, field.TypeInt, value)
	}
	if value, ok := ru.mutation.NumOfIntendedChecks(); ok {
		_spec.SetField(round.FieldNumOfIntendedChecks, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedNumOfIntendedChecks(); ok {
		_spec.AddField(round.FieldNumOfIntendedChecks, field.TypeInt, value)
	}
	if ru.mutation.NumOfIntendedChecksCleared() {
		_spec.ClearField(round.FieldNumOfIntendedChecks, field.TypeInt)
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.SetField(round.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ru.mutation.StartedAt(); ok {
		_spec.SetField(round.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.FinishedAt(); ok {
		_spec.SetField(round.FieldFinishedAt, field.TypeTime, value)
	}
	if ru.mutation.FinishedAtCleared() {
		_spec.ClearField(round.FieldFinishedAt, field.TypeTime)
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

// SetUpdateTime sets the "update_time" field.
func (ruo *RoundUpdateOne) SetUpdateTime(t time.Time) *RoundUpdateOne {
	ruo.mutation.SetUpdateTime(t)
	return ruo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (ruo *RoundUpdateOne) ClearUpdateTime() *RoundUpdateOne {
	ruo.mutation.ClearUpdateTime()
	return ruo
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

// SetNumOfIntendedChecks sets the "num_of_intended_checks" field.
func (ruo *RoundUpdateOne) SetNumOfIntendedChecks(i int) *RoundUpdateOne {
	ruo.mutation.ResetNumOfIntendedChecks()
	ruo.mutation.SetNumOfIntendedChecks(i)
	return ruo
}

// SetNillableNumOfIntendedChecks sets the "num_of_intended_checks" field if the given value is not nil.
func (ruo *RoundUpdateOne) SetNillableNumOfIntendedChecks(i *int) *RoundUpdateOne {
	if i != nil {
		ruo.SetNumOfIntendedChecks(*i)
	}
	return ruo
}

// AddNumOfIntendedChecks adds i to the "num_of_intended_checks" field.
func (ruo *RoundUpdateOne) AddNumOfIntendedChecks(i int) *RoundUpdateOne {
	ruo.mutation.AddNumOfIntendedChecks(i)
	return ruo
}

// ClearNumOfIntendedChecks clears the value of the "num_of_intended_checks" field.
func (ruo *RoundUpdateOne) ClearNumOfIntendedChecks() *RoundUpdateOne {
	ruo.mutation.ClearNumOfIntendedChecks()
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *RoundUpdateOne) SetStatus(r round.Status) *RoundUpdateOne {
	ruo.mutation.SetStatus(r)
	return ruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ruo *RoundUpdateOne) SetNillableStatus(r *round.Status) *RoundUpdateOne {
	if r != nil {
		ruo.SetStatus(*r)
	}
	return ruo
}

// SetStartedAt sets the "started_at" field.
func (ruo *RoundUpdateOne) SetStartedAt(t time.Time) *RoundUpdateOne {
	ruo.mutation.SetStartedAt(t)
	return ruo
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (ruo *RoundUpdateOne) SetNillableStartedAt(t *time.Time) *RoundUpdateOne {
	if t != nil {
		ruo.SetStartedAt(*t)
	}
	return ruo
}

// SetFinishedAt sets the "finished_at" field.
func (ruo *RoundUpdateOne) SetFinishedAt(t time.Time) *RoundUpdateOne {
	ruo.mutation.SetFinishedAt(t)
	return ruo
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (ruo *RoundUpdateOne) SetNillableFinishedAt(t *time.Time) *RoundUpdateOne {
	if t != nil {
		ruo.SetFinishedAt(*t)
	}
	return ruo
}

// ClearFinishedAt clears the value of the "finished_at" field.
func (ruo *RoundUpdateOne) ClearFinishedAt() *RoundUpdateOne {
	ruo.mutation.ClearFinishedAt()
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
	ruo.defaults()
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

// defaults sets the default values of the builder before save.
func (ruo *RoundUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateTime(); !ok && !ruo.mutation.UpdateTimeCleared() {
		v := round.UpdateDefaultUpdateTime()
		ruo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RoundUpdateOne) check() error {
	if v, ok := ruo.mutation.RoundNumber(); ok {
		if err := round.RoundNumberValidator(v); err != nil {
			return &ValidationError{Name: "round_number", err: fmt.Errorf(`entities: validator failed for field "Round.round_number": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Status(); ok {
		if err := round.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`entities: validator failed for field "Round.status": %w`, err)}
		}
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
	if ruo.mutation.CreateTimeCleared() {
		_spec.ClearField(round.FieldCreateTime, field.TypeTime)
	}
	if value, ok := ruo.mutation.UpdateTime(); ok {
		_spec.SetField(round.FieldUpdateTime, field.TypeTime, value)
	}
	if ruo.mutation.UpdateTimeCleared() {
		_spec.ClearField(round.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := ruo.mutation.RoundNumber(); ok {
		_spec.SetField(round.FieldRoundNumber, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedRoundNumber(); ok {
		_spec.AddField(round.FieldRoundNumber, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.NumOfIntendedChecks(); ok {
		_spec.SetField(round.FieldNumOfIntendedChecks, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedNumOfIntendedChecks(); ok {
		_spec.AddField(round.FieldNumOfIntendedChecks, field.TypeInt, value)
	}
	if ruo.mutation.NumOfIntendedChecksCleared() {
		_spec.ClearField(round.FieldNumOfIntendedChecks, field.TypeInt)
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.SetField(round.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ruo.mutation.StartedAt(); ok {
		_spec.SetField(round.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.FinishedAt(); ok {
		_spec.SetField(round.FieldFinishedAt, field.TypeTime, value)
	}
	if ruo.mutation.FinishedAtCleared() {
		_spec.ClearField(round.FieldFinishedAt, field.TypeTime)
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
