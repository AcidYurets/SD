package filter

// OrQueryFilter фильтр по содержанию значения в любом из нескольких полей
type OrQueryFilter struct {
	Eq string
}

func (f *OrQueryFilter) Build(fields string, b Builder) {
	if !f.IsValid() {
		return
	}

	b.EqOr(fields, f.Eq)
}

func (f *OrQueryFilter) IsValid() bool {
	if f == nil {
		return false
	}

	return f != nil && len(f.Eq) > 0
}
