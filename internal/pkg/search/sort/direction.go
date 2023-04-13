package sort

import "calend/internal/pkg/search/engine/db/ent_types"

type Direction string

func (d *Direction) IsValid() bool {

	if d == nil {
		return false
	}

	switch *d {
	case DirectionDesc, DirectionAsc:
		return true
	default:
		return false
	}
}

func (d *Direction) Build(field string, b Builder, wrapper func(p ent_types.Predicate) ent_types.Predicate) {

	if !d.IsValid() {
		return
	}

	asc := true

	if *d == DirectionDesc {
		asc = false
	}

	b.AddSort(field, asc, wrapper)

}

const (
	DirectionAsc  Direction = "ASC"
	DirectionDesc Direction = "DESC"
)
