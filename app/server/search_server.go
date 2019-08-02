package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/rerost/query-recipe-api/api"
)

// SearchServiceServer is a composite interface of api_pb.SearchServiceServer and grapiserver.Server.
type SearchServiceServer interface {
	api_pb.SearchServiceServer
	grapiserver.Server
}

// NewSearchServiceServer creates a new SearchServiceServer instance.
func NewSearchServiceServer() SearchServiceServer {
	return &searchServiceServerImpl{}
}

type searchServiceServerImpl struct {
}

func (s *searchServiceServerImpl) ListSearches(ctx context.Context, req *api_pb.ListSearchesRequest) (*api_pb.ListSearchesResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *searchServiceServerImpl) GetSearch(ctx context.Context, req *api_pb.GetSearchRequest) (*api_pb.Search, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *searchServiceServerImpl) CreateSearch(ctx context.Context, req *api_pb.CreateSearchRequest) (*api_pb.Search, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *searchServiceServerImpl) UpdateSearch(ctx context.Context, req *api_pb.UpdateSearchRequest) (*api_pb.Search, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *searchServiceServerImpl) DeleteSearch(ctx context.Context, req *api_pb.DeleteSearchRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
