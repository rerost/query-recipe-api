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

type Config interface {
	Owner() string
	Repository() string
}

func NewGHClient(githubClient *github.Client, config Config) snippet.SearchClient {
	return &ghClient{
		githubClient: githubClient,
		owner:        config.Owner(),
		repository:   config.Repository(),
	}
}

func (c *ghClient) Search(ctx context.Context, query string) ([]snippet.SnippetID, error) {
	githubSearchQuery := strings.Join([]string{
		fmt.Sprintf("repo:%s/%s", c.owner, c.repository),
		"language:markdown",
		query,
	}, " ")
	result, _, err := c.githubClient.Search.Code(ctx, githubSearchQuery, nil)
	if err != nil {
		return nil, errors.Wrap(err, "GithubClient Search.Code error")
	}

	snippetIDs := []snippet.SnippetID{}
	for _, r := range result.CodeResults {
		if r.Path == nil {
			return nil, errors.Wrapf(
				errors.New("Unexpected error Path is nil"),
				"CodeResult:%v",
				r,
			)
		}
		if !strings.HasSuffix(*r.Path, ".md") {
			return nil, errors.Wrapf(
				errors.New("Not sql file is passed"),
				"CodeResult:%v",
				r,
			)
		}
		id := strings.TrimSuffix(*r.Path, ".md")
		snippetIDs = append(snippetIDs, snippet.SnippetID(id))
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
