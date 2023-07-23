// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/check"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/host"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/hostservice"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/property"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/service"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/team"
)

// HostServiceCreate is the builder for creating a HostService entity.
type HostServiceCreate struct {
	config
	mutation *HostServiceMutation
	hooks    []Hook
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

// SetWeight sets the "weight" field.
func (hsc *HostServiceCreate) SetWeight(i int) *HostServiceCreate {
	hsc.mutation.SetWeight(i)
	return hsc
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (hsc *HostServiceCreate) SetNillableWeight(i *int) *HostServiceCreate {
	if i != nil {
		hsc.SetWeight(*i)
	}
	return hsc
}

// SetPointBoost sets the "point_boost" field.
func (hsc *HostServiceCreate) SetPointBoost(i int) *HostServiceCreate {
	hsc.mutation.SetPointBoost(i)
	return hsc
}

// SetNillablePointBoost sets the "point_boost" field if the given value is not nil.
func (hsc *HostServiceCreate) SetNillablePointBoost(i *int) *HostServiceCreate {
	if i != nil {
		hsc.SetPointBoost(*i)
	}
	return hsc
}

// SetRoundUnits sets the "round_units" field.
func (hsc *HostServiceCreate) SetRoundUnits(i int) *HostServiceCreate {
	hsc.mutation.SetRoundUnits(i)
	return hsc
}

// SetNillableRoundUnits sets the "round_units" field if the given value is not nil.
func (hsc *HostServiceCreate) SetNillableRoundUnits(i *int) *HostServiceCreate {
	if i != nil {
		hsc.SetRoundUnits(*i)
	}
	return hsc
}

// SetRoundDelay sets the "round_delay" field.
func (hsc *HostServiceCreate) SetRoundDelay(i int) *HostServiceCreate {
	hsc.mutation.SetRoundDelay(i)
	return hsc
}

// SetNillableRoundDelay sets the "round_delay" field if the given value is not nil.
func (hsc *HostServiceCreate) SetNillableRoundDelay(i *int) *HostServiceCreate {
	if i != nil {
		hsc.SetRoundDelay(*i)
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
	if _, ok := hsc.mutation.Hidden(); !ok {
		v := hostservice.DefaultHidden
		hsc.mutation.SetHidden(v)
	}
	if _, ok := hsc.mutation.Weight(); !ok {
		v := hostservice.DefaultWeight
		hsc.mutation.SetWeight(v)
	}
	if _, ok := hsc.mutation.PointBoost(); !ok {
		v := hostservice.DefaultPointBoost
		hsc.mutation.SetPointBoost(v)
	}
	if _, ok := hsc.mutation.RoundUnits(); !ok {
		v := hostservice.DefaultRoundUnits
		hsc.mutation.SetRoundUnits(v)
	}
	if _, ok := hsc.mutation.RoundDelay(); !ok {
		v := hostservice.DefaultRoundDelay
		hsc.mutation.SetRoundDelay(v)
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
	if v, ok := hsc.mutation.Name(); ok {
		if err := hostservice.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entities: validator failed for field "HostService.name": %w`, err)}
		}
	}
	if _, ok := hsc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`entities: missing required field "HostService.display_name"`)}
	}
	if v, ok := hsc.mutation.DisplayName(); ok {
		if err := hostservice.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`entities: validator failed for field "HostService.display_name": %w`, err)}
		}
	}
	if _, ok := hsc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`entities: missing required field "HostService.weight"`)}
	}
	if _, ok := hsc.mutation.PointBoost(); !ok {
		return &ValidationError{Name: "point_boost", err: errors.New(`entities: missing required field "HostService.point_boost"`)}
	}
	if _, ok := hsc.mutation.RoundUnits(); !ok {
		return &ValidationError{Name: "round_units", err: errors.New(`entities: missing required field "HostService.round_units"`)}
	}
	if _, ok := hsc.mutation.RoundDelay(); !ok {
		return &ValidationError{Name: "round_delay", err: errors.New(`entities: missing required field "HostService.round_delay"`)}
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
	if value, ok := hsc.mutation.Weight(); ok {
		_spec.SetField(hostservice.FieldWeight, field.TypeInt, value)
		_node.Weight = value
	}
	if value, ok := hsc.mutation.PointBoost(); ok {
		_spec.SetField(hostservice.FieldPointBoost, field.TypeInt, value)
		_node.PointBoost = value
	}
	if value, ok := hsc.mutation.RoundUnits(); ok {
		_spec.SetField(hostservice.FieldRoundUnits, field.TypeInt, value)
		_node.RoundUnits = value
	}
	if value, ok := hsc.mutation.RoundDelay(); ok {
		_spec.SetField(hostservice.FieldRoundDelay, field.TypeInt, value)
		_node.RoundDelay = value
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

// HostServiceCreateBulk is the builder for creating many HostService entities in bulk.
type HostServiceCreateBulk struct {
	config
	builders []*HostServiceCreate
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