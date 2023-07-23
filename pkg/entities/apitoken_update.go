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
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/apitoken"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/predicate"
)

// ApiTokenUpdate is the builder for updating ApiToken entities.
type ApiTokenUpdate struct {
	config
	hooks    []Hook
	mutation *ApiTokenMutation
}

// Where appends a list predicates to the ApiTokenUpdate builder.
func (atu *ApiTokenUpdate) Where(ps ...predicate.ApiToken) *ApiTokenUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetUpdateTime sets the "update_time" field.
func (atu *ApiTokenUpdate) SetUpdateTime(t time.Time) *ApiTokenUpdate {
	atu.mutation.SetUpdateTime(t)
	return atu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (atu *ApiTokenUpdate) ClearUpdateTime() *ApiTokenUpdate {
	atu.mutation.ClearUpdateTime()
	return atu
}

// SetToken sets the "token" field.
func (atu *ApiTokenUpdate) SetToken(s string) *ApiTokenUpdate {
	atu.mutation.SetToken(s)
	return atu
}

// SetExpiredAt sets the "expired_at" field.
func (atu *ApiTokenUpdate) SetExpiredAt(t time.Time) *ApiTokenUpdate {
	atu.mutation.SetExpiredAt(t)
	return atu
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (atu *ApiTokenUpdate) SetNillableExpiredAt(t *time.Time) *ApiTokenUpdate {
	if t != nil {
		atu.SetExpiredAt(*t)
	}
	return atu
}

// Mutation returns the ApiTokenMutation object of the builder.
func (atu *ApiTokenUpdate) Mutation() *ApiTokenMutation {
	return atu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *ApiTokenUpdate) Save(ctx context.Context) (int, error) {
	atu.defaults()
	return withHooks[int, ApiTokenMutation](ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *ApiTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *ApiTokenUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *ApiTokenUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atu *ApiTokenUpdate) defaults() {
	if _, ok := atu.mutation.UpdateTime(); !ok && !atu.mutation.UpdateTimeCleared() {
		v := apitoken.UpdateDefaultUpdateTime()
		atu.mutation.SetUpdateTime(v)
	}
}

func (atu *ApiTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(apitoken.Table, apitoken.Columns, sqlgraph.NewFieldSpec(apitoken.FieldID, field.TypeString))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atu.mutation.CreateTimeCleared() {
		_spec.ClearField(apitoken.FieldCreateTime, field.TypeTime)
	}
	if value, ok := atu.mutation.UpdateTime(); ok {
		_spec.SetField(apitoken.FieldUpdateTime, field.TypeTime, value)
	}
	if atu.mutation.UpdateTimeCleared() {
		_spec.ClearField(apitoken.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := atu.mutation.Token(); ok {
		_spec.SetField(apitoken.FieldToken, field.TypeString, value)
	}
	if value, ok := atu.mutation.ExpiredAt(); ok {
		_spec.SetField(apitoken.FieldExpiredAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apitoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// ApiTokenUpdateOne is the builder for updating a single ApiToken entity.
type ApiTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApiTokenMutation
}

// SetUpdateTime sets the "update_time" field.
func (atuo *ApiTokenUpdateOne) SetUpdateTime(t time.Time) *ApiTokenUpdateOne {
	atuo.mutation.SetUpdateTime(t)
	return atuo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (atuo *ApiTokenUpdateOne) ClearUpdateTime() *ApiTokenUpdateOne {
	atuo.mutation.ClearUpdateTime()
	return atuo
}

// SetToken sets the "token" field.
func (atuo *ApiTokenUpdateOne) SetToken(s string) *ApiTokenUpdateOne {
	atuo.mutation.SetToken(s)
	return atuo
}

// SetExpiredAt sets the "expired_at" field.
func (atuo *ApiTokenUpdateOne) SetExpiredAt(t time.Time) *ApiTokenUpdateOne {
	atuo.mutation.SetExpiredAt(t)
	return atuo
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (atuo *ApiTokenUpdateOne) SetNillableExpiredAt(t *time.Time) *ApiTokenUpdateOne {
	if t != nil {
		atuo.SetExpiredAt(*t)
	}
	return atuo
}

// Mutation returns the ApiTokenMutation object of the builder.
func (atuo *ApiTokenUpdateOne) Mutation() *ApiTokenMutation {
	return atuo.mutation
}

// Where appends a list predicates to the ApiTokenUpdate builder.
func (atuo *ApiTokenUpdateOne) Where(ps ...predicate.ApiToken) *ApiTokenUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *ApiTokenUpdateOne) Select(field string, fields ...string) *ApiTokenUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated ApiToken entity.
func (atuo *ApiTokenUpdateOne) Save(ctx context.Context) (*ApiToken, error) {
	atuo.defaults()
	return withHooks[*ApiToken, ApiTokenMutation](ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *ApiTokenUpdateOne) SaveX(ctx context.Context) *ApiToken {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *ApiTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *ApiTokenUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atuo *ApiTokenUpdateOne) defaults() {
	if _, ok := atuo.mutation.UpdateTime(); !ok && !atuo.mutation.UpdateTimeCleared() {
		v := apitoken.UpdateDefaultUpdateTime()
		atuo.mutation.SetUpdateTime(v)
	}
}

func (atuo *ApiTokenUpdateOne) sqlSave(ctx context.Context) (_node *ApiToken, err error) {
	_spec := sqlgraph.NewUpdateSpec(apitoken.Table, apitoken.Columns, sqlgraph.NewFieldSpec(apitoken.FieldID, field.TypeString))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "ApiToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apitoken.FieldID)
		for _, f := range fields {
			if !apitoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != apitoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atuo.mutation.CreateTimeCleared() {
		_spec.ClearField(apitoken.FieldCreateTime, field.TypeTime)
	}
	if value, ok := atuo.mutation.UpdateTime(); ok {
		_spec.SetField(apitoken.FieldUpdateTime, field.TypeTime, value)
	}
	if atuo.mutation.UpdateTimeCleared() {
		_spec.ClearField(apitoken.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := atuo.mutation.Token(); ok {
		_spec.SetField(apitoken.FieldToken, field.TypeString, value)
	}
	if value, ok := atuo.mutation.ExpiredAt(); ok {
		_spec.SetField(apitoken.FieldExpiredAt, field.TypeTime, value)
	}
	_node = &ApiToken{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apitoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}
