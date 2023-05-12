package search

import (
	"calend/internal/modules/config"
	"calend/internal/modules/domain/search/repo_db"
	"calend/internal/modules/domain/search/repo_elastic"
	"calend/internal/modules/domain/search/service"
	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		service.Module,
		repo_db.Module,
		repo_elastic.Module,

		fx.Provide(ProvideSearchRepo),
	)

	Invokables = fx.Options(
		service.Invokables,
		repo_db.Invokables,
		repo_elastic.Invokables,
	)
)

func ProvideSearchRepo(cfg config.Config, dbRepo *repo_db.SearchRepo, elasticRepo *repo_elastic.SearchRepo) service.ISearchRepo {
	if cfg.SearchService == "elastic" {
		return elasticRepo
	} else {
		return dbRepo
	}

}
