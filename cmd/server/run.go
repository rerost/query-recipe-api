package main

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v27/github"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/rerost/query-recipe-api/app/server"
	"github.com/rerost/query-recipe-api/infra/data"
	"github.com/rerost/query-recipe-api/infra/search"
	"github.com/rerost/query-recipe-api/repo/snippet"
)

type Config struct {
	Owner             string // e.g. rerost
	Repository        string // e.g. query-recipe-api
	GithubAccessToken string
}

func run() error {
	// Application context
	ctx := context.Background()

	// TODO(@rerost) Use wire
	cfg := Config{
		Owner:             os.Getenv("GITHUB_OWNER"),
		Repository:        os.Getenv("GITHUB_REPOSITORY"),
		GithubAccessToken: os.Getenv("GITHUB_ACCESS_TOKEN"),
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	ghc := github.NewClient(tc)
	searchClient := search.NewGHClient(ghc, cfg.Owner, cfg.Repository)
	dataClient := data.NewGHClient(ghc, cfg.Owner, cfg.Repository)
	repo := snippet.NewRepo(searchClient, dataClient)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewSearchServiceServer(repo),
		),
		grapiserver.WithGatewayAddr("tcp", fmt.Sprintf(":%s", port)),
	)

	return s.ServeContext(ctx)
}
