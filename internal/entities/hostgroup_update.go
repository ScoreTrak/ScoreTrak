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
	"github.com/ScoreTrak/ScoreTrak/internal/entities/host"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/hostgroup"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/predicate"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/team"
)

// HostGroupUpdate is the builder for updating HostGroup entities.
type HostGroupUpdate struct {
	config
	hooks    []Hook
	mutation *HostGroupMutation
}

// Where appends a list predicates to the HostGroupUpdate builder.
func (hgu *HostGroupUpdate) Where(ps ...predicate.HostGroup) *HostGroupUpdate {
	hgu.mutation.Where(ps...)
	return hgu
}

// SetUpdateTime sets the "update_time" field.
func (hgu *HostGroupUpdate) SetUpdateTime(t time.Time) *HostGroupUpdate {
	hgu.mutation.SetUpdateTime(t)
	return hgu
}

// SetPause sets the "pause" field.
func (hgu *HostGroupUpdate) SetPause(b bool) *HostGroupUpdate {
	hgu.mutation.SetPause(b)
	return hgu
}

// SetHidden sets the "hidden" field.
func (hgu *HostGroupUpdate) SetHidden(b bool) *HostGroupUpdate {
	hgu.mutation.SetHidden(b)
	return hgu
}

// SetTeamID sets the "team_id" field.
func (hgu *HostGroupUpdate) SetTeamID(i int) *HostGroupUpdate {
	hgu.mutation.SetTeamID(i)
	return hgu
}

// SetName sets the "name" field.
func (hgu *HostGroupUpdate) SetName(s string) *HostGroupUpdate {
	hgu.mutation.SetName(s)
	return hgu
}

// SetTeam sets the "team" edge to the Team entity.
func (hgu *HostGroupUpdate) SetTeam(t *Team) *HostGroupUpdate {
	return hgu.SetTeamID(t.ID)
}

// AddHostIDs adds the "hosts" edge to the Host entity by IDs.
func (hgu *HostGroupUpdate) AddHostIDs(ids ...int) *HostGroupUpdate {
	hgu.mutation.AddHostIDs(ids...)
	return hgu
}

// AddHosts adds the "hosts" edges to the Host entity.
func (hgu *HostGroupUpdate) AddHosts(h ...*Host) *HostGroupUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hgu.AddHostIDs(ids...)
}

// Mutation returns the HostGroupMutation object of the builder.
func (hgu *HostGroupUpdate) Mutation() *HostGroupMutation {
	return hgu.mutation
}

// ClearTeam clears the "team" edge to the Team entity.
func (hgu *HostGroupUpdate) ClearTeam() *HostGroupUpdate {
	hgu.mutation.ClearTeam()
	return hgu
}

// ClearHosts clears all "hosts" edges to the Host entity.
func (hgu *HostGroupUpdate) ClearHosts() *HostGroupUpdate {
	hgu.mutation.ClearHosts()
	return hgu
}

// RemoveHostIDs removes the "hosts" edge to Host entities by IDs.
func (hgu *HostGroupUpdate) RemoveHostIDs(ids ...int) *HostGroupUpdate {
	hgu.mutation.RemoveHostIDs(ids...)
	return hgu
}

// RemoveHosts removes "hosts" edges to Host entities.
func (hgu *HostGroupUpdate) RemoveHosts(h ...*Host) *HostGroupUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hgu.RemoveHostIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hgu *HostGroupUpdate) Save(ctx context.Context) (int, error) {
	hgu.defaults()
	return withHooks[int, HostGroupMutation](ctx, hgu.sqlSave, hgu.mutation, hgu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hgu *HostGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := hgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hgu *HostGroupUpdate) Exec(ctx context.Context) error {
	_, err := hgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hgu *HostGroupUpdate) ExecX(ctx context.Context) {
	if err := hgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hgu *HostGroupUpdate) defaults() {
	if _, ok := hgu.mutation.UpdateTime(); !ok {
		v := hostgroup.UpdateDefaultUpdateTime()
		hgu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hgu *HostGroupUpdate) check() error {
	if _, ok := hgu.mutation.CompetitionID(); hgu.mutation.CompetitionCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "HostGroup.competition"`)
	}
	if _, ok := hgu.mutation.TeamID(); hgu.mutation.TeamCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "HostGroup.team"`)
	}
	return nil
}

