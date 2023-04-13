package sort

import "calend/internal/pkg/search/engine/db/ent_types"

type Builder interface {
	AddSort(field string, asc bool, wrapper func(p ent_types.Predicate) ent_types.Predicate)
}
