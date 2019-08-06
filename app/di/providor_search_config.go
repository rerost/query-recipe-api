package di

import "github.com/rerost/query-recipe-api/infra/search"

func ProvidorSearchConfig(cfg Config) search.Config {
	return cfg
}
