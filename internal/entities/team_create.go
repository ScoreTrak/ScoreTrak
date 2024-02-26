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
	"github.com/scoretrak/scoretrak/internal/entities/host"
	"github.com/scoretrak/scoretrak/internal/entities/hostservice"
	"github.com/scoretrak/scoretrak/internal/entities/hostservicereport"
	"github.com/scoretrak/scoretrak/internal/entities/team"
	"github.com/scoretrak/scoretrak/internal/entities/teamreport"
)

// TeamCreate is the builder for creating a Team entity.
type TeamCreate struct {
	config
	mutation *TeamMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (tc *TeamCreate) SetName(s string) *TeamCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetDisplayName sets the "display_name" field.
func (tc *TeamCreate) SetDisplayName(s string) *TeamCreate {
	tc.mutation.SetDisplayName(s)
	return tc
}

// SetPause sets the "pause" field.
func (tc *TeamCreate) SetPause(b bool) *TeamCreate {
	tc.mutation.SetPause(b)
	return tc
}

// SetNillablePause sets the "pause" field if the given value is not nil.
func (tc *TeamCreate) SetNillablePause(b *bool) *TeamCreate {
	if b != nil {
		tc.SetPause(*b)
	}
	return tc
}

// SetHidden sets the "hidden" field.
func (tc *TeamCreate) SetHidden(b bool) *TeamCreate {
	tc.mutation.SetHidden(b)
	return tc
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (tc *TeamCreate) SetNillableHidden(b *bool) *TeamCreate {
	if b != nil {
		tc.SetHidden(*b)
	}
	return tc
}

// SetCreateTime sets the "create_time" field.
func (tc *TeamCreate) SetCreateTime(t time.Time) *TeamCreate {
	tc.mutation.SetCreateTime(t)
	return tc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tc *TeamCreate) SetNillableCreateTime(t *time.Time) *TeamCreate {
	if t != nil {
		tc.SetCreateTime(*t)
	}
	return tc
}

// SetUpdateTime sets the "update_time" field.
func (tc *TeamCreate) SetUpdateTime(t time.Time) *TeamCreate {
	tc.mutation.SetUpdateTime(t)
	return tc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (tc *TeamCreate) SetNillableUpdateTime(t *time.Time) *TeamCreate {
	if t != nil {
		tc.SetUpdateTime(*t)
	}
	return tc
}

// SetNumber sets the "number" field.
func (tc *TeamCreate) SetNumber(i int) *TeamCreate {
	tc.mutation.SetNumber(i)
	return tc
}

// SetID sets the "id" field.
func (tc *TeamCreate) SetID(s string) *TeamCreate {
	tc.mutation.SetID(s)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TeamCreate) SetNillableID(s *string) *TeamCreate {
	if s != nil {
		tc.SetID(*s)
	}
	return tc
}

// AddHostIDs adds the "hosts" edge to the Host entity by IDs.
func (tc *TeamCreate) AddHostIDs(ids ...string) *TeamCreate {
	tc.mutation.AddHostIDs(ids...)
	return tc
}

// AddHosts adds the "hosts" edges to the Host entity.
func (tc *TeamCreate) AddHosts(h ...*Host) *TeamCreate {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tc.AddHostIDs(ids...)
}

// AddHostserviceIDs adds the "hostservices" edge to the HostService entity by IDs.
func (tc *TeamCreate) AddHostserviceIDs(ids ...string) *TeamCreate {
	tc.mutation.AddHostserviceIDs(ids...)
	return tc
}

// AddHostservices adds the "hostservices" edges to the HostService entity.
func (tc *TeamCreate) AddHostservices(h ...*HostService) *TeamCreate {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tc.AddHostserviceIDs(ids...)
}

// SetTeamreportID sets the "teamreport" edge to the TeamReport entity by ID.
func (tc *TeamCreate) SetTeamreportID(id string) *TeamCreate {
	tc.mutation.SetTeamreportID(id)
	return tc
}

// SetNillableTeamreportID sets the "teamreport" edge to the TeamReport entity by ID if the given value is not nil.
func (tc *TeamCreate) SetNillableTeamreportID(id *string) *TeamCreate {
	if id != nil {
		tc = tc.SetTeamreportID(*id)
	}
	return tc
}

// SetTeamreport sets the "teamreport" edge to the TeamReport entity.
func (tc *TeamCreate) SetTeamreport(t *TeamReport) *TeamCreate {
	return tc.SetTeamreportID(t.ID)
}

// AddHostservicereportIDs adds the "hostservicereports" edge to the HostServiceReport entity by IDs.
func (tc *TeamCreate) AddHostservicereportIDs(ids ...string) *TeamCreate {
	tc.mutation.AddHostservicereportIDs(ids...)
	return tc
}

// AddHostservicereports adds the "hostservicereports" edges to the HostServiceReport entity.
func (tc *TeamCreate) AddHostservicereports(h ...*HostServiceReport) *TeamCreate {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tc.AddHostservicereportIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tc *TeamCreate) Mutation() *TeamMutation {
	return tc.mutation
}

// Save creates the Team in the database.
func (tc *TeamCreate) Save(ctx context.Context) (*Team, error) {
	tc.defaults()
	return withHooks[*Team, TeamMutation](ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TeamCreate) SaveX(ctx context.Context) *Team {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TeamCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TeamCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TeamCreate) defaults() {
	if _, ok := tc.mutation.Pause(); !ok {
		v := team.DefaultPause
		tc.mutation.SetPause(v)
	}
	if _, ok := tc.mutation.Hidden(); !ok {
		v := team.DefaultHidden
		tc.mutation.SetHidden(v)
	}
	if _, ok := tc.mutation.CreateTime(); !ok {
		v := team.DefaultCreateTime()
		tc.mutation.SetCreateTime(v)
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		v := team.DefaultUpdateTime()
		tc.mutation.SetUpdateTime(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := team.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TeamCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`entities: missing required field "Team.name"`)}
	}
	if _, ok := tc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`entities: missing required field "Team.display_name"`)}
	}
	if _, ok := tc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`entities: missing required field "Team.number"`)}
	}
	if v, ok := tc.mutation.Number(); ok {
		if err := team.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`entities: validator failed for field "Team.number": %w`, err)}
		}
	}
	if v, ok := tc.mutation.ID(); ok {
		if err := team.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`entities: validator failed for field "Team.id": %w`, err)}
		}
	}
	return nil
}

