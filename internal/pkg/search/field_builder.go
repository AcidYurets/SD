package search

import (
	"calend/internal/pkg/search/engine/db"
	"calend/internal/pkg/search/filter"
	"calend/internal/pkg/search/sort"
)

type QueryFieldBuilder interface {
	IsValid() bool
	Build(field string, b filter.Builder, wrapper func(p db.Predicate) db.Predicate)
}

type SortFieldBuilder interface {
	IsValid() bool
	Build(field string, b sort.Builder)
}
