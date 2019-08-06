package server

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rerost/query-recipe-api/app/di"

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
func NewSearchServiceServer(r di.ProvidorRepo) SearchServiceServer {
	return &searchServiceServerImpl{
		r: r,
	}
}

type searchServiceServerImpl struct {
	r di.ProvidorRepo
}

func (s *searchServiceServerImpl) Search(ctx context.Context, req *api_pb.SearchRequest) (*api_pb.SearchResult, error) {
	m := req.GetMetadata()
	if m == nil {
		return nil, status.Error(codes.InvalidArgument, "Need metadata")
	}
	repo := s.r(ctx, di.ProvidorConfig(m.AccessToken, m.Owner, m.Repository))
	query := req.GetKeyword()
	snippets, err := repo.Search(ctx, query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return nil, status.Error(codes.Internal, "Failed to search snipppets")
	}

	result := api_pb.SearchResult{}
	result.Hits = make([]*type_pb.Snippet, len(snippets))

	for i, s := range snippets {
		result.Hits[i] = &type_pb.Snippet{
			Id:       string(s.ID),
			Sql:      s.SQL,
			Document: s.Document,
		}
	}

	return &result, nil
}
