package db

import (
	"calend/internal/pkg/search"
	"entgo.io/ent/dialect/sql"
)

type SortBuilder struct {
}

func (b *SortBuilder) Build() func(options *sql.OrderTermOptions) {
	return func(options *sql.OrderTermOptions) {

	}
}

func (b *SortBuilder) AddSort(field string, asc bool) {

}

func (b *SortBuilder) AddField(name string, builder search.SortFieldBuilder) {
	if builder == nil {
		return
	}

	if !builder.IsValid() {
		return
	}

	builder.Build(name, b)
}
