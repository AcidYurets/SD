package filter

import (
	"time"
)

type Builder interface {
	Ts(field string, value string)           // Полнотекстовый поиск по полю
	Eq(field string, value interface{})      // По совпадению значений
	In(field string, values []interface{})   // По вхождению в массив значений
	Nin(field string, values []interface{})  // По не вхождению в массив значений
	From(field string, value *time.Time)     // По дате от
	To(field string, value *time.Time)       // По дате до
	Range(field string, from, to *time.Time) // Между
}
