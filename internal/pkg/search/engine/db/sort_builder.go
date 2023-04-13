package db

import (
	"calend/internal/pkg/search"
	"calend/internal/pkg/search/engine/db/types"
	"entgo.io/ent/dialect/sql"
)

type SortBuilder struct {
	predicates []types.Predicate
	wrappers   map[string]types.Wrapper
}

func NewSortBuilder(optionalWrappers ...map[string]types.Wrapper) *SortBuilder {
	wrappers := make(map[string]types.Wrapper)

	// Если optionalWrappers указаны -> переопределяем
	if len(optionalWrappers) != 0 && optionalWrappers[0] != nil {
		wrappers = optionalWrappers[0]
	}

	builder := &SortBuilder{
		wrappers: wrappers,
	}

	return builder
}

func (b *SortBuilder) AddField(name string, builder search.SortFieldBuilder) {
	if builder == nil || !builder.IsValid() {
		return
	}

	builder.Build(name, b)
}

func (b *SortBuilder) Build() []types.Predicate {
	if len(b.predicates) == 0 {
		return nil
	}

	return b.predicates
}

func (b *SortBuilder) add(pred types.Predicate) {
	b.predicates = append(b.predicates, pred)
}

func (b *SortBuilder) AddSort(field string, asc bool) {
	opt := sql.OrderAsc()
	if !asc {
		opt = sql.OrderDesc()
	}

	pred := sql.OrderByField(lastElem(field), opt).ToFunc()
	wrap, ok := b.wrappers[field]
	if ok {
		pred = wrap(pred)
	}

	b.add(pred)
}
