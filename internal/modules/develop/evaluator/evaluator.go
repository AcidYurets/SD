package evaluator

import (
	"calend/internal/modules/db/ent"
	event_elastic "calend/internal/modules/domain/event/elastic"
	"calend/internal/modules/domain/search/dto"
	"calend/internal/modules/elastic/index"
	"calend/internal/pkg/search/paginate"
	"calend/internal/utils/ptr"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

type Evaluator struct {
	dbClient     *ent.Client
	elasticIndex index.IElasticIndex
}

func NewEvaluator(dbClient *ent.Client, elasticClient *elastic.Client) *Evaluator {
	return &Evaluator{
		dbClient: dbClient,
		elasticIndex: index.NewElasticIndex(index.Config{
			Client:  elasticClient,
			Index:   event_elastic.EventIndexName,
			IdField: "Uuid",
		}),
	}
}

func (r *Evaluator) EvaluateSearchEvents(ctx context.Context, evaluationRequest *EvaluationRequest) (*EvaluationResults, error) {
	res := &EvaluationResults{}

	//sortDirection := sort.DirectionAsc
	searchRequest := &dto.EventSearchRequest{
		//Filter: &dto.EventFilter{
		//	FTSearchStr: &filter.FTSQueryFilter{
		//		Str: "Соб",
		//	},
		//	CreatorLogin: &filter.TextQueryFilter{
		//		Ts: ptr.String("Us"),
		//	},
		//	Description: &filter.TextQueryFilter{
		//		Ts: ptr.String("Опи"),
		//	},
		//	TagName: &filter.TextQueryFilter{
		//		Ts: ptr.String("Тег"),
		//	},
		//	InvitedUserUuid: &filter.IDQueryFilter{
		//		Nin: []string{"3075060f-9211-40c9-928e-404b2ff866f2"},
		//	},
		//},
		//Sort: &dto.EventSort{
		//	CreatorLogin: &sortDirection,
		//},
		Paginate: &paginate.ByPage{
			Page:     ptr.Int(1),
			PageSize: ptr.Int(evaluationRequest.PageSize),
		},
	}

	var start, stop time.Time
	var err error

	start = time.Now()
	_, err = r.searchEventsDB(ctx, searchRequest)
	if err != nil {
		return nil, err
	}
	stop = time.Now()
	res.DurationDB = stop.Sub(start)

	start = time.Now()
	_, err = r.searchEventsElastic(ctx, searchRequest)
	if err != nil {
		return nil, err
	}
	stop = time.Now()
	res.DurationElastic = stop.Sub(start)

	return res, nil
}

type EvaluationRequest struct {
	PageSize int
}

type EvaluationResults struct {
	DurationDB      time.Duration
	DurationElastic time.Duration
}

func (r *EvaluationResults) String() string {
	return fmt.Sprintf("DB: %s, Elastic: %s", r.DurationDB, r.DurationElastic)
}
