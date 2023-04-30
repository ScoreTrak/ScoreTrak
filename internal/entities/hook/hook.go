// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/internal/entities"
)

// The CheckFunc type is an adapter to allow the use of ordinary
// function as Check mutator.
type CheckFunc func(context.Context, *entities.CheckMutation) (entities.Value, error)

// Mutate calls f(ctx, m).
func (f CheckFunc) Mutate(ctx context.Context, m entities.Mutation) (entities.Value, error) {
	if mv, ok := m.(*entities.CheckMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entities.CheckMutation", m)
}

// The CompetitionFunc type is an adapter to allow the use of ordinary
// function as Competition mutator.
type CompetitionFunc func(context.Context, *entities.CompetitionMutation) (entities.Value, error)

// Mutate calls f(ctx, m).
func (f CompetitionFunc) Mutate(ctx context.Context, m entities.Mutation) (entities.Value, error) {
	if mv, ok := m.(*entities.CompetitionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entities.CompetitionMutation", m)
}

// The HostFunc type is an adapter to allow the use of ordinary
// function as Host mutator.
type HostFunc func(context.Context, *entities.HostMutation) (entities.Value, error)

// Mutate calls f(ctx, m).
func (f HostFunc) Mutate(ctx context.Context, m entities.Mutation) (entities.Value, error) {
	if mv, ok := m.(*entities.HostMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entities.HostMutation", m)
}

// The HostServiceFunc type is an adapter to allow the use of ordinary
// function as HostService mutator.
type HostServiceFunc func(context.Context, *entities.HostServiceMutation) (entities.Value, error)

// Mutate calls f(ctx, m).
func (f HostServiceFunc) Mutate(ctx context.Context, m entities.Mutation) (entities.Value, error) {
	if mv, ok := m.(*entities.HostServiceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entities.HostServiceMutation", m)
}

// The PropertyFunc type is an adapter to allow the use of ordinary
// function as Property mutator.
type PropertyFunc func(context.Context, *entities.PropertyMutation) (entities.Value, error)

// Mutate calls f(ctx, m).
func (f PropertyFunc) Mutate(ctx context.Context, m entities.Mutation) (entities.Value, error) {
	if mv, ok := m.(*entities.PropertyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entities.PropertyMutation", m)
}

// The ReportFunc type is an adapter to allow the use of ordinary
// function as Report mutator.
type ReportFunc func(context.Context, *entities.ReportMutation) (entities.Value, error)

// Mutate calls f(ctx, m).
func (f ReportFunc) Mutate(ctx context.Context, m entities.Mutation) (entities.Value, error) {
	if mv, ok := m.(*entities.ReportMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entities.ReportMutation", m)
}

// The RoundFunc type is an adapter to allow the use of ordinary
// function as Round mutator.
type RoundFunc func(context.Context, *entities.RoundMutation) (entities.Value, error)

// Mutate calls f(ctx, m).
func (f RoundFunc) Mutate(ctx context.Context, m entities.Mutation) (entities.Value, error) {
	if mv, ok := m.(*entities.RoundMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entities.RoundMutation", m)
}

// The ServiceFunc type is an adapter to allow the use of ordinary
// function as Service mutator.
type ServiceFunc func(context.Context, *entities.ServiceMutation) (entities.Value, error)

// Mutate calls f(ctx, m).
func (f ServiceFunc) Mutate(ctx context.Context, m entities.Mutation) (entities.Value, error) {
	if mv, ok := m.(*entities.ServiceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entities.ServiceMutation", m)
}

// The TeamFunc type is an adapter to allow the use of ordinary
// function as Team mutator.
type TeamFunc func(context.Context, *entities.TeamMutation) (entities.Value, error)

// Mutate calls f(ctx, m).
func (f TeamFunc) Mutate(ctx context.Context, m entities.Mutation) (entities.Value, error) {
	if mv, ok := m.(*entities.TeamMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entities.TeamMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, entities.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m entities.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m entities.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m entities.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op entities.Op) Condition {
	return func(_ context.Context, m entities.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entities.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entities.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entities.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk entities.Hook, cond Condition) entities.Hook {
	return func(next entities.Mutator) entities.Mutator {
		return entities.MutateFunc(func(ctx context.Context, m entities.Mutation) (entities.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, entities.Delete|entities.Create)
func On(hk entities.Hook, op entities.Op) entities.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, entities.Update|entities.UpdateOne)
func Unless(hk entities.Hook, op entities.Op) entities.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) entities.Hook {
	return func(entities.Mutator) entities.Mutator {
		return entities.MutateFunc(func(context.Context, entities.Mutation) (entities.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []entities.Hook {
//		return []entities.Hook{
//			Reject(entities.Delete|entities.Update),
//		}
//	}
func Reject(op entities.Op) entities.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []entities.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...entities.Hook) Chain {
	return Chain{append([]entities.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() entities.Hook {
	return func(mutator entities.Mutator) entities.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...entities.Hook) Chain {
	newHooks := make([]entities.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
