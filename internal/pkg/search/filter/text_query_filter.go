package filter

import "calend/internal/pkg/search/engine/db"

type TextQueryFilter struct {
	// Полнотекстовый поиск
	Ts *string
	// По полному совпадению значения
	Eq *string
	// По вхождению в массив значения
	In []string
	// По не вхождению в массив значения
	Nin []string
}

func (f *TextQueryFilter) Build(field string, b Builder, wrapper func(p db.Predicate) db.Predicate) {

	if !f.IsValid() {
		return
	}

	if f.Eq != nil {
		b.Eq(field, *f.Eq, wrapper)
	}

	if f.Ts != nil {
		b.Ts(field, *f.Ts, wrapper)
	}

	if len(f.In) > 0 {

		var values []interface{}
		for _, val := range f.In {
			values = append(values, val)
		}

		b.In(field, values, wrapper)
	}

	if len(f.Nin) > 0 {
		var values []interface{}
		for _, val := range f.In {
			values = append(values, val)
		}

		b.Nin(field, values, wrapper)

	}

}

func (f *TextQueryFilter) IsValid() bool {
	if f == nil {
		return false
	}

	return f.Eq != nil ||
		f.Ts != nil ||
		len(f.In) > 0 ||
		len(f.Nin) > 0
}
