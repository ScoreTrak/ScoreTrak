// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/hostservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/hostservicereport"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/team"
)

// HostServiceCreate is the builder for creating a HostService entity.
type HostServiceCreate struct {
	config
	mutation *HostServiceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (hsc *HostServiceCreate) SetName(s string) *HostServiceCreate {
	hsc.mutation.SetName(s)
	return hsc
}

// SetDisplayName sets the "display_name" field.
func (hsc *HostServiceCreate) SetDisplayName(s string) *HostServiceCreate {
	hsc.mutation.SetDisplayName(s)
	return hsc
}

// SetPause sets the "pause" field.
func (hsc *HostServiceCreate) SetPause(b bool) *HostServiceCreate {
	hsc.mutation.SetPause(b)
	return hsc
}

// SetNillablePause sets the "pause" field if the given value is not nil.
func (hsc *HostServiceCreate) SetNillablePause(b *bool) *HostServiceCreate {
	if b != nil {
		hsc.SetPause(*b)
	}
	return hsc
}

// SetHidden sets the "hidden" field.
func (hsc *HostServiceCreate) SetHidden(b bool) *HostServiceCreate {
	hsc.mutation.SetHidden(b)
	return hsc
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (hsc *HostServiceCreate) SetNillableHidden(b *bool) *HostServiceCreate {
	if b != nil {
		hsc.SetHidden(*b)
	}
	return hsc
}

// SetCreateTime sets the "create_time" field.
func (hsc *HostServiceCreate) SetCreateTime(t time.Time) *HostServiceCreate {
	hsc.mutation.SetCreateTime(t)
	return hsc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (hsc *HostServiceCreate) SetNillableCreateTime(t *time.Time) *HostServiceCreate {
	if t != nil {
		hsc.SetCreateTime(*t)
	}
	return hsc
}

// SetUpdateTime sets the "update_time" field.
func (hsc *HostServiceCreate) SetUpdateTime(t time.Time) *HostServiceCreate {
	hsc.mutation.SetUpdateTime(t)
	return hsc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (hsc *HostServiceCreate) SetNillableUpdateTime(t *time.Time) *HostServiceCreate {
	if t != nil {
		hsc.SetUpdateTime(*t)
	}
	return hsc
}

// SetServiceID sets the "service_id" field.
func (hsc *HostServiceCreate) SetServiceID(s string) *HostServiceCreate {
	hsc.mutation.SetServiceID(s)
	return hsc
}

// SetHostID sets the "host_id" field.
func (hsc *HostServiceCreate) SetHostID(s string) *HostServiceCreate {
	hsc.mutation.SetHostID(s)
	return hsc
}

// SetTeamID sets the "team_id" field.
func (hsc *HostServiceCreate) SetTeamID(s string) *HostServiceCreate {
	hsc.mutation.SetTeamID(s)
	return hsc
}

// SetID sets the "id" field.
func (hsc *HostServiceCreate) SetID(s string) *HostServiceCreate {
	hsc.mutation.SetID(s)
	return hsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (hsc *HostServiceCreate) SetNillableID(s *string) *HostServiceCreate {
	if s != nil {
		hsc.SetID(*s)
	}
	return hsc
}

// AddCheckIDs adds the "checks" edge to the Check entity by IDs.
func (hsc *HostServiceCreate) AddCheckIDs(ids ...string) *HostServiceCreate {
	hsc.mutation.AddCheckIDs(ids...)
	return hsc
}

// AddChecks adds the "checks" edges to the Check entity.
func (hsc *HostServiceCreate) AddChecks(c ...*Check) *HostServiceCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return hsc.AddCheckIDs(ids...)
}

// AddPropertyIDs adds the "properties" edge to the Property entity by IDs.
func (hsc *HostServiceCreate) AddPropertyIDs(ids ...string) *HostServiceCreate {
	hsc.mutation.AddPropertyIDs(ids...)
	return hsc
}

// AddProperties adds the "properties" edges to the Property entity.
func (hsc *HostServiceCreate) AddProperties(p ...*Property) *HostServiceCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return hsc.AddPropertyIDs(ids...)
}

// SetHostservicereportID sets the "hostservicereport" edge to the HostServiceReport entity by ID.
func (hsc *HostServiceCreate) SetHostservicereportID(id string) *HostServiceCreate {
	hsc.mutation.SetHostservicereportID(id)
	return hsc
}

// SetNillableHostservicereportID sets the "hostservicereport" edge to the HostServiceReport entity by ID if the given value is not nil.
func (hsc *HostServiceCreate) SetNillableHostservicereportID(id *string) *HostServiceCreate {
	if id != nil {
		hsc = hsc.SetHostservicereportID(*id)
	}
	return hsc
}

// SetHostservicereport sets the "hostservicereport" edge to the HostServiceReport entity.
func (hsc *HostServiceCreate) SetHostservicereport(h *HostServiceReport) *HostServiceCreate {
	return hsc.SetHostservicereportID(h.ID)
}

// SetService sets the "service" edge to the Service entity.
func (hsc *HostServiceCreate) SetService(s *Service) *HostServiceCreate {
	return hsc.SetServiceID(s.ID)
}

// SetHost sets the "host" edge to the Host entity.
func (hsc *HostServiceCreate) SetHost(h *Host) *HostServiceCreate {
	return hsc.SetHostID(h.ID)
}

// SetTeam sets the "team" edge to the Team entity.
func (hsc *HostServiceCreate) SetTeam(t *Team) *HostServiceCreate {
	return hsc.SetTeamID(t.ID)
}

// Mutation returns the HostServiceMutation object of the builder.
func (hsc *HostServiceCreate) Mutation() *HostServiceMutation {
	return hsc.mutation
}

// Save creates the HostService in the database.
func (hsc *HostServiceCreate) Save(ctx context.Context) (*HostService, error) {
	hsc.defaults()
	return withHooks[*HostService, HostServiceMutation](ctx, hsc.sqlSave, hsc.mutation, hsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hsc *HostServiceCreate) SaveX(ctx context.Context) *HostService {
	v, err := hsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hsc *HostServiceCreate) Exec(ctx context.Context) error {
	_, err := hsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hsc *HostServiceCreate) ExecX(ctx context.Context) {
	if err := hsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hsc *HostServiceCreate) defaults() {
	if _, ok := hsc.mutation.Pause(); !ok {
		v := hostservice.DefaultPause
		hsc.mutation.SetPause(v)
	}
	if _, ok := hsc.mutation.Hidden(); !ok {
		v := hostservice.DefaultHidden
		hsc.mutation.SetHidden(v)
	}
	if _, ok := hsc.mutation.CreateTime(); !ok {
		v := hostservice.DefaultCreateTime()
		hsc.mutation.SetCreateTime(v)
	}
	if _, ok := hsc.mutation.UpdateTime(); !ok {
		v := hostservice.DefaultUpdateTime()
		hsc.mutation.SetUpdateTime(v)
	}
	if _, ok := hsc.mutation.ID(); !ok {
		v := hostservice.DefaultID()
		hsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hsc *HostServiceCreate) check() error {
	if _, ok := hsc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`entities: missing required field "HostService.name"`)}
	}
	if _, ok := hsc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`entities: missing required field "HostService.display_name"`)}
	}
	if _, ok := hsc.mutation.ServiceID(); !ok {
		return &ValidationError{Name: "service_id", err: errors.New(`entities: missing required field "HostService.service_id"`)}
	}
	if _, ok := hsc.mutation.HostID(); !ok {
		return &ValidationError{Name: "host_id", err: errors.New(`entities: missing required field "HostService.host_id"`)}
	}
	if _, ok := hsc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team_id", err: errors.New(`entities: missing required field "HostService.team_id"`)}
	}
	if v, ok := hsc.mutation.ID(); ok {
		if err := hostservice.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`entities: validator failed for field "HostService.id": %w`, err)}
		}
	}
	if _, ok := hsc.mutation.ServiceID(); !ok {
		return &ValidationError{Name: "service", err: errors.New(`entities: missing required edge "HostService.service"`)}
	}
	if _, ok := hsc.mutation.HostID(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`entities: missing required edge "HostService.host"`)}
	}
	if _, ok := hsc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team", err: errors.New(`entities: missing required edge "HostService.team"`)}
	}
	return nil
}

