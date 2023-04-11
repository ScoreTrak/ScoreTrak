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
	"github.com/ScoreTrak/ScoreTrak/internal/entities/round"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/service"
)

// CheckCreate is the builder for creating a Check entity.
type CheckCreate struct {
	config
	mutation *CheckMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (cc *CheckCreate) SetCreateTime(t time.Time) *CheckCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *CheckCreate) SetNillableCreateTime(t *time.Time) *CheckCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *CheckCreate) SetUpdateTime(t time.Time) *CheckCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *CheckCreate) SetNillableUpdateTime(t *time.Time) *CheckCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetPause sets the "pause" field.
func (cc *CheckCreate) SetPause(b bool) *CheckCreate {
	cc.mutation.SetPause(b)
	return cc
}

// SetHidden sets the "hidden" field.
func (cc *CheckCreate) SetHidden(b bool) *CheckCreate {
	cc.mutation.SetHidden(b)
	return cc
}

// SetCompetitionID sets the "competition_id" field.
func (cc *CheckCreate) SetCompetitionID(i int) *CheckCreate {
	cc.mutation.SetCompetitionID(i)
	return cc
}

// SetLog sets the "log" field.
func (cc *CheckCreate) SetLog(s string) *CheckCreate {
	cc.mutation.SetLog(s)
	return cc
}

// SetError sets the "error" field.
func (cc *CheckCreate) SetError(s string) *CheckCreate {
	cc.mutation.SetError(s)
	return cc
}

// SetPassed sets the "passed" field.
func (cc *CheckCreate) SetPassed(b bool) *CheckCreate {
	cc.mutation.SetPassed(b)
	return cc
}

// SetCompetition sets the "competition" edge to the Competition entity.
func (cc *CheckCreate) SetCompetition(c *Competition) *CheckCreate {
	return cc.SetCompetitionID(c.ID)
}

// SetRoundsID sets the "rounds" edge to the Round entity by ID.
func (cc *CheckCreate) SetRoundsID(id int) *CheckCreate {
	cc.mutation.SetRoundsID(id)
	return cc
}

// SetRounds sets the "rounds" edge to the Round entity.
func (cc *CheckCreate) SetRounds(r *Round) *CheckCreate {
	return cc.SetRoundsID(r.ID)
}

// SetServicesID sets the "services" edge to the Service entity by ID.
func (cc *CheckCreate) SetServicesID(id int) *CheckCreate {
	cc.mutation.SetServicesID(id)
	return cc
}

// SetServices sets the "services" edge to the Service entity.
func (cc *CheckCreate) SetServices(s *Service) *CheckCreate {
	return cc.SetServicesID(s.ID)
}

// Mutation returns the CheckMutation object of the builder.
func (cc *CheckCreate) Mutation() *CheckMutation {
	return cc.mutation
}

// Save creates the Check in the database.
func (cc *CheckCreate) Save(ctx context.Context) (*Check, error) {
	cc.defaults()
	return withHooks[*Check, CheckMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CheckCreate) SaveX(ctx context.Context) *Check {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CheckCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CheckCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CheckCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := check.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := check.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CheckCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`entities: missing required field "Check.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`entities: missing required field "Check.update_time"`)}
	}
	if _, ok := cc.mutation.Pause(); !ok {
		return &ValidationError{Name: "pause", err: errors.New(`entities: missing required field "Check.pause"`)}
	}
	if _, ok := cc.mutation.Hidden(); !ok {
		return &ValidationError{Name: "hidden", err: errors.New(`entities: missing required field "Check.hidden"`)}
	}
	if _, ok := cc.mutation.CompetitionID(); !ok {
		return &ValidationError{Name: "competition_id", err: errors.New(`entities: missing required field "Check.competition_id"`)}
	}
	if _, ok := cc.mutation.Log(); !ok {
		return &ValidationError{Name: "log", err: errors.New(`entities: missing required field "Check.log"`)}
	}
	if _, ok := cc.mutation.Error(); !ok {
		return &ValidationError{Name: "error", err: errors.New(`entities: missing required field "Check.error"`)}
	}
	if _, ok := cc.mutation.Passed(); !ok {
		return &ValidationError{Name: "passed", err: errors.New(`entities: missing required field "Check.passed"`)}
	}
	if _, ok := cc.mutation.CompetitionID(); !ok {
		return &ValidationError{Name: "competition", err: errors.New(`entities: missing required edge "Check.competition"`)}
	}
	if _, ok := cc.mutation.RoundsID(); !ok {
		return &ValidationError{Name: "rounds", err: errors.New(`entities: missing required edge "Check.rounds"`)}
	}
	if _, ok := cc.mutation.ServicesID(); !ok {
		return &ValidationError{Name: "services", err: errors.New(`entities: missing required edge "Check.services"`)}
	}
	return nil
}

func (cc *CheckCreate) sqlSave(ctx context.Context) (*Check, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CheckCreate) createSpec() (*Check, *sqlgraph.CreateSpec) {
	var (
		_node = &Check{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(check.Table, sqlgraph.NewFieldSpec(check.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.SetField(check.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.SetField(check.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.Pause(); ok {
		_spec.SetField(check.FieldPause, field.TypeBool, value)
		_node.Pause = value
	}
	if value, ok := cc.mutation.Hidden(); ok {
		_spec.SetField(check.FieldHidden, field.TypeBool, value)
		_node.Hidden = value
	}
	if value, ok := cc.mutation.Log(); ok {
		_spec.SetField(check.FieldLog, field.TypeString, value)
		_node.Log = value
	}
	if value, ok := cc.mutation.Error(); ok {
		_spec.SetField(check.FieldError, field.TypeString, value)
		_node.Error = value
	}
	if value, ok := cc.mutation.Passed(); ok {
		_spec.SetField(check.FieldPassed, field.TypeBool, value)
		_node.Passed = value
	}
	if nodes := cc.mutation.CompetitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   check.CompetitionTable,
			Columns: []string{check.CompetitionColumn},
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
	if nodes := cc.mutation.RoundsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   check.RoundsTable,
			Columns: []string{check.RoundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.round_checks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   check.ServicesTable,
			Columns: []string{check.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.service_checks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CheckCreateBulk is the builder for creating many Check entities in bulk.
type CheckCreateBulk struct {
	config
	builders []*CheckCreate
}

// Save creates the Check entities in the database.
func (ccb *CheckCreateBulk) Save(ctx context.Context) ([]*Check, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Check, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CheckMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CheckCreateBulk) SaveX(ctx context.Context) []*Check {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CheckCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CheckCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
