package elastic

import "github.com/olivere/elastic/v7"

type PaginateBuilder struct {
	page        *int
	pageSize    *int
	searchAfter interface{}
}

func (b *PaginateBuilder) Build() func(searchService *elastic.SearchService) *elastic.SearchService {
	return func(searchService *elastic.SearchService) *elastic.SearchService {

		// По умолчанию пустая выборка
		searchService.Size(0)
		if b.pageSize != nil && *b.pageSize > 0 {
			searchService.Size(*b.pageSize)
		}
		switch {
		case b.searchAfter != nil:
			searchService.SearchAfter(b.searchAfter)
		case b.page != nil && b.pageSize != nil && *b.page > 0 && *b.pageSize > 0:
			searchService.From(*b.pageSize * (*b.page - 1))
		}

		return searchService
	}
}

func (b *PaginateBuilder) Paginate(page *int, pageSize *int, searchAfter interface{}) {
	b.page = page
	b.pageSize = pageSize
	b.searchAfter = searchAfter
}
