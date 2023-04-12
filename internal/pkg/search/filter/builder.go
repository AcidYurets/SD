package filter

import (
	"calend/internal/pkg/search/engine/db/ent_types"
	"time"
)

type Builder interface {
	Ts(field string, value string, wrapper func(p ent_types.Predicate) ent_types.Predicate)           // Полнотекстовый поиск по полю
	Eq(field string, value interface{}, wrapper func(p ent_types.Predicate) ent_types.Predicate)      // По совпадению значений
	In(field string, values []interface{}, wrapper func(p ent_types.Predicate) ent_types.Predicate)   // По вхождению в массив значений
	Nin(field string, values []interface{}, wrapper func(p ent_types.Predicate) ent_types.Predicate)  // По не вхождению в массив значений
	From(field string, value *time.Time, wrapper func(p ent_types.Predicate) ent_types.Predicate)     // По дате от
	To(field string, value *time.Time, wrapper func(p ent_types.Predicate) ent_types.Predicate)       // По дате до
	Range(field string, from, to *time.Time, wrapper func(p ent_types.Predicate) ent_types.Predicate) // Между
}