func (tc *TeamCreate) sqlSave(ctx context.Context) (*Team, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Team.ID type: %T", _spec.ID.Value)
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TeamCreate) createSpec() (*Team, *sqlgraph.CreateSpec) {
	var (
		_node = &Team{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(team.Table, sqlgraph.NewFieldSpec(team.FieldID, field.TypeString))
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.DisplayName(); ok {
		_spec.SetField(team.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := tc.mutation.Pause(); ok {
		_spec.SetField(team.FieldPause, field.TypeBool, value)
		_node.Pause = value
	}
	if value, ok := tc.mutation.Hidden(); ok {
		_spec.SetField(team.FieldHidden, field.TypeBool, value)
		_node.Hidden = value
	}
	if value, ok := tc.mutation.CreateTime(); ok {
		_spec.SetField(team.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := tc.mutation.UpdateTime(); ok {
		_spec.SetField(team.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := tc.mutation.Number(); ok {
		_spec.SetField(team.FieldNumber, field.TypeInt, value)
		_node.Number = value
	}
	if nodes := tc.mutation.HostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.HostsTable,
			Columns: []string{team.HostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.HostservicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.HostservicesTable,
			Columns: []string{team.HostservicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostservice.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TeamreportIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   team.TeamreportTable,
			Columns: []string{team.TeamreportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teamreport.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.HostservicereportsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.HostservicereportsTable,
			Columns: []string{team.HostservicereportsColumn},
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Team.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeamUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (tc *TeamCreate) OnConflict(opts ...sql.ConflictOption) *TeamUpsertOne {
	tc.conflict = opts
	return &TeamUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Team.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TeamCreate) OnConflictColumns(columns ...string) *TeamUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TeamUpsertOne{
		create: tc,
	}
}

type (
	// TeamUpsertOne is the builder for "upsert"-ing
	//  one Team node.
	TeamUpsertOne struct {
		create *TeamCreate
	}

	// TeamUpsert is the "OnConflict" setter.
	TeamUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *TeamUpsert) SetName(v string) *TeamUpsert {
	u.Set(team.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TeamUpsert) UpdateName() *TeamUpsert {
	u.SetExcluded(team.FieldName)
	return u
}

// SetDisplayName sets the "display_name" field.
func (u *TeamUpsert) SetDisplayName(v string) *TeamUpsert {
	u.Set(team.FieldDisplayName, v)
	return u
}

// UpdateDisplayName sets the "display_name" field to the value that was provided on create.
func (u *TeamUpsert) UpdateDisplayName() *TeamUpsert {
	u.SetExcluded(team.FieldDisplayName)
	return u
}

// SetPause sets the "pause" field.
func (u *TeamUpsert) SetPause(v bool) *TeamUpsert {
	u.Set(team.FieldPause, v)
	return u
}

// UpdatePause sets the "pause" field to the value that was provided on create.
func (u *TeamUpsert) UpdatePause() *TeamUpsert {
	u.SetExcluded(team.FieldPause)
	return u
}

// ClearPause clears the value of the "pause" field.
func (u *TeamUpsert) ClearPause() *TeamUpsert {
	u.SetNull(team.FieldPause)
	return u
}

// SetHidden sets the "hidden" field.
func (u *TeamUpsert) SetHidden(v bool) *TeamUpsert {
	u.Set(team.FieldHidden, v)
	return u
}

// UpdateHidden sets the "hidden" field to the value that was provided on create.
func (u *TeamUpsert) UpdateHidden() *TeamUpsert {
	u.SetExcluded(team.FieldHidden)
	return u
}

// ClearHidden clears the value of the "hidden" field.
func (u *TeamUpsert) ClearHidden() *TeamUpsert {
	u.SetNull(team.FieldHidden)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *TeamUpsert) SetUpdateTime(v time.Time) *TeamUpsert {
	u.Set(team.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *TeamUpsert) UpdateUpdateTime() *TeamUpsert {
	u.SetExcluded(team.FieldUpdateTime)
	return u
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *TeamUpsert) ClearUpdateTime() *TeamUpsert {
	u.SetNull(team.FieldUpdateTime)
	return u
}

// SetNumber sets the "number" field.
func (u *TeamUpsert) SetNumber(v int) *TeamUpsert {
	u.Set(team.FieldNumber, v)
	return u
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *TeamUpsert) UpdateNumber() *TeamUpsert {
	u.SetExcluded(team.FieldNumber)
	return u
}

// AddNumber adds v to the "number" field.
func (u *TeamUpsert) AddNumber(v int) *TeamUpsert {
	u.Add(team.FieldNumber, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Team.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(team.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TeamUpsertOne) UpdateNewValues() *TeamUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(team.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(team.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Team.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TeamUpsertOne) Ignore() *TeamUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeamUpsertOne) DoNothing() *TeamUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeamCreate.OnConflict
// documentation for more info.
func (u *TeamUpsertOne) Update(set func(*TeamUpsert)) *TeamUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeamUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TeamUpsertOne) SetName(v string) *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TeamUpsertOne) UpdateName() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateName()
	})
}

// SetDisplayName sets the "display_name" field.
func (u *TeamUpsertOne) SetDisplayName(v string) *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.SetDisplayName(v)
	})
}

// UpdateDisplayName sets the "display_name" field to the value that was provided on create.
func (u *TeamUpsertOne) UpdateDisplayName() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateDisplayName()
	})
}

// SetPause sets the "pause" field.
func (u *TeamUpsertOne) SetPause(v bool) *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.SetPause(v)
	})
}

// UpdatePause sets the "pause" field to the value that was provided on create.
func (u *TeamUpsertOne) UpdatePause() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.UpdatePause()
	})
}

