package main

import (
	"fmt"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		//modules.AppModule,
		//modules.AppInvokables,
		fx.Provide(
			ProvideSearchRepo,
		),
		fx.Invoke(
			InvokeService,
		),
	).Run()
}

type IRepo interface {
	SearchEvents() string
}

type DBRepo struct{}

func (r *DBRepo) SearchEvents() string { return "db" }

type ElasticRepo struct{}

func (r *ElasticRepo) SearchEvents() string { return "elastic" }

func NewDBRepo() *DBRepo {
	return &DBRepo{}
}

func NewElasticRepo() *ElasticRepo {
	return &ElasticRepo{}
}

func InvokeService(repo IRepo) {
	fmt.Println(repo.SearchEvents())
}

func ProvideSearchRepo( /*config*/ ) IRepo {
	var choice int
	_, _ = fmt.Scanln(&choice)

	if choice == 0 {
		return NewDBRepo()
	} else {
		return NewElasticRepo()
	}
}
