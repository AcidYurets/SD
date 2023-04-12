package db

import (
	"calend/internal/pkg/search"
	"entgo.io/ent/dialect/sql"
	"strings"
	"time"
	"unicode"
)

type Predicate func(*sql.Selector)

type QueryBuilder struct {
	predicates []Predicate
}

func (b *QueryBuilder) AddField(name string, builder search.QueryFieldBuilder, optionalWrapper ...func(p Predicate) Predicate) {
	if builder == nil || !builder.IsValid() {
		return
	}

	wrapper := func(p Predicate) Predicate {
		return p
	}
	// Если optionalWrapper указан -> переопределяем
	if len(optionalWrapper) != 0 && optionalWrapper[0] != nil {
		wrapper = optionalWrapper[0]
	}

	builder.Build(name, b, wrapper)
}

func (b *QueryBuilder) Build() func(*sql.Selector) {
	if len(b.predicates) == 0 {
		return nil
	}

	// Соединяем все фильтры через И
	return and(b.predicates...)
}

func (b *QueryBuilder) add(pred Predicate) {
	b.predicates = append(b.predicates, pred)
}

func (b *QueryBuilder) Ts(field string, value string, wrapper func(p Predicate) Predicate) {
	fields := strings.Fields(field)
	if len(fields) == 0 {
		return
	}

	sValue := strings.ToLower(value)
	words := strings.FieldsFunc(sValue, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsNumber(r) })
	var queries []Predicate

	for _, field := range fields {
		var likes []Predicate

		// Проверяем все слова на наличие в заданном поле
		for _, word := range words {
			likes = append(likes, sql.FieldContains(field, word))
		}
		// Поле должно содержать все слова из запроса, поэтому соединяем через И
		queries = append(queries, and(likes...))
	}

	if len(queries) == 1 {
		b.add(wrapper(queries[0]))
	} else {
		// Любое поле может содержать все слова из запроса, поэтому соединяем через ИЛИ
		b.add(wrapper(or(queries...)))
	}
}

func (b *QueryBuilder) Eq(field string, value interface{}, wrapper func(p Predicate) Predicate) {
	b.add(wrapper(sql.FieldEQ(field, value)))
}

func (b *QueryBuilder) In(field string, values []interface{}, wrapper func(p Predicate) Predicate) {
	b.add(wrapper(sql.FieldIn(field, values...)))
}

func (b *QueryBuilder) Nin(field string, values []interface{}, wrapper func(p Predicate) Predicate) {
	b.add(wrapper(sql.FieldNotIn(field, values...)))
}

func (b *QueryBuilder) From(field string, value *time.Time, wrapper func(p Predicate) Predicate) {
	b.add(wrapper(sql.FieldGTE(field, value)))

}

func (b *QueryBuilder) To(field string, value *time.Time, wrapper func(p Predicate) Predicate) {
	b.add(wrapper(sql.FieldLTE(field, value)))
}

func (b *QueryBuilder) Range(field string, from, to *time.Time, wrapper func(p Predicate) Predicate) {
	b.add(wrapper(and(sql.FieldGTE(field, from), sql.FieldLTE(field, to))))
}

// ===================== Вспомогательные функции =======================

// and groups predicates with the AND operator between them.
func and(predicates ...Predicate) Predicate {
	return Predicate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// or groups predicates with the OR operator between them.
func or(predicates ...Predicate) Predicate {
	return Predicate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// not applies the not operator on the given predicate.
func not(p Predicate) Predicate {
	return Predicate(func(s *sql.Selector) {
		p(s.Not())
	})
}
