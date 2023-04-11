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
	"github.com/ScoreTrak/ScoreTrak/internal/entities/service"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/team"
)

// HostUpdate is the builder for updating Host entities.
type HostUpdate struct {
	config
	hooks    []Hook
	mutation *HostMutation
}

// Where appends a list predicates to the HostUpdate builder.
func (hu *HostUpdate) Where(ps ...predicate.Host) *HostUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetUpdateTime sets the "update_time" field.
func (hu *HostUpdate) SetUpdateTime(t time.Time) *HostUpdate {
	hu.mutation.SetUpdateTime(t)
	return hu
}

// SetPause sets the "pause" field.
func (hu *HostUpdate) SetPause(b bool) *HostUpdate {
	hu.mutation.SetPause(b)
	return hu
}

// SetHidden sets the "hidden" field.
func (hu *HostUpdate) SetHidden(b bool) *HostUpdate {
	hu.mutation.SetHidden(b)
	return hu
}

// SetTeamID sets the "team_id" field.
func (hu *HostUpdate) SetTeamID(i int) *HostUpdate {
	hu.mutation.SetTeamID(i)
	return hu
}

// SetAddress sets the "address" field.
func (hu *HostUpdate) SetAddress(s string) *HostUpdate {
	hu.mutation.SetAddress(s)
	return hu
}

// SetAddressListRange sets the "address_list_range" field.
func (hu *HostUpdate) SetAddressListRange(s string) *HostUpdate {
	hu.mutation.SetAddressListRange(s)
	return hu
}

// SetEditable sets the "editable" field.
func (hu *HostUpdate) SetEditable(b bool) *HostUpdate {
	hu.mutation.SetEditable(b)
	return hu
}

// SetTeam sets the "team" edge to the Team entity.
func (hu *HostUpdate) SetTeam(t *Team) *HostUpdate {
	return hu.SetTeamID(t.ID)
}

// AddServiceIDs adds the "services" edge to the Service entity by IDs.
func (hu *HostUpdate) AddServiceIDs(ids ...int) *HostUpdate {
	hu.mutation.AddServiceIDs(ids...)
	return hu
}

// AddServices adds the "services" edges to the Service entity.
func (hu *HostUpdate) AddServices(s ...*Service) *HostUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return hu.AddServiceIDs(ids...)
}

// SetHostGroupID sets the "host_group" edge to the HostGroup entity by ID.
func (hu *HostUpdate) SetHostGroupID(id int) *HostUpdate {
	hu.mutation.SetHostGroupID(id)
	return hu
}

// SetHostGroup sets the "host_group" edge to the HostGroup entity.
func (hu *HostUpdate) SetHostGroup(h *HostGroup) *HostUpdate {
	return hu.SetHostGroupID(h.ID)
}

// Mutation returns the HostMutation object of the builder.
func (hu *HostUpdate) Mutation() *HostMutation {
	return hu.mutation
}

// ClearTeam clears the "team" edge to the Team entity.
func (hu *HostUpdate) ClearTeam() *HostUpdate {
	hu.mutation.ClearTeam()
	return hu
}

// ClearServices clears all "services" edges to the Service entity.
func (hu *HostUpdate) ClearServices() *HostUpdate {
	hu.mutation.ClearServices()
	return hu
}

// RemoveServiceIDs removes the "services" edge to Service entities by IDs.
func (hu *HostUpdate) RemoveServiceIDs(ids ...int) *HostUpdate {
	hu.mutation.RemoveServiceIDs(ids...)
	return hu
}

// RemoveServices removes "services" edges to Service entities.
func (hu *HostUpdate) RemoveServices(s ...*Service) *HostUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return hu.RemoveServiceIDs(ids...)
}

