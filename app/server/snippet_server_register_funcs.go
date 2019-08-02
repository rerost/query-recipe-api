// Code generated by github.com/izumin5210/grapi. DO NOT EDIT.

package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api_pb "github.com/rerost/query-recipe-api/api"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *snippetServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	api_pb.RegisterSnippetServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *snippetServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api_pb.RegisterSnippetServiceHandler(ctx, mux, conn)
}