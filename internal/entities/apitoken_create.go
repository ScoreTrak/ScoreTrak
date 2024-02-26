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
	"github.com/scoretrak/scoretrak/internal/entities/apitoken"
)

// ApiTokenCreate is the builder for creating a ApiToken entity.
type ApiTokenCreate struct {
	config
	mutation *ApiTokenMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (atc *ApiTokenCreate) SetCreateTime(t time.Time) *ApiTokenCreate {
	atc.mutation.SetCreateTime(t)
	return atc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableCreateTime(t *time.Time) *ApiTokenCreate {
	if t != nil {
		atc.SetCreateTime(*t)
	}
	return atc
}

// SetUpdateTime sets the "update_time" field.
func (atc *ApiTokenCreate) SetUpdateTime(t time.Time) *ApiTokenCreate {
	atc.mutation.SetUpdateTime(t)
	return atc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableUpdateTime(t *time.Time) *ApiTokenCreate {
	if t != nil {
		atc.SetUpdateTime(*t)
	}
	return atc
}

// SetToken sets the "token" field.
func (atc *ApiTokenCreate) SetToken(s string) *ApiTokenCreate {
	atc.mutation.SetToken(s)
	return atc
}

// SetExpiredAt sets the "expired_at" field.
func (atc *ApiTokenCreate) SetExpiredAt(t time.Time) *ApiTokenCreate {
	atc.mutation.SetExpiredAt(t)
	return atc
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableExpiredAt(t *time.Time) *ApiTokenCreate {
	if t != nil {
		atc.SetExpiredAt(*t)
	}
	return atc
}

// SetID sets the "id" field.
func (atc *ApiTokenCreate) SetID(s string) *ApiTokenCreate {
	atc.mutation.SetID(s)
	return atc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableID(s *string) *ApiTokenCreate {
	if s != nil {
		atc.SetID(*s)
	}
	return atc
}

// Mutation returns the ApiTokenMutation object of the builder.
func (atc *ApiTokenCreate) Mutation() *ApiTokenMutation {
	return atc.mutation
}

// Save creates the ApiToken in the database.
func (atc *ApiTokenCreate) Save(ctx context.Context) (*ApiToken, error) {
	atc.defaults()
	return withHooks[*ApiToken, ApiTokenMutation](ctx, atc.sqlSave, atc.mutation, atc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atc *ApiTokenCreate) SaveX(ctx context.Context) *ApiToken {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *ApiTokenCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *ApiTokenCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atc *ApiTokenCreate) defaults() {
	if _, ok := atc.mutation.CreateTime(); !ok {
		v := apitoken.DefaultCreateTime()
		atc.mutation.SetCreateTime(v)
	}
	if _, ok := atc.mutation.UpdateTime(); !ok {
		v := apitoken.DefaultUpdateTime()
		atc.mutation.SetUpdateTime(v)
	}
	if _, ok := atc.mutation.ExpiredAt(); !ok {
		v := apitoken.DefaultExpiredAt
		atc.mutation.SetExpiredAt(v)
	}
	if _, ok := atc.mutation.ID(); !ok {
		v := apitoken.DefaultID()
		atc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *ApiTokenCreate) check() error {
	if _, ok := atc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`entities: missing required field "ApiToken.token"`)}
	}
	if v, ok := atc.mutation.Token(); ok {
		if err := apitoken.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`entities: validator failed for field "ApiToken.token": %w`, err)}
		}
	}
	if _, ok := atc.mutation.ExpiredAt(); !ok {
		return &ValidationError{Name: "expired_at", err: errors.New(`entities: missing required field "ApiToken.expired_at"`)}
	}
	if v, ok := atc.mutation.ID(); ok {
		if err := apitoken.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`entities: validator failed for field "ApiToken.id": %w`, err)}
		}
	}
	return nil
}

func (atc *ApiTokenCreate) sqlSave(ctx context.Context) (*ApiToken, error) {
	if err := atc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ApiToken.ID type: %T", _spec.ID.Value)
		}
	}
	atc.mutation.id = &_node.ID
	atc.mutation.done = true
	return _node, nil
}

