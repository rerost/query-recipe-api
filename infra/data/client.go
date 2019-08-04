package data

import (
	"context"
	"io/ioutil"

	"github.com/google/go-github/v27/github"
	"github.com/pkg/errors"
	"github.com/rerost/query-recipe-api/repo/snippet"
)

type ghClient struct {
	githubClient *github.Client
	owner        string
	repository   string
}

func NewGHClient(githubClient *github.Client, owner, repository string) snippet.DataClient {
	return &ghClient{
		githubClient: githubClient,
		owner:        owner,
		repository:   repository,
	}
}

func (c *ghClient) Create(ctx context.Context, sp snippet.Snippet) (snippet.Snippet, error) {
	// TODO(@rerost) Implement
	return snippet.Snippet{}, nil
}

func convertIDToFilepath(snippetID snippet.SnippetID) (string, error) {
	path := string(snippetID)
	return path, nil
}

func (c *ghClient) getFile(ctx context.Context, filepath string) (string, error) {
	reader, err := c.githubClient.Repositories.DownloadContents(
		ctx,
		c.owner,
		c.repository,
		filepath,
		nil,
	)
	if err != nil {
		return "", errors.Wrap(err, "Failed to download contents")
	}
	defer reader.Close()

	contents, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", errors.Wrap(err, "Failed to ioutil.ReadAll(reader)")
	}

	return string(contents), nil
}

func (c *ghClient) Get(ctx context.Context, snippetID snippet.SnippetID) (snippet.Snippet, error) {
	filepath, err := convertIDToFilepath(snippetID)
	if err != nil {
		return snippet.Snippet{}, errors.Wrap(err, "Failed to convertIDToFilepath")
	}

	document, err := c.getFile(ctx, filepath+".md")
	if err != nil {
		return snippet.Snippet{}, errors.Wrap(err, "Failed to get document")
	}

	sql, err := c.getFile(ctx, filepath+".sql")
	if err != nil {
		return snippet.Snippet{}, errors.Wrap(err, "Failed to get sql")
	}

	return snippet.Snippet{
		ID:       snippetID,
		Document: string(document),
		SQL:      string(sql),
	}, nil
}
func (c *ghClient) BulkGet(ctx context.Context, snippetIDs []snippet.SnippetID) ([]snippet.Snippet, error) {
	snippets := []snippet.Snippet{}
	for _, id := range snippetIDs {
		s, err := c.Get(ctx, id)
		if err != nil {
			return nil, errors.Wrap(err, "GHClient Get error")
		}
		snippets = append(snippets, s)
	}
	return snippets, nil
}
func (c *ghClient) Update(ctx context.Context, sp snippet.Snippet) error {
	// TODO(@rerost) Implement
	return nil
}
func (c *ghClient) Delete(ctx context.Context, snippetID snippet.SnippetID) error {
	// TODO(@rerost) Implement
	return nil
}
