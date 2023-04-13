package filter

type BoolQueryFilter struct {
	// По полному совпадению значения
	Eq *bool
}

func (f *BoolQueryFilter) Build(field string, b Builder) {

	if !f.IsValid() {
		return
	}

	if f.Eq != nil {
		b.Eq(field, f.Eq)
	}
}

func (f *BoolQueryFilter) IsValid() bool {
	if f == nil {
		return false
	}

	return f.Eq != nil
}
