package filter

import "calend/internal/pkg/search/engine/db/ent_types"

// FTSQueryFilter Полнотекстовый поиск
type FTSQueryFilter struct {
	Str string
}

func (f *FTSQueryFilter) Build(field string, b Builder, wrapper func(p ent_types.Predicate) ent_types.Predicate) {

	if !f.IsValid() {
		return
	}

	b.Ts(field, f.Str, wrapper)
}

func (f *FTSQueryFilter) IsValid() bool {
	if f == nil {
		return false
	}

	return f != nil && len(f.Str) > 0
}
