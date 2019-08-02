package server

import (
	"context"
	"testing"

	api_pb "github.com/rerost/query-recipe-api/api"
)

func Test_SnippetServiceServer_ListSnippets(t *testing.T) {
	svr := NewSnippetServiceServer()

	ctx := context.Background()
	req := &api_pb.ListSnippetsRequest{}

	resp, err := svr.ListSnippets(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_SnippetServiceServer_GetSnippet(t *testing.T) {
	svr := NewSnippetServiceServer()

	ctx := context.Background()
	req := &api_pb.GetSnippetRequest{}

	resp, err := svr.GetSnippet(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_SnippetServiceServer_CreateSnippet(t *testing.T) {
	svr := NewSnippetServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateSnippetRequest{}

	resp, err := svr.CreateSnippet(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_SnippetServiceServer_UpdateSnippet(t *testing.T) {
	svr := NewSnippetServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateSnippetRequest{}

	resp, err := svr.UpdateSnippet(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_SnippetServiceServer_DeleteSnippet(t *testing.T) {
	svr := NewSnippetServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteSnippetRequest{}

	resp, err := svr.DeleteSnippet(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
