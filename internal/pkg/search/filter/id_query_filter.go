package filter

// IDQueryFilter - фильтр по id и uuid
//
//	uuid - требуется передать поле с типа keyword, например Uuid.keyword
type IDQueryFilter struct {
	// По полному совпадению значения
	Eq *string
	// По вхождению в массив значения
	In []string
	// По не вхождению в массив значения
	Nin []string
}

func (f *IDQueryFilter) Build(field string, b Builder) {

	if !f.IsValid() {
		return
	}

	if f.Eq != nil {
		b.Eq(field, *f.Eq)
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

func (f *IDQueryFilter) IsValid() bool {
	if f == nil {
		return false
	}

	return f.Eq != nil ||
		len(f.In) > 0 ||
		len(f.Nin) > 0
}
