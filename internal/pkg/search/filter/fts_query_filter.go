package filter

import "calend/internal/pkg/search/engine/db"

// FTSQueryFilter Полнотекстовый поиск
type FTSQueryFilter struct {
	Str string
}

func (f *FTSQueryFilter) Build(field string, b Builder, wrapper func(p db.Predicate) db.Predicate) {

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
