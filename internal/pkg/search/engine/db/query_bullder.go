package db

import (
	"calend/internal/pkg/search"
	"calend/internal/pkg/search/engine/db/types"
	"entgo.io/ent/dialect/sql"
	"strings"
	"time"
	"unicode"
)

type QueryBuilder struct {
	predicates []types.Predicate
	wrappers   map[string]types.Wrapper
}

func NewQueryBuilder(optionalWrappers ...map[string]types.Wrapper) *QueryBuilder {
	wrappers := make(map[string]types.Wrapper)

	// Если optionalWrappers указаны -> переопределяем
	if len(optionalWrappers) != 0 && optionalWrappers[0] != nil {
		wrappers = optionalWrappers[0]
	}

	builder := &QueryBuilder{
		wrappers: wrappers,
	}

	return builder
}

func (b *QueryBuilder) AddField(name string, builder search.QueryFieldBuilder) {
	if builder == nil || !builder.IsValid() {
		return
	}

	builder.Build(name, b)
}

func (b *QueryBuilder) Build() []types.Predicate {
	if len(b.predicates) == 0 {
		return nil
	}

	return b.predicates
}

func (b *QueryBuilder) add(pred types.Predicate) {
	b.predicates = append(b.predicates, pred)
}

func (b *QueryBuilder) Ts(field string, value string) {
	fields := strings.Fields(field)
	if len(fields) == 0 {
		return
	}

	sValue := strings.ToLower(value)
	words := strings.FieldsFunc(sValue, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsNumber(r) })
	var queries []types.Predicate

	for _, field := range fields {
		var likes []types.Predicate

		wrap, ok := b.wrappers[field]
		for _, word := range words {
			pred := sql.FieldContainsFold(lastElem(field), word)
			if ok {
				pred = wrap(pred)
			}

			likes = append(likes, pred)
		}
		// Поле должно содержать все слова из запроса, поэтому соединяем через И
		queries = append(queries, and(likes...))
	}

	if len(queries) == 1 {
		b.add(queries[0])
	} else {
		// Любое поле может содержать все слова из запроса, поэтому соединяем через ИЛИ
		b.add(or(queries...))
	}
}

func (b *QueryBuilder) Eq(field string, value interface{}) {
	pred := sql.FieldEQ(lastElem(field), value)
	wrap, ok := b.wrappers[field]
	if ok {
		pred = wrap(pred)
	}

	b.add(pred)
}

func (b *QueryBuilder) In(field string, values []interface{}) {
	pred := sql.FieldIn(lastElem(field), values...)
	wrap, ok := b.wrappers[field]
	if ok {
		pred = wrap(pred)
	}

	b.add(pred)
}

func (b *QueryBuilder) Nin(field string, values []interface{}) {
	pred := sql.FieldNotIn(lastElem(field), values...)
	wrap, ok := b.wrappers[field]
	if ok {
		pred = wrap(pred)
	}

	b.add(pred)
}

func (b *QueryBuilder) From(field string, value *time.Time) {
	pred := sql.FieldGTE(lastElem(field), value)
	wrap, ok := b.wrappers[field]
	if ok {
		pred = wrap(pred)
	}

	b.add(pred)
}

func (b *QueryBuilder) To(field string, value *time.Time) {
	pred := sql.FieldLTE(lastElem(field), value)
	wrap, ok := b.wrappers[field]
	if ok {
		pred = wrap(pred)
	}

	b.add(pred)
}

func (b *QueryBuilder) Range(field string, from, to *time.Time) {
	pred := and(sql.FieldGTE(lastElem(field), from), sql.FieldLTE(lastElem(field), to))
	wrap, ok := b.wrappers[field]
	if ok {
		pred = wrap(pred)
	}

	b.add(pred)
}

// ===================== Вспомогательные функции =======================

// and groups predicates with the AND operator between them.
func and(predicates ...types.Predicate) types.Predicate {
	return types.Predicate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// or groups predicates with the OR operator between them.
func or(predicates ...types.Predicate) types.Predicate {
	return types.Predicate(func(s *sql.Selector) {
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
func not(p types.Predicate) types.Predicate {
	return types.Predicate(func(s *sql.Selector) {
		p(s.Not())
	})
}

// lastElem получает последний элемент в цепочке (например, из aa.bb.cc получит cc)
func lastElem(s string) string {
	slice := strings.Split(s, ".")
	return slice[len(slice)-1]
}
