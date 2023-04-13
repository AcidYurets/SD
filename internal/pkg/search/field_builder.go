package search

import (
	"calend/internal/pkg/search/engine/db/ent_types"
	"calend/internal/pkg/search/filter"
	"calend/internal/pkg/search/sort"
)

type QueryFieldBuilder interface {
	IsValid() bool
	Build(field string, b filter.Builder, wrapper func(p ent_types.Predicate) ent_types.Predicate)
}

type SortFieldBuilder interface {
	IsValid() bool
	Build(field string, b sort.Builder, wrapper func(p ent_types.Predicate) ent_types.Predicate)
}
