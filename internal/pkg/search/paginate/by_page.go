package paginate

import "calend/internal/pkg/search/engine/db"

type ByPage struct {
	PageSize *int // Размер страницы выборки
	Page     *int // Номер страницы выборки
}

func BuildPaginate(paginate *ByPage) (int, int) {
	if paginate == nil {
		return 0, 0
	}
	builder := &db.PaginateBuilder{}

	builder.Paginate(paginate.Page, paginate.PageSize)

	return builder.Build()
}
