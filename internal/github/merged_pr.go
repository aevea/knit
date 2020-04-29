package github

import (
	"context"
	"errors"

	"github.com/shurcooL/githubv4"
)

// MergedPRs gets the last 100 merged PRs, unless noLimit is true. It will proceed to get all PRs if possible at that point.
func (client *Client) MergedPRs(noLimit bool) ([]PR, error) {
	var mergedPRQuery mergedPRQuery

	variables := map[string]interface{}{
		"repositoryOwner": githubv4.String(client.Owner),
		"repositoryName":  githubv4.String(client.Repository),
		"prCursor":        (*githubv4.String)(nil),
	}

	var prs []PR

	for {
		err := client.GHClient.Query(context.Background(), &mergedPRQuery, variables)

		if err != nil {
			return nil, err
		}

		if len(mergedPRQuery.Repository.PullRequests.Nodes) == 0 {
			return nil, errors.New("No merged PR found")
		}

		for _, pr := range mergedPRQuery.Repository.PullRequests.Nodes {
			merged := pr.MergedAt.Sub(pr.CreatedAt)

			prs = append(prs, PR{Title: pr.Title, MergedAfter: merged, URL: pr.URL})
		}

		if !mergedPRQuery.Repository.PullRequests.PageInfo.HasPreviousPage {
			break
		}

		if !noLimit {
			break
		}

		variables["prCursor"] = githubv4.NewString(mergedPRQuery.Repository.PullRequests.PageInfo.StartCursor)
	}

	return prs, nil
}
