// +build wireinject

package di

import (
	"context"

	"github.com/google/go-github/v27/github"
	"github.com/google/wire"
	"github.com/rerost/query-recipe-api/infra/data"
	"github.com/rerost/query-recipe-api/infra/search"
	"github.com/rerost/query-recipe-api/repo/snippet"
	"golang.org/x/oauth2"
)

func NewRepo(context.Context, Config) snippet.Repo {
	wire.Build(
		oauth2.StaticTokenSource,
		oauth2.NewClient,
		ProvidorToken,
		ProvidorDataConfig,
		ProvidorSearchConfig,
		github.NewClient,
		search.NewGHClient,
		data.NewGHClient,
		snippet.NewRepo,
	)

	return nil
}
