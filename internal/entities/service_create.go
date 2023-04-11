// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/check"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/competition"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/host"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/property"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/service"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/team"
)

// ServiceCreate is the builder for creating a Service entity.
type ServiceCreate struct {
	config
	mutation *ServiceMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (sc *ServiceCreate) SetCreateTime(t time.Time) *ServiceCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableCreateTime(t *time.Time) *ServiceCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *ServiceCreate) SetUpdateTime(t time.Time) *ServiceCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableUpdateTime(t *time.Time) *ServiceCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetPause sets the "pause" field.
func (sc *ServiceCreate) SetPause(b bool) *ServiceCreate {
	sc.mutation.SetPause(b)
	return sc
}

// SetHidden sets the "hidden" field.
func (sc *ServiceCreate) SetHidden(b bool) *ServiceCreate {
	sc.mutation.SetHidden(b)
	return sc
}

// SetCompetitionID sets the "competition_id" field.
func (sc *ServiceCreate) SetCompetitionID(i int) *ServiceCreate {
	sc.mutation.SetCompetitionID(i)
	return sc
}

// SetTeamID sets the "team_id" field.
func (sc *ServiceCreate) SetTeamID(i int) *ServiceCreate {
	sc.mutation.SetTeamID(i)
	return sc
}

// SetName sets the "name" field.
func (sc *ServiceCreate) SetName(s string) *ServiceCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDisplayName sets the "display_name" field.
func (sc *ServiceCreate) SetDisplayName(s string) *ServiceCreate {
	sc.mutation.SetDisplayName(s)
	return sc
}

// SetWeight sets the "weight" field.
func (sc *ServiceCreate) SetWeight(i int) *ServiceCreate {
	sc.mutation.SetWeight(i)
	return sc
}

// SetPointBoost sets the "point_boost" field.
func (sc *ServiceCreate) SetPointBoost(i int) *ServiceCreate {
	sc.mutation.SetPointBoost(i)
	return sc
}

// SetRoundUnits sets the "round_units" field.
func (sc *ServiceCreate) SetRoundUnits(i int) *ServiceCreate {
	sc.mutation.SetRoundUnits(i)
	return sc
}

// SetRoundDelay sets the "round_delay" field.
func (sc *ServiceCreate) SetRoundDelay(i int) *ServiceCreate {
	sc.mutation.SetRoundDelay(i)
	return sc
}

// SetCompetition sets the "competition" edge to the Competition entity.
func (sc *ServiceCreate) SetCompetition(c *Competition) *ServiceCreate {
	return sc.SetCompetitionID(c.ID)
}

// SetTeam sets the "team" edge to the Team entity.
func (sc *ServiceCreate) SetTeam(t *Team) *ServiceCreate {
	return sc.SetTeamID(t.ID)
}

// SetHostsID sets the "hosts" edge to the Host entity by ID.
func (sc *ServiceCreate) SetHostsID(id int) *ServiceCreate {
	sc.mutation.SetHostsID(id)
	return sc
}

// SetNillableHostsID sets the "hosts" edge to the Host entity by ID if the given value is not nil.
func (sc *ServiceCreate) SetNillableHostsID(id *int) *ServiceCreate {
	if id != nil {
		sc = sc.SetHostsID(*id)
	}
	return sc
}

// SetHosts sets the "hosts" edge to the Host entity.
func (sc *ServiceCreate) SetHosts(h *Host) *ServiceCreate {
	return sc.SetHostsID(h.ID)
}

// AddCheckIDs adds the "checks" edge to the Check entity by IDs.
func (sc *ServiceCreate) AddCheckIDs(ids ...int) *ServiceCreate {
	sc.mutation.AddCheckIDs(ids...)
	return sc
}

// AddChecks adds the "checks" edges to the Check entity.
func (sc *ServiceCreate) AddChecks(c ...*Check) *ServiceCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddCheckIDs(ids...)
}

// AddPropertyIDs adds the "properties" edge to the Property entity by IDs.
func (sc *ServiceCreate) AddPropertyIDs(ids ...int) *ServiceCreate {
	sc.mutation.AddPropertyIDs(ids...)
	return sc
}

// AddProperties adds the "properties" edges to the Property entity.
func (sc *ServiceCreate) AddProperties(p ...*Property) *ServiceCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddPropertyIDs(ids...)
}

// Mutation returns the ServiceMutation object of the builder.
func (sc *ServiceCreate) Mutation() *ServiceMutation {
	return sc.mutation
}

