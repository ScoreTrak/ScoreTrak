// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/predicate"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/service"
)

// ServiceUpdate is the builder for updating Service entities.
type ServiceUpdate struct {
	config
	hooks    []Hook
	mutation *ServiceMutation
}

// Where appends a list predicates to the ServiceUpdate builder.
func (su *ServiceUpdate) Where(ps ...predicate.Service) *ServiceUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *ServiceUpdate) SetName(s string) *ServiceUpdate {
	su.mutation.SetName(s)
	return su
}

// SetDisplayName sets the "display_name" field.
func (su *ServiceUpdate) SetDisplayName(s string) *ServiceUpdate {
	su.mutation.SetDisplayName(s)
	return su
}

// SetPause sets the "pause" field.
func (su *ServiceUpdate) SetPause(b bool) *ServiceUpdate {
	su.mutation.SetPause(b)
	return su
}

// SetNillablePause sets the "pause" field if the given value is not nil.
func (su *ServiceUpdate) SetNillablePause(b *bool) *ServiceUpdate {
	if b != nil {
		su.SetPause(*b)
	}
	return su
}

// ClearPause clears the value of the "pause" field.
func (su *ServiceUpdate) ClearPause() *ServiceUpdate {
	su.mutation.ClearPause()
	return su
}

// SetHidden sets the "hidden" field.
func (su *ServiceUpdate) SetHidden(b bool) *ServiceUpdate {
	su.mutation.SetHidden(b)
	return su
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (su *ServiceUpdate) SetNillableHidden(b *bool) *ServiceUpdate {
	if b != nil {
		su.SetHidden(*b)
	}
	return su
}

// ClearHidden clears the value of the "hidden" field.
func (su *ServiceUpdate) ClearHidden() *ServiceUpdate {
	su.mutation.ClearHidden()
	return su
}

// Mutation returns the ServiceMutation object of the builder.
func (su *ServiceUpdate) Mutation() *ServiceMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ServiceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ServiceMutation](ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *ServiceUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ServiceUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ServiceUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ServiceUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := service.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entities: validator failed for field "Service.name": %w`, err)}
		}
	}
	if v, ok := su.mutation.DisplayName(); ok {
		if err := service.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`entities: validator failed for field "Service.display_name": %w`, err)}
		}
	}
	if _, ok := su.mutation.CompetitionID(); su.mutation.CompetitionCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Service.competition"`)
	}
	return nil
}

func (su *ServiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(service.Table, service.Columns, sqlgraph.NewFieldSpec(service.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(service.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.DisplayName(); ok {
		_spec.SetField(service.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := su.mutation.Pause(); ok {
		_spec.SetField(service.FieldPause, field.TypeBool, value)
	}
	if su.mutation.PauseCleared() {
		_spec.ClearField(service.FieldPause, field.TypeBool)
	}
	if value, ok := su.mutation.Hidden(); ok {
		_spec.SetField(service.FieldHidden, field.TypeBool, value)
	}
	if su.mutation.HiddenCleared() {
		_spec.ClearField(service.FieldHidden, field.TypeBool)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{service.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// ServiceUpdateOne is the builder for updating a single Service entity.
type ServiceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServiceMutation
}

// SetName sets the "name" field.
func (suo *ServiceUpdateOne) SetName(s string) *ServiceUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetDisplayName sets the "display_name" field.
func (suo *ServiceUpdateOne) SetDisplayName(s string) *ServiceUpdateOne {
	suo.mutation.SetDisplayName(s)
	return suo
}

// SetPause sets the "pause" field.
func (suo *ServiceUpdateOne) SetPause(b bool) *ServiceUpdateOne {
	suo.mutation.SetPause(b)
	return suo
}

// SetNillablePause sets the "pause" field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillablePause(b *bool) *ServiceUpdateOne {
	if b != nil {
		suo.SetPause(*b)
	}
	return suo
}

// ClearPause clears the value of the "pause" field.
func (suo *ServiceUpdateOne) ClearPause() *ServiceUpdateOne {
	suo.mutation.ClearPause()
	return suo
}

// SetHidden sets the "hidden" field.
func (suo *ServiceUpdateOne) SetHidden(b bool) *ServiceUpdateOne {
	suo.mutation.SetHidden(b)
	return suo
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableHidden(b *bool) *ServiceUpdateOne {
	if b != nil {
		suo.SetHidden(*b)
	}
	return suo
}

// ClearHidden clears the value of the "hidden" field.
func (suo *ServiceUpdateOne) ClearHidden() *ServiceUpdateOne {
	suo.mutation.ClearHidden()
	return suo
}

// Mutation returns the ServiceMutation object of the builder.
func (suo *ServiceUpdateOne) Mutation() *ServiceMutation {
	return suo.mutation
}

// Where appends a list predicates to the ServiceUpdate builder.
func (suo *ServiceUpdateOne) Where(ps ...predicate.Service) *ServiceUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ServiceUpdateOne) Select(field string, fields ...string) *ServiceUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Service entity.
func (suo *ServiceUpdateOne) Save(ctx context.Context) (*Service, error) {
	return withHooks[*Service, ServiceMutation](ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ServiceUpdateOne) SaveX(ctx context.Context) *Service {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ServiceUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ServiceUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ServiceUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := service.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entities: validator failed for field "Service.name": %w`, err)}
		}
	}
	if v, ok := suo.mutation.DisplayName(); ok {
		if err := service.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`entities: validator failed for field "Service.display_name": %w`, err)}
		}
	}
	if _, ok := suo.mutation.CompetitionID(); suo.mutation.CompetitionCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Service.competition"`)
	}
	return nil
}

func (suo *ServiceUpdateOne) sqlSave(ctx context.Context) (_node *Service, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(service.Table, service.Columns, sqlgraph.NewFieldSpec(service.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "Service.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, service.FieldID)
		for _, f := range fields {
			if !service.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != service.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(service.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.DisplayName(); ok {
		_spec.SetField(service.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Pause(); ok {
		_spec.SetField(service.FieldPause, field.TypeBool, value)
	}
	if suo.mutation.PauseCleared() {
		_spec.ClearField(service.FieldPause, field.TypeBool)
	}
	if value, ok := suo.mutation.Hidden(); ok {
		_spec.SetField(service.FieldHidden, field.TypeBool, value)
	}
	if suo.mutation.HiddenCleared() {
		_spec.ClearField(service.FieldHidden, field.TypeBool)
	}
	_node = &Service{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{service.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
