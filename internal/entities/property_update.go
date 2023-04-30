// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/hostservice"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/predicate"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/property"
)

// PropertyUpdate is the builder for updating Property entities.
type PropertyUpdate struct {
	config
	hooks    []Hook
	mutation *PropertyMutation
}

// Where appends a list predicates to the PropertyUpdate builder.
func (pu *PropertyUpdate) Where(ps ...predicate.Property) *PropertyUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetKey sets the "key" field.
func (pu *PropertyUpdate) SetKey(s string) *PropertyUpdate {
	pu.mutation.SetKey(s)
	return pu
}

// SetValue sets the "value" field.
func (pu *PropertyUpdate) SetValue(s string) *PropertyUpdate {
	pu.mutation.SetValue(s)
	return pu
}

// SetStatus sets the "status" field.
func (pu *PropertyUpdate) SetStatus(pr property.Status) *PropertyUpdate {
	pu.mutation.SetStatus(pr)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableStatus(pr *property.Status) *PropertyUpdate {
	if pr != nil {
		pu.SetStatus(*pr)
	}
	return pu
}

// SetHostServiceID sets the "host_service_id" field.
func (pu *PropertyUpdate) SetHostServiceID(s string) *PropertyUpdate {
	pu.mutation.SetHostServiceID(s)
	return pu
}

// SetHostserviceID sets the "hostservice" edge to the HostService entity by ID.
func (pu *PropertyUpdate) SetHostserviceID(id string) *PropertyUpdate {
	pu.mutation.SetHostserviceID(id)
	return pu
}

// SetHostservice sets the "hostservice" edge to the HostService entity.
func (pu *PropertyUpdate) SetHostservice(h *HostService) *PropertyUpdate {
	return pu.SetHostserviceID(h.ID)
}

// Mutation returns the PropertyMutation object of the builder.
func (pu *PropertyUpdate) Mutation() *PropertyMutation {
	return pu.mutation
}

// ClearHostservice clears the "hostservice" edge to the HostService entity.
func (pu *PropertyUpdate) ClearHostservice() *PropertyUpdate {
	pu.mutation.ClearHostservice()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PropertyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PropertyMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PropertyUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PropertyUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PropertyUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PropertyUpdate) check() error {
	if v, ok := pu.mutation.Status(); ok {
		if err := property.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`entities: validator failed for field "Property.status": %w`, err)}
		}
	}
	if _, ok := pu.mutation.HostserviceID(); pu.mutation.HostserviceCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Property.hostservice"`)
	}
	if _, ok := pu.mutation.TeamID(); pu.mutation.TeamCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Property.team"`)
	}
	return nil
}

func (pu *PropertyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(property.Table, property.Columns, sqlgraph.NewFieldSpec(property.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Key(); ok {
		_spec.SetField(property.FieldKey, field.TypeString, value)
	}
	if value, ok := pu.mutation.Value(); ok {
		_spec.SetField(property.FieldValue, field.TypeString, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(property.FieldStatus, field.TypeEnum, value)
	}
	if pu.mutation.HostserviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   property.HostserviceTable,
			Columns: []string{property.HostserviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostservice.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.HostserviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   property.HostserviceTable,
			Columns: []string{property.HostserviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostservice.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{property.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PropertyUpdateOne is the builder for updating a single Property entity.
type PropertyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PropertyMutation
}

// SetKey sets the "key" field.
func (puo *PropertyUpdateOne) SetKey(s string) *PropertyUpdateOne {
	puo.mutation.SetKey(s)
	return puo
}

// SetValue sets the "value" field.
func (puo *PropertyUpdateOne) SetValue(s string) *PropertyUpdateOne {
	puo.mutation.SetValue(s)
	return puo
}

// SetStatus sets the "status" field.
func (puo *PropertyUpdateOne) SetStatus(pr property.Status) *PropertyUpdateOne {
	puo.mutation.SetStatus(pr)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableStatus(pr *property.Status) *PropertyUpdateOne {
	if pr != nil {
		puo.SetStatus(*pr)
	}
	return puo
}

// SetHostServiceID sets the "host_service_id" field.
func (puo *PropertyUpdateOne) SetHostServiceID(s string) *PropertyUpdateOne {
	puo.mutation.SetHostServiceID(s)
	return puo
}

// SetHostserviceID sets the "hostservice" edge to the HostService entity by ID.
func (puo *PropertyUpdateOne) SetHostserviceID(id string) *PropertyUpdateOne {
	puo.mutation.SetHostserviceID(id)
	return puo
}

// SetHostservice sets the "hostservice" edge to the HostService entity.
func (puo *PropertyUpdateOne) SetHostservice(h *HostService) *PropertyUpdateOne {
	return puo.SetHostserviceID(h.ID)
}

// Mutation returns the PropertyMutation object of the builder.
func (puo *PropertyUpdateOne) Mutation() *PropertyMutation {
	return puo.mutation
}

// ClearHostservice clears the "hostservice" edge to the HostService entity.
func (puo *PropertyUpdateOne) ClearHostservice() *PropertyUpdateOne {
	puo.mutation.ClearHostservice()
	return puo
}

// Where appends a list predicates to the PropertyUpdate builder.
func (puo *PropertyUpdateOne) Where(ps ...predicate.Property) *PropertyUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PropertyUpdateOne) Select(field string, fields ...string) *PropertyUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Property entity.
func (puo *PropertyUpdateOne) Save(ctx context.Context) (*Property, error) {
	return withHooks[*Property, PropertyMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PropertyUpdateOne) SaveX(ctx context.Context) *Property {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PropertyUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PropertyUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PropertyUpdateOne) check() error {
	if v, ok := puo.mutation.Status(); ok {
		if err := property.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`entities: validator failed for field "Property.status": %w`, err)}
		}
	}
	if _, ok := puo.mutation.HostserviceID(); puo.mutation.HostserviceCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Property.hostservice"`)
	}
	if _, ok := puo.mutation.TeamID(); puo.mutation.TeamCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Property.team"`)
	}
	return nil
}

func (puo *PropertyUpdateOne) sqlSave(ctx context.Context) (_node *Property, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(property.Table, property.Columns, sqlgraph.NewFieldSpec(property.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "Property.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, property.FieldID)
		for _, f := range fields {
			if !property.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != property.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Key(); ok {
		_spec.SetField(property.FieldKey, field.TypeString, value)
	}
	if value, ok := puo.mutation.Value(); ok {
		_spec.SetField(property.FieldValue, field.TypeString, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(property.FieldStatus, field.TypeEnum, value)
	}
	if puo.mutation.HostserviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   property.HostserviceTable,
			Columns: []string{property.HostserviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostservice.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.HostserviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   property.HostserviceTable,
			Columns: []string{property.HostserviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostservice.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Property{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{property.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
