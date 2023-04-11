// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/competition"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/property"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/service"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/team"
)

// PropertyCreate is the builder for creating a Property entity.
type PropertyCreate struct {
	config
	mutation *PropertyMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *PropertyCreate) SetCreateTime(t time.Time) *PropertyCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PropertyCreate) SetNillableCreateTime(t *time.Time) *PropertyCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PropertyCreate) SetUpdateTime(t time.Time) *PropertyCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PropertyCreate) SetNillableUpdateTime(t *time.Time) *PropertyCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetCompetitionID sets the "competition_id" field.
func (pc *PropertyCreate) SetCompetitionID(i int) *PropertyCreate {
	pc.mutation.SetCompetitionID(i)
	return pc
}

// SetTeamID sets the "team_id" field.
func (pc *PropertyCreate) SetTeamID(i int) *PropertyCreate {
	pc.mutation.SetTeamID(i)
	return pc
}

// SetKey sets the "key" field.
func (pc *PropertyCreate) SetKey(s string) *PropertyCreate {
	pc.mutation.SetKey(s)
	return pc
}

// SetValue sets the "value" field.
func (pc *PropertyCreate) SetValue(s string) *PropertyCreate {
	pc.mutation.SetValue(s)
	return pc
}

// SetStatus sets the "status" field.
func (pc *PropertyCreate) SetStatus(pr property.Status) *PropertyCreate {
	pc.mutation.SetStatus(pr)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *PropertyCreate) SetNillableStatus(pr *property.Status) *PropertyCreate {
	if pr != nil {
		pc.SetStatus(*pr)
	}
	return pc
}

// SetCompetition sets the "competition" edge to the Competition entity.
func (pc *PropertyCreate) SetCompetition(c *Competition) *PropertyCreate {
	return pc.SetCompetitionID(c.ID)
}

// SetTeam sets the "team" edge to the Team entity.
func (pc *PropertyCreate) SetTeam(t *Team) *PropertyCreate {
	return pc.SetTeamID(t.ID)
}

// SetServicesID sets the "services" edge to the Service entity by ID.
func (pc *PropertyCreate) SetServicesID(id int) *PropertyCreate {
	pc.mutation.SetServicesID(id)
	return pc
}

// SetServices sets the "services" edge to the Service entity.
func (pc *PropertyCreate) SetServices(s *Service) *PropertyCreate {
	return pc.SetServicesID(s.ID)
}

// Mutation returns the PropertyMutation object of the builder.
func (pc *PropertyCreate) Mutation() *PropertyMutation {
	return pc.mutation
}

// Save creates the Property in the database.
func (pc *PropertyCreate) Save(ctx context.Context) (*Property, error) {
	pc.defaults()
	return withHooks[*Property, PropertyMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PropertyCreate) SaveX(ctx context.Context) *Property {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PropertyCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PropertyCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PropertyCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := property.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := property.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.Status(); !ok {
		v := property.DefaultStatus
		pc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PropertyCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`entities: missing required field "Property.create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`entities: missing required field "Property.update_time"`)}
	}
	if _, ok := pc.mutation.CompetitionID(); !ok {
		return &ValidationError{Name: "competition_id", err: errors.New(`entities: missing required field "Property.competition_id"`)}
	}
	if _, ok := pc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team_id", err: errors.New(`entities: missing required field "Property.team_id"`)}
	}
	if _, ok := pc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`entities: missing required field "Property.key"`)}
	}
	if _, ok := pc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`entities: missing required field "Property.value"`)}
	}
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`entities: missing required field "Property.status"`)}
	}
	if v, ok := pc.mutation.Status(); ok {
		if err := property.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`entities: validator failed for field "Property.status": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CompetitionID(); !ok {
		return &ValidationError{Name: "competition", err: errors.New(`entities: missing required edge "Property.competition"`)}
	}
	if _, ok := pc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team", err: errors.New(`entities: missing required edge "Property.team"`)}
	}
	if _, ok := pc.mutation.ServicesID(); !ok {
		return &ValidationError{Name: "services", err: errors.New(`entities: missing required edge "Property.services"`)}
	}
	return nil
}

func (pc *PropertyCreate) sqlSave(ctx context.Context) (*Property, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PropertyCreate) createSpec() (*Property, *sqlgraph.CreateSpec) {
	var (
		_node = &Property{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(property.Table, sqlgraph.NewFieldSpec(property.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(property.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(property.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Key(); ok {
		_spec.SetField(property.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := pc.mutation.Value(); ok {
		_spec.SetField(property.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(property.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := pc.mutation.CompetitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   property.CompetitionTable,
			Columns: []string{property.CompetitionColumn},
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
	if nodes := pc.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   property.TeamTable,
			Columns: []string{property.TeamColumn},
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
	if nodes := pc.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   property.ServicesTable,
			Columns: []string{property.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.service_properties = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PropertyCreateBulk is the builder for creating many Property entities in bulk.
type PropertyCreateBulk struct {
	config
	builders []*PropertyCreate
}

// Save creates the Property entities in the database.
func (pcb *PropertyCreateBulk) Save(ctx context.Context) ([]*Property, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Property, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PropertyMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PropertyCreateBulk) SaveX(ctx context.Context) []*Property {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PropertyCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PropertyCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
