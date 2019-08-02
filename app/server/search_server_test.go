package server

import (
	"context"
	"testing"

	api_pb "github.com/rerost/query-recipe-api/api"
)

func Test_SearchServiceServer_ListSearches(t *testing.T) {
	svr := NewSearchServiceServer()

	ctx := context.Background()
	req := &api_pb.ListSearchesRequest{}

	resp, err := svr.ListSearches(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_SearchServiceServer_GetSearch(t *testing.T) {
	svr := NewSearchServiceServer()

	ctx := context.Background()
	req := &api_pb.GetSearchRequest{}

	resp, err := svr.GetSearch(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_SearchServiceServer_CreateSearch(t *testing.T) {
	svr := NewSearchServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateSearchRequest{}

	resp, err := svr.CreateSearch(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_SearchServiceServer_UpdateSearch(t *testing.T) {
	svr := NewSearchServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateSearchRequest{}

	resp, err := svr.UpdateSearch(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_SearchServiceServer_DeleteSearch(t *testing.T) {
	svr := NewSearchServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteSearchRequest{}

	resp, err := svr.DeleteSearch(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
