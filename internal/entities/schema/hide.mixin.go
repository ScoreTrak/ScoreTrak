package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type HideMixin struct {
	mixin.Schema
}

func (HideMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("hidden").Optional().Default(false),
	}
}

//func FilterHiddenRule() privacy.QueryMutationRule {
//	type HiddenFilter interface {
//		WhereHidden(p entql.BoolP)
//	}
//	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
//		hf, ok := f.(HiddenFilter)
//		if !ok {
//			return privacy.Denyf("unexpected filter type")
//		}
//		hf.WhereHidden(entql.BoolEQ(false))
//		return privacy.Skip
//	})
//}
