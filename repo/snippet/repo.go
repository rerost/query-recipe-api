package snippet

import (
	"context"

	"github.com/pkg/errors"
)

// SnippetID = /snippet/test_query (Not /snippet/test_query.sql, /snippet/test_query.md)
// TODO(@rerost) Think type. null.String or string or int64
type SnippetID string

type Snippet struct {
	ID       SnippetID
	Document string
	SQL      string
}

type Repo interface {
	// Create(ctx context.Context, snippet Snippet) (Snippet, error)
	// Get(ctx context.Context, snippetID SnippetID) (Snippet, error)
	// Update(ctx context.Context, snippet Snippet) error
	// Delete(ctx context.Context, snippetID SnippetID) error
	Search(ctx context.Context, query string) ([]Snippet, error)
}

type SearchClient interface {
	Search(ctx context.Context, query string) ([]SnippetID, error)
	// Create(ctx context.Context, snippet Snippet) error
	// Update(ctx context.Context, snippet Snippet) error
	// Delete(ctx context.Context, snippetID SnippetID) error
}

type DataClient interface {
	// Create(ctx context.Context, snippet Snippet) (Snippet, error)
	Get(ctx context.Context, snippetID SnippetID) (Snippet, error)
	BulkGet(ctx context.Context, snippetIDs []SnippetID) ([]Snippet, error) // Return in order
	// Update(ctx context.Context, snippet Snippet) error
	// Delete(ctx context.Context, snippetID SnippetID) error
}

type repoImpl struct {
	sc SearchClient
	dc DataClient
}

func NewRepo(sc SearchClient, dc DataClient) Repo {
	return &repoImpl{
		sc: sc,
		dc: dc,
	}
}

func (r *repoImpl) Create(ctx context.Context, snippet Snippet) (Snippet, error) {
	// s, err := r.dc.Create(ctx, snippet)
	// if err != nil {
	// 	return Snippet{}, errors.Wrap(err, "DataClient Create error")
	// }

	// err = r.sc.Create(ctx, snippet)
	// if err != nil {
	// 	// TODO(@rerost) Rollback Data
	// 	return Snippet{}, errors.Wrap(err, "SearchClient Register error")
	// }

	// return s, nil
	return Snippet{}, nil
}

func (r *repoImpl) Delete(ctx context.Context, snippetID SnippetID) error {
	// err := r.dc.Delete(ctx, snippetID)
	// if err != nil {
	// 	return errors.Wrap(err, "DataClient Delete error")
	// }

	// err = r.sc.Delete(ctx, snippetID)
	// if err != nil {
	// 	// TODO(@rerost) Rollback Data
	// 	return errors.Wrap(err, "SearchClient Delete error")
	// }

	return nil
}

func (r *repoImpl) Get(ctx context.Context, snippetID SnippetID) (Snippet, error) {
	s, err := r.dc.Get(ctx, snippetID)
	if err != nil {
		return Snippet{}, errors.Wrap(err, "DataClient Get error")
	}

	return s, nil
}

func (r *repoImpl) Update(ctx context.Context, snippet Snippet) error {
	// err := r.dc.Update(ctx, snippet)
	// if err != nil {
	// 	return errors.Wrap(err, "DataClient Update error")
	// }

	// err = r.sc.Update(ctx, snippet)
	// if err != nil {
	// 	// TODO(@rerost) Rollback Data
	// 	return errors.Wrap(err, "SearchClient Update error")
	// }

	return nil
}

func (r *repoImpl) Search(ctx context.Context, query string) ([]Snippet, error) {
	ids, err := r.sc.Search(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "SearchClient Search error")
	}

	ss, err := r.dc.BulkGet(ctx, ids)
	if err != nil {
		return nil, errors.Wrap(err, "DataClient BulkGet error")
	}
	return ss, nil
}