func (atc *ApiTokenCreate) createSpec() (*ApiToken, *sqlgraph.CreateSpec) {
	var (
		_node = &ApiToken{config: atc.config}
		_spec = sqlgraph.NewCreateSpec(apitoken.Table, sqlgraph.NewFieldSpec(apitoken.FieldID, field.TypeString))
	)
	_spec.OnConflict = atc.conflict
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := atc.mutation.CreateTime(); ok {
		_spec.SetField(apitoken.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := atc.mutation.UpdateTime(); ok {
		_spec.SetField(apitoken.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := atc.mutation.Token(); ok {
		_spec.SetField(apitoken.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := atc.mutation.ExpiredAt(); ok {
		_spec.SetField(apitoken.FieldExpiredAt, field.TypeTime, value)
		_node.ExpiredAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ApiToken.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ApiTokenUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (atc *ApiTokenCreate) OnConflict(opts ...sql.ConflictOption) *ApiTokenUpsertOne {
	atc.conflict = opts
	return &ApiTokenUpsertOne{
		create: atc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (atc *ApiTokenCreate) OnConflictColumns(columns ...string) *ApiTokenUpsertOne {
	atc.conflict = append(atc.conflict, sql.ConflictColumns(columns...))
	return &ApiTokenUpsertOne{
		create: atc,
	}
}

type (
	// ApiTokenUpsertOne is the builder for "upsert"-ing
	//  one ApiToken node.
	ApiTokenUpsertOne struct {
		create *ApiTokenCreate
	}

	// ApiTokenUpsert is the "OnConflict" setter.
	ApiTokenUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *ApiTokenUpsert) SetUpdateTime(v time.Time) *ApiTokenUpsert {
	u.Set(apitoken.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ApiTokenUpsert) UpdateUpdateTime() *ApiTokenUpsert {
	u.SetExcluded(apitoken.FieldUpdateTime)
	return u
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *ApiTokenUpsert) ClearUpdateTime() *ApiTokenUpsert {
	u.SetNull(apitoken.FieldUpdateTime)
	return u
}

// SetToken sets the "token" field.
func (u *ApiTokenUpsert) SetToken(v string) *ApiTokenUpsert {
	u.Set(apitoken.FieldToken, v)
	return u
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *ApiTokenUpsert) UpdateToken() *ApiTokenUpsert {
	u.SetExcluded(apitoken.FieldToken)
	return u
}

// SetExpiredAt sets the "expired_at" field.
func (u *ApiTokenUpsert) SetExpiredAt(v time.Time) *ApiTokenUpsert {
	u.Set(apitoken.FieldExpiredAt, v)
	return u
}

// UpdateExpiredAt sets the "expired_at" field to the value that was provided on create.
func (u *ApiTokenUpsert) UpdateExpiredAt() *ApiTokenUpsert {
	u.SetExcluded(apitoken.FieldExpiredAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apitoken.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ApiTokenUpsertOne) UpdateNewValues() *ApiTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(apitoken.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(apitoken.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ApiTokenUpsertOne) Ignore() *ApiTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ApiTokenUpsertOne) DoNothing() *ApiTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ApiTokenCreate.OnConflict
// documentation for more info.
func (u *ApiTokenUpsertOne) Update(set func(*ApiTokenUpsert)) *ApiTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ApiTokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ApiTokenUpsertOne) SetUpdateTime(v time.Time) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ApiTokenUpsertOne) UpdateUpdateTime() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *ApiTokenUpsertOne) ClearUpdateTime() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.ClearUpdateTime()
	})
}

// SetToken sets the "token" field.
func (u *ApiTokenUpsertOne) SetToken(v string) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *ApiTokenUpsertOne) UpdateToken() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateToken()
	})
}

// SetExpiredAt sets the "expired_at" field.
func (u *ApiTokenUpsertOne) SetExpiredAt(v time.Time) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetExpiredAt(v)
	})
}

// UpdateExpiredAt sets the "expired_at" field to the value that was provided on create.
func (u *ApiTokenUpsertOne) UpdateExpiredAt() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateExpiredAt()
	})
}

// Exec executes the query.
func (u *ApiTokenUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("entities: missing options for ApiTokenCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ApiTokenUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ApiTokenUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("entities: ApiTokenUpsertOne.ID is not supported by MySQL driver. Use ApiTokenUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ApiTokenUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ApiTokenCreateBulk is the builder for creating many ApiToken entities in bulk.
type ApiTokenCreateBulk struct {
	config
	builders []*ApiTokenCreate
	conflict []sql.ConflictOption
}

// Save creates the ApiToken entities in the database.
func (atcb *ApiTokenCreateBulk) Save(ctx context.Context) ([]*ApiToken, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*ApiToken, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApiTokenMutation)
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
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = atcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *ApiTokenCreateBulk) SaveX(ctx context.Context) []*ApiToken {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *ApiTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *ApiTokenCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ApiToken.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ApiTokenUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (atcb *ApiTokenCreateBulk) OnConflict(opts ...sql.ConflictOption) *ApiTokenUpsertBulk {
	atcb.conflict = opts
	return &ApiTokenUpsertBulk{
		create: atcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (atcb *ApiTokenCreateBulk) OnConflictColumns(columns ...string) *ApiTokenUpsertBulk {
	atcb.conflict = append(atcb.conflict, sql.ConflictColumns(columns...))
	return &ApiTokenUpsertBulk{
		create: atcb,
	}
}

// ApiTokenUpsertBulk is the builder for "upsert"-ing
// a bulk of ApiToken nodes.
type ApiTokenUpsertBulk struct {
	create *ApiTokenCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apitoken.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ApiTokenUpsertBulk) UpdateNewValues() *ApiTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(apitoken.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(apitoken.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ApiTokenUpsertBulk) Ignore() *ApiTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ApiTokenUpsertBulk) DoNothing() *ApiTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ApiTokenCreateBulk.OnConflict
// documentation for more info.
func (u *ApiTokenUpsertBulk) Update(set func(*ApiTokenUpsert)) *ApiTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ApiTokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ApiTokenUpsertBulk) SetUpdateTime(v time.Time) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ApiTokenUpsertBulk) UpdateUpdateTime() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *ApiTokenUpsertBulk) ClearUpdateTime() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.ClearUpdateTime()
	})
}

// SetToken sets the "token" field.
func (u *ApiTokenUpsertBulk) SetToken(v string) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *ApiTokenUpsertBulk) UpdateToken() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateToken()
	})
}

// SetExpiredAt sets the "expired_at" field.
func (u *ApiTokenUpsertBulk) SetExpiredAt(v time.Time) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetExpiredAt(v)
	})
}

// UpdateExpiredAt sets the "expired_at" field to the value that was provided on create.
func (u *ApiTokenUpsertBulk) UpdateExpiredAt() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateExpiredAt()
	})
}

// Exec executes the query.
func (u *ApiTokenUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("entities: OnConflict was set for builder %d. Set it on the ApiTokenCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("entities: missing options for ApiTokenCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ApiTokenUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
