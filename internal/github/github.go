package github

import (
	"net/http"

	"github.com/shurcooL/githubv4"
)

// Client is a wrapper for easier use of Github
type Client struct {
	GHClient   *githubv4.Client
	Owner      string
	Repository string
}

// NewGithubClient creates a new githubClient based on the HTTPClient provided
func NewGithubClient(HTTPClient *http.Client, owner string, repository string) *Client {
	client := githubv4.NewClient(HTTPClient)

	return &Client{GHClient: client, Owner: owner, Repository: repository}
}
