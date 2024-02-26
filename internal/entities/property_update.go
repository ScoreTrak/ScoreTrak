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
	"github.com/scoretrak/scoretrak/internal/entities/predicate"
	"github.com/scoretrak/scoretrak/internal/entities/property"
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

// SetUpdateTime sets the "update_time" field.
func (pu *PropertyUpdate) SetUpdateTime(t time.Time) *PropertyUpdate {
	pu.mutation.SetUpdateTime(t)
	return pu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (pu *PropertyUpdate) ClearUpdateTime() *PropertyUpdate {
	pu.mutation.ClearUpdateTime()
	return pu
}

// SetValue sets the "value" field.
func (pu *PropertyUpdate) SetValue(s string) *PropertyUpdate {
	pu.mutation.SetValue(s)
	return pu
}

// Mutation returns the PropertyMutation object of the builder.
func (pu *PropertyUpdate) Mutation() *PropertyMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PropertyUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
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

// defaults sets the default values of the builder before save.
func (pu *PropertyUpdate) defaults() {
	if _, ok := pu.mutation.UpdateTime(); !ok && !pu.mutation.UpdateTimeCleared() {
		v := property.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PropertyUpdate) check() error {
	if _, ok := pu.mutation.HostserviceID(); pu.mutation.HostserviceCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Property.hostservice"`)
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
	if pu.mutation.CreateTimeCleared() {
		_spec.ClearField(property.FieldCreateTime, field.TypeTime)
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.SetField(property.FieldUpdateTime, field.TypeTime, value)
	}
	if pu.mutation.UpdateTimeCleared() {
		_spec.ClearField(property.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := pu.mutation.Value(); ok {
		_spec.SetField(property.FieldValue, field.TypeString, value)
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

// SetUpdateTime sets the "update_time" field.
func (puo *PropertyUpdateOne) SetUpdateTime(t time.Time) *PropertyUpdateOne {
	puo.mutation.SetUpdateTime(t)
	return puo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (puo *PropertyUpdateOne) ClearUpdateTime() *PropertyUpdateOne {
	puo.mutation.ClearUpdateTime()
	return puo
}

// SetValue sets the "value" field.
func (puo *PropertyUpdateOne) SetValue(s string) *PropertyUpdateOne {
	puo.mutation.SetValue(s)
	return puo
}

// Mutation returns the PropertyMutation object of the builder.
func (puo *PropertyUpdateOne) Mutation() *PropertyMutation {
	return puo.mutation
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
	puo.defaults()
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

// defaults sets the default values of the builder before save.
func (puo *PropertyUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdateTime(); !ok && !puo.mutation.UpdateTimeCleared() {
		v := property.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PropertyUpdateOne) check() error {
	if _, ok := puo.mutation.HostserviceID(); puo.mutation.HostserviceCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Property.hostservice"`)
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
	if puo.mutation.CreateTimeCleared() {
		_spec.ClearField(property.FieldCreateTime, field.TypeTime)
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.SetField(property.FieldUpdateTime, field.TypeTime, value)
	}
	if puo.mutation.UpdateTimeCleared() {
		_spec.ClearField(property.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := puo.mutation.Value(); ok {
		_spec.SetField(property.FieldValue, field.TypeString, value)
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
