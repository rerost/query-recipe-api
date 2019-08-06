package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gorilla/handlers"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/rerost/query-recipe-api/app/di"
	"github.com/rerost/query-recipe-api/app/server"
)

type Config struct {
	Owner             string // e.g. rerost
	Repository        string // e.g. query-recipe-api
	GithubAccessToken string
}

func run() error {
	// Application context
	ctx := context.Background()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewSearchServiceServer(di.NewRepo),
		),
		grapiserver.WithGatewayAddr("tcp", fmt.Sprintf(":%s", port)),
		grapiserver.WithGatewayServerMiddlewares(
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	)

	return s.ServeContext(ctx)
}
