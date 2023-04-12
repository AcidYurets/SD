package filter

import "calend/internal/pkg/search/engine/db"

type BoolQueryFilter struct {
	// По полному совпадению значения
	Eq *bool
}

func (f *BoolQueryFilter) Build(field string, b Builder, wrapper func(p db.Predicate) db.Predicate) {

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