// ClearPause clears the value of the "pause" field.
func (u *TeamUpsertOne) ClearPause() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.ClearPause()
	})
}

// SetHidden sets the "hidden" field.
func (u *TeamUpsertOne) SetHidden(v bool) *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.SetHidden(v)
	})
}

// UpdateHidden sets the "hidden" field to the value that was provided on create.
func (u *TeamUpsertOne) UpdateHidden() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateHidden()
	})
}

// ClearHidden clears the value of the "hidden" field.
func (u *TeamUpsertOne) ClearHidden() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.ClearHidden()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *TeamUpsertOne) SetUpdateTime(v time.Time) *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *TeamUpsertOne) UpdateUpdateTime() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *TeamUpsertOne) ClearUpdateTime() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.ClearUpdateTime()
	})
}

// SetNumber sets the "number" field.
func (u *TeamUpsertOne) SetNumber(v int) *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.SetNumber(v)
	})
}

// AddNumber adds v to the "number" field.
func (u *TeamUpsertOne) AddNumber(v int) *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.AddNumber(v)
	})
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *TeamUpsertOne) UpdateNumber() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateNumber()
	})
}

// Exec executes the query.
func (u *TeamUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("entities: missing options for TeamCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeamUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TeamUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("entities: TeamUpsertOne.ID is not supported by MySQL driver. Use TeamUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TeamUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TeamCreateBulk is the builder for creating many Team entities in bulk.
type TeamCreateBulk struct {
	config
	builders []*TeamCreate
	conflict []sql.ConflictOption
}

// Save creates the Team entities in the database.
func (tcb *TeamCreateBulk) Save(ctx context.Context) ([]*Team, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Team, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeamMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TeamCreateBulk) SaveX(ctx context.Context) []*Team {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TeamCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TeamCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Team.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeamUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (tcb *TeamCreateBulk) OnConflict(opts ...sql.ConflictOption) *TeamUpsertBulk {
	tcb.conflict = opts
	return &TeamUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Team.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TeamCreateBulk) OnConflictColumns(columns ...string) *TeamUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TeamUpsertBulk{
		create: tcb,
	}
}

// TeamUpsertBulk is the builder for "upsert"-ing
// a bulk of Team nodes.
type TeamUpsertBulk struct {
	create *TeamCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Team.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(team.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TeamUpsertBulk) UpdateNewValues() *TeamUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(team.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(team.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Team.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TeamUpsertBulk) Ignore() *TeamUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeamUpsertBulk) DoNothing() *TeamUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeamCreateBulk.OnConflict
// documentation for more info.
func (u *TeamUpsertBulk) Update(set func(*TeamUpsert)) *TeamUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeamUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TeamUpsertBulk) SetName(v string) *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TeamUpsertBulk) UpdateName() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateName()
	})
}

