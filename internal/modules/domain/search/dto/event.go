package dto

import (
	"calend/internal/pkg/search/filter"
	"calend/internal/pkg/search/paginate"
	"calend/internal/pkg/search/sort"
)

type EventSearchRequest struct {
	Filter   *EventFilter
	Sort     *EventSort
	Paginate *paginate.ByPage
}

type EventFilter struct {
	FTSearchStr *filter.FTSQueryFilter // Полнотекстовый поиск по всем полям

	Timestamp   *filter.DateQueryFilter // Поиск по временной метке
	Name        *filter.TextQueryFilter // Поиск по названию
	Description *filter.TextQueryFilter // Поиск по описанию
	Type        *filter.TextQueryFilter // Поиск по типу
	IsWholeDay  *filter.BoolQueryFilter // Поиск по признаку полного дня
	CreatorUuid *filter.IDQueryFilter   // Поиск по Uuid создателя

	CreatorLogin *filter.TextQueryFilter // Поиск по логину создателя
	TagName      *filter.TextQueryFilter // Поиск по названию тегов
}

type EventSort struct {
	Timestamp    *sort.Direction // По временной метке
	Name         *sort.Direction // По названию
	Description  *sort.Direction // По описанию
	Type         *sort.Direction // По типу
	CreatorLogin *sort.Direction // По имени создателя
}
