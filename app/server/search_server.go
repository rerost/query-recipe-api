package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rerost/query-recipe-api/repo/snippet"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	api_pb "github.com/rerost/query-recipe-api/api"
	type_pb "github.com/rerost/query-recipe-api/api/type"
)

// SearchServiceServer is a composite interface of api_pb.SearchServiceServer and grapiserver.Server.
type SearchServiceServer interface {
	api_pb.SearchServiceServer
	grapiserver.Server
}

// NewSearchServiceServer creates a new SearchServiceServer instance.
func NewSearchServiceServer(sr snippet.Repo) SearchServiceServer {
	return &searchServiceServerImpl{
		sr: sr,
	}
}

type searchServiceServerImpl struct {
	sr snippet.Repo
}

func (s *searchServiceServerImpl) Search(ctx context.Context, req *api_pb.SearchRequest) (*api_pb.SearchResult, error) {
	query := req.GetKeyword()
	snippets, err := s.sr.Search(ctx, query)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to search snipppets")
	}

	result := api_pb.SearchResult{}
	result.Hits = make([]*type_pb.Snippet, len(snippets))

	for i, s := range snippets {
		result.Hits[i] = &type_pb.Snippet{
			Id:  string(s.ID),
			Sql: s.SQL,
		}
	}

	return &result, nil
}
