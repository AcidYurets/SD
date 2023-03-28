package filter

type BoolQueryFilter struct {
	// По полному совпадению значения
	Eq *bool
	// По полному совпадению значения
	In []bool
	// По полному совпадению значения
	Nin []bool
}

func (f *BoolQueryFilter) Build(field string, b Builder) {

	if !f.IsValid() {
		return
	}

	if f.Eq != nil {
		b.Eq(field, f.Eq)
	}

	if len(f.In) > 0 {
		var values []interface{}
		for _, val := range f.In {
			values = append(values, val)
		}
		b.In(field, values)
	}
	if len(f.Nin) > 0 {
		var values []interface{}
		for _, val := range f.Nin {
			values = append(values, val)
		}
		b.Nin(field, values)
	}
}

func (f *BoolQueryFilter) IsValid() bool {
	if f == nil {
		return false
	}

	return f.Eq != nil ||
		len(f.In) > 0 ||
		len(f.Nin) > 0
}
