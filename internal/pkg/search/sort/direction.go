package sort

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

func (d *Direction) Build(field string, b Builder) {

	if !d.IsValid() {
		return
	}

	asc := true

	if *d == DirectionDesc {
		asc = false
	}

	b.AddSort(field, asc)

}

const (
	DirectionAsc  Direction = "ASC"
	DirectionDesc Direction = "DESC"
)
