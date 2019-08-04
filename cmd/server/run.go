package main

import (
	"context"
	"net/http"

	"github.com/google/go-github/v27/github"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/rerost/query-recipe-api/app/server"
	"github.com/rerost/query-recipe-api/infra/data"
	"github.com/rerost/query-recipe-api/infra/search"
	"github.com/rerost/query-recipe-api/repo/snippet"
)

type Config struct {
	Owner      string // e.g. rerost
	Repository string // e.g. query-recipe-api
}

func run() error {
	// Application context
	ctx := context.Background()

	// TODO(@rerost) Use wire
	cfg := Config{
		Owner:      "rerost",
		Repository: "test-query-recipe",
	}
	httpClient := new(http.Client)
	ghc := github.NewClient(httpClient)
	searchClient := search.NewGHClient(ghc, cfg.Owner, cfg.Repository)
	dataClient := data.NewGHClient(ghc, cfg.Owner, cfg.Repository)
	repo := snippet.NewRepo(searchClient, dataClient)

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewSearchServiceServer(repo),
		),
	)
	return s.ServeContext(ctx)
}
