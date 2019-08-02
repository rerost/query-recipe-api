package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/rerost/query-recipe-api/api"
)

// SnippetServiceServer is a composite interface of api_pb.SnippetServiceServer and grapiserver.Server.
type SnippetServiceServer interface {
	api_pb.SnippetServiceServer
	grapiserver.Server
}

// NewSnippetServiceServer creates a new SnippetServiceServer instance.
func NewSnippetServiceServer() SnippetServiceServer {
	return &snippetServiceServerImpl{}
}

type snippetServiceServerImpl struct {
}

func (s *snippetServiceServerImpl) ListSnippets(ctx context.Context, req *api_pb.ListSnippetsRequest) (*api_pb.ListSnippetsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *snippetServiceServerImpl) GetSnippet(ctx context.Context, req *api_pb.GetSnippetRequest) (*api_pb.Snippet, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *snippetServiceServerImpl) CreateSnippet(ctx context.Context, req *api_pb.CreateSnippetRequest) (*api_pb.Snippet, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *snippetServiceServerImpl) UpdateSnippet(ctx context.Context, req *api_pb.UpdateSnippetRequest) (*api_pb.Snippet, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *snippetServiceServerImpl) DeleteSnippet(ctx context.Context, req *api_pb.DeleteSnippetRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