func (hgu *HostGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hgu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(hostgroup.Table, hostgroup.Columns, sqlgraph.NewFieldSpec(hostgroup.FieldID, field.TypeInt))
	if ps := hgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hgu.mutation.UpdateTime(); ok {
		_spec.SetField(hostgroup.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := hgu.mutation.Pause(); ok {
		_spec.SetField(hostgroup.FieldPause, field.TypeBool, value)
	}
	if value, ok := hgu.mutation.Hidden(); ok {
		_spec.SetField(hostgroup.FieldHidden, field.TypeBool, value)
	}
	if value, ok := hgu.mutation.Name(); ok {
		_spec.SetField(hostgroup.FieldName, field.TypeString, value)
	}
	if hgu.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hostgroup.TeamTable,
			Columns: []string{hostgroup.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hgu.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hostgroup.TeamTable,
			Columns: []string{hostgroup.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hgu.mutation.HostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hostgroup.HostsTable,
			Columns: []string{hostgroup.HostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hgu.mutation.RemovedHostsIDs(); len(nodes) > 0 && !hgu.mutation.HostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hostgroup.HostsTable,
			Columns: []string{hostgroup.HostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hgu.mutation.HostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hostgroup.HostsTable,
			Columns: []string{hostgroup.HostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hostgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hgu.mutation.done = true
	return n, nil
}

// HostGroupUpdateOne is the builder for updating a single HostGroup entity.
type HostGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HostGroupMutation
}

// SetUpdateTime sets the "update_time" field.
func (hguo *HostGroupUpdateOne) SetUpdateTime(t time.Time) *HostGroupUpdateOne {
	hguo.mutation.SetUpdateTime(t)
	return hguo
}

// SetPause sets the "pause" field.
func (hguo *HostGroupUpdateOne) SetPause(b bool) *HostGroupUpdateOne {
	hguo.mutation.SetPause(b)
	return hguo
}

// SetHidden sets the "hidden" field.
func (hguo *HostGroupUpdateOne) SetHidden(b bool) *HostGroupUpdateOne {
	hguo.mutation.SetHidden(b)
	return hguo
}

// SetTeamID sets the "team_id" field.
func (hguo *HostGroupUpdateOne) SetTeamID(i int) *HostGroupUpdateOne {
	hguo.mutation.SetTeamID(i)
	return hguo
}

// SetName sets the "name" field.
func (hguo *HostGroupUpdateOne) SetName(s string) *HostGroupUpdateOne {
	hguo.mutation.SetName(s)
	return hguo
}

// SetTeam sets the "team" edge to the Team entity.
func (hguo *HostGroupUpdateOne) SetTeam(t *Team) *HostGroupUpdateOne {
	return hguo.SetTeamID(t.ID)
}

// AddHostIDs adds the "hosts" edge to the Host entity by IDs.
func (hguo *HostGroupUpdateOne) AddHostIDs(ids ...int) *HostGroupUpdateOne {
	hguo.mutation.AddHostIDs(ids...)
	return hguo
}

// AddHosts adds the "hosts" edges to the Host entity.
func (hguo *HostGroupUpdateOne) AddHosts(h ...*Host) *HostGroupUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hguo.AddHostIDs(ids...)
}

// Mutation returns the HostGroupMutation object of the builder.
func (hguo *HostGroupUpdateOne) Mutation() *HostGroupMutation {
	return hguo.mutation
}

// ClearTeam clears the "team" edge to the Team entity.
func (hguo *HostGroupUpdateOne) ClearTeam() *HostGroupUpdateOne {
	hguo.mutation.ClearTeam()
	return hguo
}

// ClearHosts clears all "hosts" edges to the Host entity.
func (hguo *HostGroupUpdateOne) ClearHosts() *HostGroupUpdateOne {
	hguo.mutation.ClearHosts()
	return hguo
}

// RemoveHostIDs removes the "hosts" edge to Host entities by IDs.
func (hguo *HostGroupUpdateOne) RemoveHostIDs(ids ...int) *HostGroupUpdateOne {
	hguo.mutation.RemoveHostIDs(ids...)
	return hguo
}

// RemoveHosts removes "hosts" edges to Host entities.
func (hguo *HostGroupUpdateOne) RemoveHosts(h ...*Host) *HostGroupUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hguo.RemoveHostIDs(ids...)
}

// Where appends a list predicates to the HostGroupUpdate builder.
func (hguo *HostGroupUpdateOne) Where(ps ...predicate.HostGroup) *HostGroupUpdateOne {
	hguo.mutation.Where(ps...)
	return hguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hguo *HostGroupUpdateOne) Select(field string, fields ...string) *HostGroupUpdateOne {
	hguo.fields = append([]string{field}, fields...)
	return hguo
}

// Save executes the query and returns the updated HostGroup entity.
func (hguo *HostGroupUpdateOne) Save(ctx context.Context) (*HostGroup, error) {
	hguo.defaults()
	return withHooks[*HostGroup, HostGroupMutation](ctx, hguo.sqlSave, hguo.mutation, hguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hguo *HostGroupUpdateOne) SaveX(ctx context.Context) *HostGroup {
	node, err := hguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hguo *HostGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := hguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hguo *HostGroupUpdateOne) ExecX(ctx context.Context) {
	if err := hguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hguo *HostGroupUpdateOne) defaults() {
	if _, ok := hguo.mutation.UpdateTime(); !ok {
		v := hostgroup.UpdateDefaultUpdateTime()
		hguo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hguo *HostGroupUpdateOne) check() error {
	if _, ok := hguo.mutation.CompetitionID(); hguo.mutation.CompetitionCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "HostGroup.competition"`)
	}
	if _, ok := hguo.mutation.TeamID(); hguo.mutation.TeamCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "HostGroup.team"`)
	}
	return nil
}

func (hguo *HostGroupUpdateOne) sqlSave(ctx context.Context) (_node *HostGroup, err error) {
	if err := hguo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(hostgroup.Table, hostgroup.Columns, sqlgraph.NewFieldSpec(hostgroup.FieldID, field.TypeInt))
	id, ok := hguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "HostGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hostgroup.FieldID)
		for _, f := range fields {
			if !hostgroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != hostgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hguo.mutation.UpdateTime(); ok {
		_spec.SetField(hostgroup.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := hguo.mutation.Pause(); ok {
		_spec.SetField(hostgroup.FieldPause, field.TypeBool, value)
	}
	if value, ok := hguo.mutation.Hidden(); ok {
		_spec.SetField(hostgroup.FieldHidden, field.TypeBool, value)
	}
	if value, ok := hguo.mutation.Name(); ok {
		_spec.SetField(hostgroup.FieldName, field.TypeString, value)
	}
	if hguo.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hostgroup.TeamTable,
			Columns: []string{hostgroup.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hguo.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hostgroup.TeamTable,
			Columns: []string{hostgroup.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hguo.mutation.HostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hostgroup.HostsTable,
			Columns: []string{hostgroup.HostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hguo.mutation.RemovedHostsIDs(); len(nodes) > 0 && !hguo.mutation.HostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hostgroup.HostsTable,
			Columns: []string{hostgroup.HostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hguo.mutation.HostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hostgroup.HostsTable,
			Columns: []string{hostgroup.HostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HostGroup{config: hguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hostgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	hguo.mutation.done = true
	return _node, nil
}
