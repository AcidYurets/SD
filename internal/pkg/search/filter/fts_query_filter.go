package filter

// FTSQueryFilter Полнотекстовый поиск
type FTSQueryFilter struct {
	Str string
}

func (f *FTSQueryFilter) Build(field string, b Builder) {

	if !f.IsValid() {
		return
	}

	b.Ts(field, f.Str)
}

func (f *FTSQueryFilter) IsValid() bool {
	if f == nil {
		return false
	}

	return f != nil && len(f.Str) > 0
}
