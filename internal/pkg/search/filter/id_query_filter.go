package filter

import (
	"calend/internal/pkg/search/engine/db/ent_types"
)

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
	// Учитывать ли дочерние элементы
	WithHierarchy *bool
}

func (f *IDQueryFilter) Build(field string, b Builder, wrapper func(p ent_types.Predicate) ent_types.Predicate) {

	if !f.IsValid() {
		return
	}

	if f.Eq != nil {
		b.Eq(field, *f.Eq, wrapper)
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
		for _, val := range f.Nin {
			values = append(values, val)
		}
		b.Nin(field, values, wrapper)
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
