package db

import (
	"calend/internal/pkg/search"
	"calend/internal/pkg/search/engine/db/ent_types"
	"entgo.io/ent/dialect/sql"
)

type SortBuilder struct {
	predicates []ent_types.Predicate
}

func (b *SortBuilder) add(pred ent_types.Predicate) {
	b.predicates = append(b.predicates, pred)
}

func (b *SortBuilder) Build() []ent_types.Predicate {
	if len(b.predicates) == 0 {
		return nil
	}

	return b.predicates
}

func (b *SortBuilder) AddSort(field string, asc bool, wrapper func(p ent_types.Predicate) ent_types.Predicate) {
	opt := sql.OrderAsc()
	if !asc {
		opt = sql.OrderDesc()
	}

	b.add(wrapper(sql.OrderByField(field, opt).ToFunc()))
}

func (b *SortBuilder) AddField(name string, builder search.SortFieldBuilder, optionalWrapper ...func(p ent_types.Predicate) ent_types.Predicate) {
	if builder == nil || !builder.IsValid() {
		return
	}

	wrapper := func(p ent_types.Predicate) ent_types.Predicate {
		return p
	}
	// Если optionalWrapper указан -> переопределяем
	if len(optionalWrapper) != 0 && optionalWrapper[0] != nil {
		wrapper = optionalWrapper[0]
	}

	builder.Build(name, b, wrapper)
}