// ClearHostGroup clears the "host_group" edge to the HostGroup entity.
func (hu *HostUpdate) ClearHostGroup() *HostUpdate {
	hu.mutation.ClearHostGroup()
	return hu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HostUpdate) Save(ctx context.Context) (int, error) {
	hu.defaults()
	return withHooks[int, HostMutation](ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HostUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HostUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HostUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hu *HostUpdate) defaults() {
	if _, ok := hu.mutation.UpdateTime(); !ok {
		v := host.UpdateDefaultUpdateTime()
		hu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hu *HostUpdate) check() error {
	if v, ok := hu.mutation.Address(); ok {
		if err := host.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`entities: validator failed for field "Host.address": %w`, err)}
		}
	}
	if _, ok := hu.mutation.CompetitionID(); hu.mutation.CompetitionCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Host.competition"`)
	}
	if _, ok := hu.mutation.TeamID(); hu.mutation.TeamCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Host.team"`)
	}
	if _, ok := hu.mutation.HostGroupID(); hu.mutation.HostGroupCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Host.host_group"`)
	}
	return nil
}

func (hu *HostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(host.Table, host.Columns, sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.UpdateTime(); ok {
		_spec.SetField(host.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := hu.mutation.Pause(); ok {
		_spec.SetField(host.FieldPause, field.TypeBool, value)
	}
	if value, ok := hu.mutation.Hidden(); ok {
		_spec.SetField(host.FieldHidden, field.TypeBool, value)
	}
	if value, ok := hu.mutation.Address(); ok {
		_spec.SetField(host.FieldAddress, field.TypeString, value)
	}
	if value, ok := hu.mutation.AddressListRange(); ok {
		_spec.SetField(host.FieldAddressListRange, field.TypeString, value)
	}
	if value, ok := hu.mutation.Editable(); ok {
		_spec.SetField(host.FieldEditable, field.TypeBool, value)
	}
	if hu.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   host.TeamTable,
			Columns: []string{host.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   host.TeamTable,
			Columns: []string{host.TeamColumn},
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
	if hu.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.ServicesTable,
			Columns: []string{host.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedServicesIDs(); len(nodes) > 0 && !hu.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.ServicesTable,
			Columns: []string{host.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.ServicesTable,
			Columns: []string{host.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.HostGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   host.HostGroupTable,
			Columns: []string{host.HostGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostgroup.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.HostGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   host.HostGroupTable,
			Columns: []string{host.HostGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{host.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}

// HostUpdateOne is the builder for updating a single Host entity.
type HostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HostMutation
}

// SetUpdateTime sets the "update_time" field.
func (huo *HostUpdateOne) SetUpdateTime(t time.Time) *HostUpdateOne {
	huo.mutation.SetUpdateTime(t)
	return huo
}

// SetPause sets the "pause" field.
func (huo *HostUpdateOne) SetPause(b bool) *HostUpdateOne {
	huo.mutation.SetPause(b)
	return huo
}

// SetHidden sets the "hidden" field.
func (huo *HostUpdateOne) SetHidden(b bool) *HostUpdateOne {
	huo.mutation.SetHidden(b)
	return huo
}

// SetTeamID sets the "team_id" field.
func (huo *HostUpdateOne) SetTeamID(i int) *HostUpdateOne {
	huo.mutation.SetTeamID(i)
	return huo
}

// SetAddress sets the "address" field.
func (huo *HostUpdateOne) SetAddress(s string) *HostUpdateOne {
	huo.mutation.SetAddress(s)
	return huo
}

// SetAddressListRange sets the "address_list_range" field.
func (huo *HostUpdateOne) SetAddressListRange(s string) *HostUpdateOne {
	huo.mutation.SetAddressListRange(s)
	return huo
}

// SetEditable sets the "editable" field.
func (huo *HostUpdateOne) SetEditable(b bool) *HostUpdateOne {
	huo.mutation.SetEditable(b)
	return huo
}

// SetTeam sets the "team" edge to the Team entity.
func (huo *HostUpdateOne) SetTeam(t *Team) *HostUpdateOne {
	return huo.SetTeamID(t.ID)
}

// AddServiceIDs adds the "services" edge to the Service entity by IDs.
func (huo *HostUpdateOne) AddServiceIDs(ids ...int) *HostUpdateOne {
	huo.mutation.AddServiceIDs(ids...)
	return huo
}

// AddServices adds the "services" edges to the Service entity.
func (huo *HostUpdateOne) AddServices(s ...*Service) *HostUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return huo.AddServiceIDs(ids...)
}

// SetHostGroupID sets the "host_group" edge to the HostGroup entity by ID.
func (huo *HostUpdateOne) SetHostGroupID(id int) *HostUpdateOne {
	huo.mutation.SetHostGroupID(id)
	return huo
}

// SetHostGroup sets the "host_group" edge to the HostGroup entity.
func (huo *HostUpdateOne) SetHostGroup(h *HostGroup) *HostUpdateOne {
	return huo.SetHostGroupID(h.ID)
}

// Mutation returns the HostMutation object of the builder.
func (huo *HostUpdateOne) Mutation() *HostMutation {
	return huo.mutation
}

// ClearTeam clears the "team" edge to the Team entity.
func (huo *HostUpdateOne) ClearTeam() *HostUpdateOne {
	huo.mutation.ClearTeam()
	return huo
}

// ClearServices clears all "services" edges to the Service entity.
func (huo *HostUpdateOne) ClearServices() *HostUpdateOne {
	huo.mutation.ClearServices()
	return huo
}

// RemoveServiceIDs removes the "services" edge to Service entities by IDs.
func (huo *HostUpdateOne) RemoveServiceIDs(ids ...int) *HostUpdateOne {
	huo.mutation.RemoveServiceIDs(ids...)
	return huo
}

// RemoveServices removes "services" edges to Service entities.
func (huo *HostUpdateOne) RemoveServices(s ...*Service) *HostUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return huo.RemoveServiceIDs(ids...)
}

// ClearHostGroup clears the "host_group" edge to the HostGroup entity.
func (huo *HostUpdateOne) ClearHostGroup() *HostUpdateOne {
	huo.mutation.ClearHostGroup()
	return huo
}

// Where appends a list predicates to the HostUpdate builder.
func (huo *HostUpdateOne) Where(ps ...predicate.Host) *HostUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HostUpdateOne) Select(field string, fields ...string) *HostUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Host entity.
func (huo *HostUpdateOne) Save(ctx context.Context) (*Host, error) {
	huo.defaults()
	return withHooks[*Host, HostMutation](ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HostUpdateOne) SaveX(ctx context.Context) *Host {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HostUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HostUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (huo *HostUpdateOne) defaults() {
	if _, ok := huo.mutation.UpdateTime(); !ok {
		v := host.UpdateDefaultUpdateTime()
		huo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (huo *HostUpdateOne) check() error {
	if v, ok := huo.mutation.Address(); ok {
		if err := host.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`entities: validator failed for field "Host.address": %w`, err)}
		}
	}
	if _, ok := huo.mutation.CompetitionID(); huo.mutation.CompetitionCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Host.competition"`)
	}
	if _, ok := huo.mutation.TeamID(); huo.mutation.TeamCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Host.team"`)
	}
	if _, ok := huo.mutation.HostGroupID(); huo.mutation.HostGroupCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Host.host_group"`)
	}
	return nil
}

func (huo *HostUpdateOne) sqlSave(ctx context.Context) (_node *Host, err error) {
	if err := huo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(host.Table, host.Columns, sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt))
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "Host.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, host.FieldID)
		for _, f := range fields {
			if !host.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != host.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.UpdateTime(); ok {
		_spec.SetField(host.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := huo.mutation.Pause(); ok {
		_spec.SetField(host.FieldPause, field.TypeBool, value)
	}
	if value, ok := huo.mutation.Hidden(); ok {
		_spec.SetField(host.FieldHidden, field.TypeBool, value)
	}
	if value, ok := huo.mutation.Address(); ok {
		_spec.SetField(host.FieldAddress, field.TypeString, value)
	}
	if value, ok := huo.mutation.AddressListRange(); ok {
		_spec.SetField(host.FieldAddressListRange, field.TypeString, value)
	}
	if value, ok := huo.mutation.Editable(); ok {
		_spec.SetField(host.FieldEditable, field.TypeBool, value)
	}
	if huo.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   host.TeamTable,
			Columns: []string{host.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   host.TeamTable,
			Columns: []string{host.TeamColumn},
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
	if huo.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.ServicesTable,
			Columns: []string{host.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedServicesIDs(); len(nodes) > 0 && !huo.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.ServicesTable,
			Columns: []string{host.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.ServicesTable,
			Columns: []string{host.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.HostGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   host.HostGroupTable,
			Columns: []string{host.HostGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostgroup.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.HostGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   host.HostGroupTable,
			Columns: []string{host.HostGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Host{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{host.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}
