package search

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v27/github"
	"github.com/pkg/errors"
	"github.com/rerost/query-recipe-api/repo/snippet"
)

type ghClient struct {
	githubClient *github.Client
	owner        string
	repository   string
}

func NewGHClient(githubClient *github.Client, owner string, repository string) snippet.SearchClient {
	return &ghClient{
		githubClient: githubClient,
		owner:        owner,
		repository:   repository,
	}
}

func (c *ghClient) Search(ctx context.Context, query string) ([]snippet.SnippetID, error) {
	githubSearchQuery := strings.Join([]string{
		query,
		fmt.Sprintf("repo:%s/%s", c.owner, c.repository),
	}, "+")
	result, _, err := c.githubClient.Search.Code(ctx, githubSearchQuery, nil)
	if err != nil {
		return nil, errors.Wrap(err, "GithubClient Search.Code error")
	}

	snippetIDs := make([]snippet.SnippetID, len(result.CodeResults))
	for i, r := range result.CodeResults {
		if r.Path == nil {
			return nil, errors.Wrapf(
				errors.New("Unexpected error Path is nil"),
				"CodeResult:%v",
				r,
			)
		}
		snippetIDs[i] = snippet.SnippetID(*r.Path)
	}
	return snippetIDs, nil
}

func (c *ghClient) Create(ctx context.Context, snippet snippet.Snippet) error {
	// Nothing todo
	return nil
}

func (c *ghClient) Update(ctx context.Context, snippet snippet.Snippet) error {
	// Nothing todo
	return nil
}

func (c *ghClient) Delete(ctx context.Context, snippetID snippet.SnippetID) error {
	// Nothing todo
	return nil
}