func (hsc *HostServiceCreate) sqlSave(ctx context.Context) (*HostService, error) {
	if err := hsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected HostService.ID type: %T", _spec.ID.Value)
		}
	}
	hsc.mutation.id = &_node.ID
	hsc.mutation.done = true
	return _node, nil
}

func (hsc *HostServiceCreate) createSpec() (*HostService, *sqlgraph.CreateSpec) {
	var (
		_node = &HostService{config: hsc.config}
		_spec = sqlgraph.NewCreateSpec(hostservice.Table, sqlgraph.NewFieldSpec(hostservice.FieldID, field.TypeString))
	)
	_spec.OnConflict = hsc.conflict
	if id, ok := hsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := hsc.mutation.Name(); ok {
		_spec.SetField(hostservice.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := hsc.mutation.DisplayName(); ok {
		_spec.SetField(hostservice.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := hsc.mutation.Pause(); ok {
		_spec.SetField(hostservice.FieldPause, field.TypeBool, value)
		_node.Pause = value
	}
	if value, ok := hsc.mutation.Hidden(); ok {
		_spec.SetField(hostservice.FieldHidden, field.TypeBool, value)
		_node.Hidden = value
	}
	if value, ok := hsc.mutation.CreateTime(); ok {
		_spec.SetField(hostservice.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := hsc.mutation.UpdateTime(); ok {
		_spec.SetField(hostservice.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if nodes := hsc.mutation.ChecksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hostservice.ChecksTable,
			Columns: []string{hostservice.ChecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(check.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hsc.mutation.PropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hostservice.PropertiesTable,
			Columns: []string{hostservice.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(property.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hsc.mutation.HostservicereportIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   hostservice.HostservicereportTable,
			Columns: []string{hostservice.HostservicereportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostservicereport.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hsc.mutation.ServiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hostservice.ServiceTable,
			Columns: []string{hostservice.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ServiceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hsc.mutation.HostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hostservice.HostTable,
			Columns: []string{hostservice.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.HostID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hsc.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hostservice.TeamTable,
			Columns: []string{hostservice.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TeamID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.HostService.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HostServiceUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (hsc *HostServiceCreate) OnConflict(opts ...sql.ConflictOption) *HostServiceUpsertOne {
	hsc.conflict = opts
	return &HostServiceUpsertOne{
		create: hsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.HostService.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hsc *HostServiceCreate) OnConflictColumns(columns ...string) *HostServiceUpsertOne {
	hsc.conflict = append(hsc.conflict, sql.ConflictColumns(columns...))
	return &HostServiceUpsertOne{
		create: hsc,
	}
}

type (
	// HostServiceUpsertOne is the builder for "upsert"-ing
	//  one HostService node.
	HostServiceUpsertOne struct {
		create *HostServiceCreate
	}

	// HostServiceUpsert is the "OnConflict" setter.
	HostServiceUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *HostServiceUpsert) SetName(v string) *HostServiceUpsert {
	u.Set(hostservice.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *HostServiceUpsert) UpdateName() *HostServiceUpsert {
	u.SetExcluded(hostservice.FieldName)
	return u
}

// SetDisplayName sets the "display_name" field.
func (u *HostServiceUpsert) SetDisplayName(v string) *HostServiceUpsert {
	u.Set(hostservice.FieldDisplayName, v)
	return u
}

// UpdateDisplayName sets the "display_name" field to the value that was provided on create.
func (u *HostServiceUpsert) UpdateDisplayName() *HostServiceUpsert {
	u.SetExcluded(hostservice.FieldDisplayName)
	return u
}

// SetPause sets the "pause" field.
func (u *HostServiceUpsert) SetPause(v bool) *HostServiceUpsert {
	u.Set(hostservice.FieldPause, v)
	return u
}

// UpdatePause sets the "pause" field to the value that was provided on create.
func (u *HostServiceUpsert) UpdatePause() *HostServiceUpsert {
	u.SetExcluded(hostservice.FieldPause)
	return u
}

// ClearPause clears the value of the "pause" field.
func (u *HostServiceUpsert) ClearPause() *HostServiceUpsert {
	u.SetNull(hostservice.FieldPause)
	return u
}

// SetHidden sets the "hidden" field.
func (u *HostServiceUpsert) SetHidden(v bool) *HostServiceUpsert {
	u.Set(hostservice.FieldHidden, v)
	return u
}

// UpdateHidden sets the "hidden" field to the value that was provided on create.
func (u *HostServiceUpsert) UpdateHidden() *HostServiceUpsert {
	u.SetExcluded(hostservice.FieldHidden)
	return u
}

// ClearHidden clears the value of the "hidden" field.
func (u *HostServiceUpsert) ClearHidden() *HostServiceUpsert {
	u.SetNull(hostservice.FieldHidden)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *HostServiceUpsert) SetUpdateTime(v time.Time) *HostServiceUpsert {
	u.Set(hostservice.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *HostServiceUpsert) UpdateUpdateTime() *HostServiceUpsert {
	u.SetExcluded(hostservice.FieldUpdateTime)
	return u
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *HostServiceUpsert) ClearUpdateTime() *HostServiceUpsert {
	u.SetNull(hostservice.FieldUpdateTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.HostService.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(hostservice.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *HostServiceUpsertOne) UpdateNewValues() *HostServiceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(hostservice.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(hostservice.FieldCreateTime)
		}
		if _, exists := u.create.mutation.ServiceID(); exists {
			s.SetIgnore(hostservice.FieldServiceID)
		}
		if _, exists := u.create.mutation.HostID(); exists {
			s.SetIgnore(hostservice.FieldHostID)
		}
		if _, exists := u.create.mutation.TeamID(); exists {
			s.SetIgnore(hostservice.FieldTeamID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.HostService.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *HostServiceUpsertOne) Ignore() *HostServiceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HostServiceUpsertOne) DoNothing() *HostServiceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HostServiceCreate.OnConflict
// documentation for more info.
func (u *HostServiceUpsertOne) Update(set func(*HostServiceUpsert)) *HostServiceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HostServiceUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *HostServiceUpsertOne) SetName(v string) *HostServiceUpsertOne {
	return u.Update(func(s *HostServiceUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *HostServiceUpsertOne) UpdateName() *HostServiceUpsertOne {
	return u.Update(func(s *HostServiceUpsert) {
		s.UpdateName()
	})
}

// SetDisplayName sets the "display_name" field.
func (u *HostServiceUpsertOne) SetDisplayName(v string) *HostServiceUpsertOne {
	return u.Update(func(s *HostServiceUpsert) {
		s.SetDisplayName(v)
	})
}

// UpdateDisplayName sets the "display_name" field to the value that was provided on create.
func (u *HostServiceUpsertOne) UpdateDisplayName() *HostServiceUpsertOne {
	return u.Update(func(s *HostServiceUpsert) {
		s.UpdateDisplayName()
	})
}

// SetPause sets the "pause" field.
func (u *HostServiceUpsertOne) SetPause(v bool) *HostServiceUpsertOne {
	return u.Update(func(s *HostServiceUpsert) {
		s.SetPause(v)
	})
}

// UpdatePause sets the "pause" field to the value that was provided on create.
func (u *HostServiceUpsertOne) UpdatePause() *HostServiceUpsertOne {
	return u.Update(func(s *HostServiceUpsert) {
		s.UpdatePause()
	})
}

// ClearPause clears the value of the "pause" field.
func (u *HostServiceUpsertOne) ClearPause() *HostServiceUpsertOne {
	return u.Update(func(s *HostServiceUpsert) {
		s.ClearPause()
	})
}

// SetHidden sets the "hidden" field.
func (u *HostServiceUpsertOne) SetHidden(v bool) *HostServiceUpsertOne {
	return u.Update(func(s *HostServiceUpsert) {
		s.SetHidden(v)
	})
}

// UpdateHidden sets the "hidden" field to the value that was provided on create.
func (u *HostServiceUpsertOne) UpdateHidden() *HostServiceUpsertOne {
	return u.Update(func(s *HostServiceUpsert) {
		s.UpdateHidden()
	})
}

// ClearHidden clears the value of the "hidden" field.
func (u *HostServiceUpsertOne) ClearHidden() *HostServiceUpsertOne {
	return u.Update(func(s *HostServiceUpsert) {
		s.ClearHidden()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *HostServiceUpsertOne) SetUpdateTime(v time.Time) *HostServiceUpsertOne {
	return u.Update(func(s *HostServiceUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *HostServiceUpsertOne) UpdateUpdateTime() *HostServiceUpsertOne {
	return u.Update(func(s *HostServiceUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *HostServiceUpsertOne) ClearUpdateTime() *HostServiceUpsertOne {
	return u.Update(func(s *HostServiceUpsert) {
		s.ClearUpdateTime()
	})
}

// Exec executes the query.
func (u *HostServiceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("entities: missing options for HostServiceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HostServiceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *HostServiceUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("entities: HostServiceUpsertOne.ID is not supported by MySQL driver. Use HostServiceUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *HostServiceUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// HostServiceCreateBulk is the builder for creating many HostService entities in bulk.
type HostServiceCreateBulk struct {
	config
	builders []*HostServiceCreate
	conflict []sql.ConflictOption
}

// Save creates the HostService entities in the database.
func (hscb *HostServiceCreateBulk) Save(ctx context.Context) ([]*HostService, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hscb.builders))
	nodes := make([]*HostService, len(hscb.builders))
	mutators := make([]Mutator, len(hscb.builders))
	for i := range hscb.builders {
		func(i int, root context.Context) {
			builder := hscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HostServiceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = hscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hscb *HostServiceCreateBulk) SaveX(ctx context.Context) []*HostService {
	v, err := hscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hscb *HostServiceCreateBulk) Exec(ctx context.Context) error {
	_, err := hscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hscb *HostServiceCreateBulk) ExecX(ctx context.Context) {
	if err := hscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.HostService.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HostServiceUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (hscb *HostServiceCreateBulk) OnConflict(opts ...sql.ConflictOption) *HostServiceUpsertBulk {
	hscb.conflict = opts
	return &HostServiceUpsertBulk{
		create: hscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.HostService.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hscb *HostServiceCreateBulk) OnConflictColumns(columns ...string) *HostServiceUpsertBulk {
	hscb.conflict = append(hscb.conflict, sql.ConflictColumns(columns...))
	return &HostServiceUpsertBulk{
		create: hscb,
	}
}

// HostServiceUpsertBulk is the builder for "upsert"-ing
// a bulk of HostService nodes.
type HostServiceUpsertBulk struct {
	create *HostServiceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.HostService.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(hostservice.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *HostServiceUpsertBulk) UpdateNewValues() *HostServiceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(hostservice.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(hostservice.FieldCreateTime)
			}
			if _, exists := b.mutation.ServiceID(); exists {
				s.SetIgnore(hostservice.FieldServiceID)
			}
			if _, exists := b.mutation.HostID(); exists {
				s.SetIgnore(hostservice.FieldHostID)
			}
			if _, exists := b.mutation.TeamID(); exists {
				s.SetIgnore(hostservice.FieldTeamID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.HostService.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *HostServiceUpsertBulk) Ignore() *HostServiceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HostServiceUpsertBulk) DoNothing() *HostServiceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HostServiceCreateBulk.OnConflict
// documentation for more info.
func (u *HostServiceUpsertBulk) Update(set func(*HostServiceUpsert)) *HostServiceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HostServiceUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *HostServiceUpsertBulk) SetName(v string) *HostServiceUpsertBulk {
	return u.Update(func(s *HostServiceUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *HostServiceUpsertBulk) UpdateName() *HostServiceUpsertBulk {
	return u.Update(func(s *HostServiceUpsert) {
		s.UpdateName()
	})
}

// SetDisplayName sets the "display_name" field.
func (u *HostServiceUpsertBulk) SetDisplayName(v string) *HostServiceUpsertBulk {
	return u.Update(func(s *HostServiceUpsert) {
		s.SetDisplayName(v)
	})
}

// UpdateDisplayName sets the "display_name" field to the value that was provided on create.
func (u *HostServiceUpsertBulk) UpdateDisplayName() *HostServiceUpsertBulk {
	return u.Update(func(s *HostServiceUpsert) {
		s.UpdateDisplayName()
	})
}

// SetPause sets the "pause" field.
func (u *HostServiceUpsertBulk) SetPause(v bool) *HostServiceUpsertBulk {
	return u.Update(func(s *HostServiceUpsert) {
		s.SetPause(v)
	})
}

// UpdatePause sets the "pause" field to the value that was provided on create.
func (u *HostServiceUpsertBulk) UpdatePause() *HostServiceUpsertBulk {
	return u.Update(func(s *HostServiceUpsert) {
		s.UpdatePause()
	})
}

// ClearPause clears the value of the "pause" field.
func (u *HostServiceUpsertBulk) ClearPause() *HostServiceUpsertBulk {
	return u.Update(func(s *HostServiceUpsert) {
		s.ClearPause()
	})
}

// SetHidden sets the "hidden" field.
func (u *HostServiceUpsertBulk) SetHidden(v bool) *HostServiceUpsertBulk {
	return u.Update(func(s *HostServiceUpsert) {
		s.SetHidden(v)
	})
}

// UpdateHidden sets the "hidden" field to the value that was provided on create.
func (u *HostServiceUpsertBulk) UpdateHidden() *HostServiceUpsertBulk {
	return u.Update(func(s *HostServiceUpsert) {
		s.UpdateHidden()
	})
}

// ClearHidden clears the value of the "hidden" field.
func (u *HostServiceUpsertBulk) ClearHidden() *HostServiceUpsertBulk {
	return u.Update(func(s *HostServiceUpsert) {
		s.ClearHidden()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *HostServiceUpsertBulk) SetUpdateTime(v time.Time) *HostServiceUpsertBulk {
	return u.Update(func(s *HostServiceUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *HostServiceUpsertBulk) UpdateUpdateTime() *HostServiceUpsertBulk {
	return u.Update(func(s *HostServiceUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *HostServiceUpsertBulk) ClearUpdateTime() *HostServiceUpsertBulk {
	return u.Update(func(s *HostServiceUpsert) {
		s.ClearUpdateTime()
	})
}

// Exec executes the query.
func (u *HostServiceUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("entities: OnConflict was set for builder %d. Set it on the HostServiceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("entities: missing options for HostServiceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HostServiceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
