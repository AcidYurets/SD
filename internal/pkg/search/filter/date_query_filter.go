package filter

import (
	"calend/internal/pkg/search/engine/db"
	"time"
)

type DateQueryFilter struct {
	// По полному совпадению значения
	Eq *time.Time
	// Начало периода
	From *time.Time
	// Окончание периода
	To *time.Time
	// Зона времени
	TimeZone *string
	// Включает нижнюю границу
	IncludeLower *bool
	// Включает верхнюю границу
	IncludeUpper *bool
}

func (f *DateQueryFilter) Build(field string, b Builder, wrapper func(p db.Predicate) db.Predicate) {

	if !f.IsValid() {
		return
	}

	if f.Eq != nil {
		b.Eq(field, f.Eq, wrapper)
	}

	if f.isValidRangeQuery() {

		switch {
		case f.From != nil && f.To != nil:
			b.Range(field, f.From, f.To, wrapper)
		case f.From != nil:
			b.From(field, f.From, wrapper)
		case f.To != nil:
			b.To(field, f.To, wrapper)
		}
	}

}

func (f *DateQueryFilter) IsValid() bool {
	if f == nil {
		return false
	}

	return f.Eq != nil ||
		f.isValidRangeQuery()
}

func (f *DateQueryFilter) isValidRangeQuery() bool {
	return f.From != nil ||
		f.To != nil
}
