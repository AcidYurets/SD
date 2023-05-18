package elastic

import (
	"calend/internal/pkg/search"
	"github.com/olivere/elastic/v7"
)

// SortBuilder облегчает сборку фильтров для БД
type SortBuilder struct {
	sorters []elastic.Sorter
}

func (b *SortBuilder) Build() []elastic.Sorter {
	return b.sorters
}

func (b *SortBuilder) AddSort(field string, asc bool) {
	keywordField := trimKeyword(field) + ".keyword"

	fieldSort := elastic.NewFieldSort(keywordField)
	fieldSort.Order(asc)

	b.sorters = append(b.sorters, fieldSort)
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
