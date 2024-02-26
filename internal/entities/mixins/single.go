package mixins

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/mixin"
	"errors"
)

const FIRST_ID = "00000000000000000000000000"

type SingleMixin struct {
	mixin.Schema
}

func (SingleMixin) Hooks() []ent.Hook {
	type IDSetter interface {
		SetID(string)
	}
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				entity, ok := m.(IDSetter)
				if !ok {
					return nil, errors.New("failed")
				}
				entity.SetID(FIRST_ID)
				return next.Mutate(ctx, m)
			})
		},
	}
}
