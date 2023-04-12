package db

type PaginateBuilder struct {
	page     *int
	pageSize *int
}

func (b *PaginateBuilder) Paginate(page *int, pageSize *int) {
	b.page = page
	b.pageSize = pageSize
}

func (b *PaginateBuilder) Build() (int, int) {
	if *b.page > 0 && *b.pageSize > 0 {
		return *b.pageSize, *b.pageSize * (*b.page - 1)
	}

	return 0, 0 // значения, не оказывающие эффекта в ent
}
