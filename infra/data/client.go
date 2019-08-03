package data

import (
	"context"

	"github.com/rerost/query-recipe-api/repo/snippet"
)

type ghClient struct {
}

func NewGHClient() snippet.DataClient {
	return nil
}

func (c *ghClient) Create(ctx context.Context, sp snippet.Snippet) (snippet.Snippet, error) {
	// TODO(@rerost) Implement
	return snippet.Snippet{}, nil
}
func (c *ghClient) Get(ctx context.Context, snippetID snippet.SnippetID) (snippet.Snippet, error) {
	return snippet.Snippet{}, nil
}
func (c *ghClient) BulkGet(ctx context.Context, snippetIDs []snippet.SnippetID) ([]snippet.Snippet, error) {
	return nil, nil
}
func (c *ghClient) Update(ctx context.Context, sp snippet.Snippet) error {
	// TODO(@rerost) Implement
	return nil
}
func (c *ghClient) Delete(ctx context.Context, snippetID snippet.SnippetID) error {
	// TODO(@rerost) Implement
	return nil
}
