package github

import (
	"context"
	"errors"
	"time"

	"github.com/shurcooL/githubv4"
)

// OldestPR queries Github for the first OPEN PR.
func (client *Client) OldestPR() (*PR, error) {
	var oldestPRQuery oldestPRQuery

	variables := map[string]interface{}{
		"repositoryOwner": githubv4.String(client.Owner),
		"repositoryName":  githubv4.String(client.Repository),
	}

	err := client.GHClient.Query(context.Background(), &oldestPRQuery, variables)

	if err != nil {
		return nil, err
	}

	if len(oldestPRQuery.Repository.PullRequests.Nodes) == 0 {
		return nil, errors.New("No oldest PR found")
	}

	oldestNode := oldestPRQuery.Repository.PullRequests.Nodes[0]

	open := time.Since(oldestNode.CreatedAt)

	return &PR{Title: oldestNode.Title, OpenFor: open, URL: oldestNode.URL}, nil
}