// SetDisplayName sets the "display_name" field.
func (u *TeamUpsertBulk) SetDisplayName(v string) *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.SetDisplayName(v)
	})
}

// UpdateDisplayName sets the "display_name" field to the value that was provided on create.
func (u *TeamUpsertBulk) UpdateDisplayName() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateDisplayName()
	})
}

// SetPause sets the "pause" field.
func (u *TeamUpsertBulk) SetPause(v bool) *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.SetPause(v)
	})
}

// UpdatePause sets the "pause" field to the value that was provided on create.
func (u *TeamUpsertBulk) UpdatePause() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.UpdatePause()
	})
}

// ClearPause clears the value of the "pause" field.
func (u *TeamUpsertBulk) ClearPause() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.ClearPause()
	})
}

// SetHidden sets the "hidden" field.
func (u *TeamUpsertBulk) SetHidden(v bool) *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.SetHidden(v)
	})
}

// UpdateHidden sets the "hidden" field to the value that was provided on create.
func (u *TeamUpsertBulk) UpdateHidden() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateHidden()
	})
}

// ClearHidden clears the value of the "hidden" field.
func (u *TeamUpsertBulk) ClearHidden() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.ClearHidden()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *TeamUpsertBulk) SetUpdateTime(v time.Time) *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *TeamUpsertBulk) UpdateUpdateTime() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *TeamUpsertBulk) ClearUpdateTime() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.ClearUpdateTime()
	})
}

// SetNumber sets the "number" field.
func (u *TeamUpsertBulk) SetNumber(v int) *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.SetNumber(v)
	})
}

// AddNumber adds v to the "number" field.
func (u *TeamUpsertBulk) AddNumber(v int) *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.AddNumber(v)
	})
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *TeamUpsertBulk) UpdateNumber() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateNumber()
	})
}

// Exec executes the query.
func (u *TeamUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("entities: OnConflict was set for builder %d. Set it on the TeamCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("entities: missing options for TeamCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeamUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
