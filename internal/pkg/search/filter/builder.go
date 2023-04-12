package filter

import (
	"calend/internal/pkg/search/engine/db"
	"time"
)

type Builder interface {
	Ts(field string, value string, wrapper func(p db.Predicate) db.Predicate)           // Полнотекстовый поиск по полю
	Eq(field string, value interface{}, wrapper func(p db.Predicate) db.Predicate)      // По совпадению значений
	In(field string, values []interface{}, wrapper func(p db.Predicate) db.Predicate)   // По вхождению в массив значений
	Nin(field string, values []interface{}, wrapper func(p db.Predicate) db.Predicate)  // По не вхождению в массив значений
	From(field string, value *time.Time, wrapper func(p db.Predicate) db.Predicate)     // По дате от
	To(field string, value *time.Time, wrapper func(p db.Predicate) db.Predicate)       // По дате до
	Range(field string, from, to *time.Time, wrapper func(p db.Predicate) db.Predicate) // Между
}