// Save creates the Service in the database.
func (sc *ServiceCreate) Save(ctx context.Context) (*Service, error) {
	sc.defaults()
	return withHooks[*Service, ServiceMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ServiceCreate) SaveX(ctx context.Context) *Service {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ServiceCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ServiceCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ServiceCreate) defaults() {
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := service.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := service.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ServiceCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`entities: missing required field "Service.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`entities: missing required field "Service.update_time"`)}
	}
	if _, ok := sc.mutation.Pause(); !ok {
		return &ValidationError{Name: "pause", err: errors.New(`entities: missing required field "Service.pause"`)}
	}
	if _, ok := sc.mutation.Hidden(); !ok {
		return &ValidationError{Name: "hidden", err: errors.New(`entities: missing required field "Service.hidden"`)}
	}
	if _, ok := sc.mutation.CompetitionID(); !ok {
		return &ValidationError{Name: "competition_id", err: errors.New(`entities: missing required field "Service.competition_id"`)}
	}
	if _, ok := sc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team_id", err: errors.New(`entities: missing required field "Service.team_id"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`entities: missing required field "Service.name"`)}
	}
	if _, ok := sc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`entities: missing required field "Service.display_name"`)}
	}
	if _, ok := sc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`entities: missing required field "Service.weight"`)}
	}
	if _, ok := sc.mutation.PointBoost(); !ok {
		return &ValidationError{Name: "point_boost", err: errors.New(`entities: missing required field "Service.point_boost"`)}
	}
	if _, ok := sc.mutation.RoundUnits(); !ok {
		return &ValidationError{Name: "round_units", err: errors.New(`entities: missing required field "Service.round_units"`)}
	}
	if _, ok := sc.mutation.RoundDelay(); !ok {
		return &ValidationError{Name: "round_delay", err: errors.New(`entities: missing required field "Service.round_delay"`)}
	}
	if _, ok := sc.mutation.CompetitionID(); !ok {
		return &ValidationError{Name: "competition", err: errors.New(`entities: missing required edge "Service.competition"`)}
	}
	if _, ok := sc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team", err: errors.New(`entities: missing required edge "Service.team"`)}
	}
	return nil
}

func (sc *ServiceCreate) sqlSave(ctx context.Context) (*Service, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ServiceCreate) createSpec() (*Service, *sqlgraph.CreateSpec) {
	var (
		_node = &Service{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(service.Table, sqlgraph.NewFieldSpec(service.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.SetField(service.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.SetField(service.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.Pause(); ok {
		_spec.SetField(service.FieldPause, field.TypeBool, value)
		_node.Pause = value
	}
	if value, ok := sc.mutation.Hidden(); ok {
		_spec.SetField(service.FieldHidden, field.TypeBool, value)
		_node.Hidden = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(service.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.DisplayName(); ok {
		_spec.SetField(service.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := sc.mutation.Weight(); ok {
		_spec.SetField(service.FieldWeight, field.TypeInt, value)
		_node.Weight = value
	}
	if value, ok := sc.mutation.PointBoost(); ok {
		_spec.SetField(service.FieldPointBoost, field.TypeInt, value)
		_node.PointBoost = value
	}
	if value, ok := sc.mutation.RoundUnits(); ok {
		_spec.SetField(service.FieldRoundUnits, field.TypeInt, value)
		_node.RoundUnits = value
	}
	if value, ok := sc.mutation.RoundDelay(); ok {
		_spec.SetField(service.FieldRoundDelay, field.TypeInt, value)
		_node.RoundDelay = value
	}
	if nodes := sc.mutation.CompetitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   service.CompetitionTable,
			Columns: []string{service.CompetitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competition.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CompetitionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   service.TeamTable,
			Columns: []string{service.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TeamID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.HostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   service.HostsTable,
			Columns: []string{service.HostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.host_services = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ChecksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   service.ChecksTable,
			Columns: []string{service.ChecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(check.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.PropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   service.PropertiesTable,
			Columns: []string{service.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(property.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ServiceCreateBulk is the builder for creating many Service entities in bulk.
type ServiceCreateBulk struct {
	config
	builders []*ServiceCreate
}

// Save creates the Service entities in the database.
func (scb *ServiceCreateBulk) Save(ctx context.Context) ([]*Service, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Service, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServiceMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ServiceCreateBulk) SaveX(ctx context.Context) []*Service {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ServiceCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ServiceCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
