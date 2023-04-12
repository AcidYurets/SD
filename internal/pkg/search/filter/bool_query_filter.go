package filter

import (
	"calend/internal/pkg/search/engine/db/ent_types"
)

type BoolQueryFilter struct {
	// По полному совпадению значения
	Eq *bool
}

func (f *BoolQueryFilter) Build(field string, b Builder, wrapper func(p ent_types.Predicate) ent_types.Predicate) {

	if !f.IsValid() {
		return
	}

	if f.Eq != nil {
		b.Eq(field, f.Eq, wrapper)
	}
}

func (f *BoolQueryFilter) IsValid() bool {
	if f == nil {
		return false
	}

	return f.Eq != nil
}
