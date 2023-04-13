package filter

import (
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

func (f *DateQueryFilter) Build(field string, b Builder) {

	if !f.IsValid() {
		return
	}

	if f.Eq != nil {
		b.Eq(field, f.Eq)
	}

	if f.isValidRangeQuery() {

		switch {
		case f.From != nil && f.To != nil:
			b.Range(field, f.From, f.To)
		case f.From != nil:
			b.From(field, f.From)
		case f.To != nil:
			b.To(field, f.To)
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
