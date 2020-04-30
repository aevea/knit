package github

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/shurcooL/githubv4"
)

// Client is a wrapper for easier use of Github
type Client struct {
	GHClient   *githubv4.Client
	Owner      string
	Repository string
}

// NewGithubClient creates a new githubClient based on the HTTPClient provided
func NewGithubClient(HTTPClient *http.Client, repository string) (*Client, error) {
	repoName := strings.Split(repository, "/")

	if len(repoName) != 2 {
		return nil, fmt.Errorf("expected repository format of owner/repository, but received %s", repository)
	}

	owner := repoName[0]
	repo := repoName[1]
	client := githubv4.NewClient(HTTPClient)

	return &Client{GHClient: client, Owner: owner, Repository: repo}, nil
}
