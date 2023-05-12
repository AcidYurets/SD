package elastic

import (
	"calend/internal/pkg/search"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

type QueryBuilder struct {
	filters []elastic.Query
}

func (b *QueryBuilder) AddField(name string, builder search.QueryFieldBuilder) {

	if builder == nil || !builder.IsValid() {
		return
	}

	builder.Build(name, b)
}

func (b *QueryBuilder) add(filter elastic.Query) {
	b.filters = append(b.filters, filter)
}

func (b *QueryBuilder) Build() *elastic.BoolQuery {
	if len(b.filters) == 0 {
		return nil
	}
	return elastic.NewBoolQuery().Filter(b.filters...)
}

func (b *QueryBuilder) Eq(field string, value interface{}) {
	b.add(elastic.NewTermQuery(field, value))
}

func (b *QueryBuilder) In(field string, values []interface{}) {
	if len(values) == 0 {
		return
	}

	b.add(elastic.NewTermsQuery(field, values...))
}

func (b *QueryBuilder) Nin(field string, values []interface{}) {
	if len(values) == 0 {
		return
	}

	b.add(elastic.NewBoolQuery().MustNot(
		elastic.NewTermsQuery(field, values...),
	))
}

func (b *QueryBuilder) Ts(field string, value string) {
	fields := strings.Fields(field)

	if len(fields) == 0 {
		return
	}
	sValue := strings.ToLower(value)
	words := strings.Fields(sValue)
	queries := []elastic.Query{
		elastic.NewMultiMatchQuery(sValue, fields...).Type("phrase"),
	}
	for _, field := range fields {
		queries = append(queries, []elastic.Query{
			elastic.NewWildcardQuery(field, "*"+sValue+"*"),
			elastic.NewQueryStringQuery("*" + strings.Join(words, "* ") + "*").Field(field).DefaultOperator("AND"),
		}...)
	}

	b.add(elastic.NewBoolQuery().Should(
		queries...,
	))
}

func (b *QueryBuilder) Range(field string, from, to *time.Time) {
	rangeQuery := elastic.NewRangeQuery(field)
	rangeQuery.From(from).To(to)

	b.add(rangeQuery)
}

func (b *QueryBuilder) From(field string, value *time.Time) {
	rangeQuery := elastic.NewRangeQuery(field)
	rangeQuery.From(value)

	b.add(rangeQuery)
}

func (b *QueryBuilder) To(field string, value *time.Time) {
	rangeQuery := elastic.NewRangeQuery(field)
	rangeQuery.To(value)

	b.add(rangeQuery)
}
