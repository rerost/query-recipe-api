package di

import (
	"context"

	"github.com/rerost/query-recipe-api/repo/snippet"
)

type ProvidorRepo func(context.Context, Config) snippet.Repo
