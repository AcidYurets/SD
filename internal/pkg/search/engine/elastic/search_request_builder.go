package elastic

import (
	"calend/internal/pkg/search/paginate"
	"github.com/olivere/elastic/v7"
)

type SearchRequestBuilder struct {
	Query    *elastic.BoolQuery
	Filter   func() *elastic.BoolQuery
	Sort     func() func(search *elastic.SearchService) *elastic.SearchService
	Paginate *paginate.ByPage
}

func (b *SearchRequestBuilder) Build() (
	sort func(search *elastic.SearchService) *elastic.SearchService,
	paginate func(search *elastic.SearchService) *elastic.SearchService,
) {

	emptyScope := func(search *elastic.SearchService) *elastic.SearchService {
		return search
	}
	sort, paginate = emptyScope, emptyScope
	if b == nil {
		return
	}

	if b.Filter != nil {
		if val := b.Filter(); val != nil {
			if b.Query != nil {
				b.Query.Filter(val)
			} else {
				b.Query = val
			}

		}
	}

	if b.Sort != nil {
		if val := b.Sort(); val != nil {
			sort = val
		}
	}

	if b.Paginate != nil {
		builder := &PaginateBuilder{}

		builder.Paginate(b.Paginate.Page, b.Paginate.PageSize, nil)

		if val := builder.Build(); val != nil {
			paginate = val
		}
	}

	return
}
